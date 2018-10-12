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
	"github.com/spf13/cobra"
	"github.com/stellar/go/keypair"
	"github.com/themarkrizkallah/zenith/tnet"
	"log"
)

func generateWallet() {
	pair, err := keypair.Random()
	if err != nil {
		log.Println("Failed to generate wallet")
		log.Fatal(err)
	}

	fmt.Println("Your public key is:", pair.Address())
	fmt.Println("Your secret key is:", pair.Seed())
	fmt.Println("Save both of those keys offline somewhere safe.")
	fmt.Printf("Note: The account must have a minimum of %d %q in order to be active.\n", minBalance, currency)

	if isTest {
		log.Println("Funding the account on the testnet...")
		tnet.FundAccount(pair.Address())
	}
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a new stellar wallet",
	Long:  `Generate a stellar public and seed keypair to be saved offline`,
	Run: func(cmd *cobra.Command, args []string) {
		generateWallet()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
