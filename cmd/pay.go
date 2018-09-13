// Copyright © 2018 Mark Rizkallah <mark.g.rizkallah@gmail.com>
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

package cmd

import (
	"fmt"
	"log"

	"github.com/themarkrizkallah/zenith/client"
	"github.com/themarkrizkallah/zenith/tnet"

	"github.com/spf13/cobra"
)

func pay() {
	var (
		amount      string // XLM amount
		destination string // Receiver's secret key
		source      string // Sender's public key
		memo        string // Optional transaction memo
		err         error
	)

	fmt.Print("Amount to send (XLM): ")
	fmt.Scan(&amount)

	fmt.Print("Destination's address: ")
	fmt.Scan(&destination)

	_, err = tnet.LoadAccount(destination)
	if err != nil {
		log.Println("Error loading account")
		log.Fatal(err)
	}

	fmt.Print("Memo (optional): ")
	fmt.Scanln(&memo)

	fmt.Print("Sender's secret key: ")
	fmt.Scanln(&source)

	if isTest {
		err = tnet.MakeTransaction(amount, source, destination, memo)
	} else {
		err = client.MakeTransaction(amount, source, destination, memo)
	}

	if err != nil {
		log.Println("Transaction failed")
		log.Fatal(err)
	}
}

// payCmd represents the pay command
var payCmd = &cobra.Command{
	Use:   "pay",
	Short: "Send XLM",
	Long: `Send XLM. Can be used as follows:

zenith-cli pay
`,
	Run: func(cmd *cobra.Command, args []string) {
		pay()
	},
}

func init() {
	rootCmd.AddCommand(payCmd)
}
