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
	"github.com/stellar/go/clients/horizon"
)

// Get account from stdin
func readAccount() horizon.Account {
	var (
		address string
		account horizon.Account
		err     error
	)

	fmt.Print("Enter account ID (public key): ")
	fmt.Scan(&address)

	if isTest {
		account, err = tnet.LoadAccount(address)
		if err != nil {
			log.Println("Error loading account")
			log.Fatal(err)
		}
	} else {
		account, err = client.LoadAccount(address)
		if err != nil {
			log.Println("Error loading account")
			log.Fatal(err)
		}
	}

	return account
}

// Returns the native balance in account
func getBalance(account horizon.Account) string {
	balance, _ := account.GetNativeBalance()
	return balance
}

// Print the XLM balance of an account
func printBalance(account horizon.Account) {
	address := account.AccountID
	balance := getBalance(account)

	fmt.Println("Balances for account:", address)
	fmt.Println("XLM Balance:", balance)
}

// Display account balance
func displayBalance(args []string) {
	var (
		account  horizon.Account
		accounts []horizon.Account
		err      error
	)

	if len(args) > 0 {
		for _, arg := range args {
			if isTest {
				account, err = tnet.LoadAccount(arg)
			} else {
				account, err = client.LoadAccount(arg)
			}

			if err == nil {
				accounts = append(accounts, account)
			}
		}
	}

	if len(accounts) > 0 {
		for _, account = range accounts {
			printBalance(account)
		}
	} else {
		account = readAccount()
		printBalance(account)
	}
}

// balanceCmd represents the balance command
var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Display the balance of an account",
	Long: `Display the balance of an account, can be used as follows:

zenith-cli balance 
or
zenith-cli balance [address(es)]`,
	Run: func(cmd *cobra.Command, args []string) {
		displayBalance(args)
	},
}

func init() {
	rootCmd.AddCommand(balanceCmd)
}
