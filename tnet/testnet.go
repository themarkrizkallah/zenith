// Copyright © Mark Rizkallah <mark.g.rizkallah@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tnet

import (
	"fmt"
	"log"
	"net/http"

	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
)

// FundAccount funds an account on the testnet using friendbot
func FundAccount(address string) {
	resp, err := http.Get("https://friendbot.stellar.org/?addr=" + address)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
}

// LoadAccount loads account from testnet
func LoadAccount(address string) (horizon.Account, error) {
	account, err := horizon.DefaultTestNetClient.LoadAccount(address)
	return account, err
}

// MakeTransaction creates, signs, and submits a transaction
func MakeTransaction(amount, source, destination, memo string) error {
	tx, err := build.Transaction(
		build.TestNetwork,
		build.SourceAccount{AddressOrSeed: source},
		build.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		build.Payment(
			build.Destination{AddressOrSeed: destination},
			build.NativeAmount{Amount: amount},
		),
	)
	if err != nil {
		return err
	}

	// Sign the transaction
	txe, err := tx.Sign(source)
	if err != nil {
		return err
	}

	txeB64, err := txe.Base64()
	if err != nil {
		return err
	}

	// Submit transaction
	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)
	if err != nil {
		return err
	}

	fmt.Println("Transaction Successful!")
	fmt.Println("Ledger:", resp.Ledger)
	fmt.Println("Hash:", resp.Hash)

	return err
}
