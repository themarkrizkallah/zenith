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
	"database/sql"
	"fmt"
	"github.com/matthewhartstonge/argon2"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
	"syscall"
)

func comparePassToHash(password, dbpass []byte) bool {
	success, err := argon2.VerifyEncoded(password, dbpass)
	if err != nil {
		log.Printf("Failed to verify.\n%s\n", err)
		os.Exit(1)
	}

	return success
}

func login() {
	var (
		username string
		password string
		uid      string
		success  bool
	)

	// Connect to database
	dbinfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable",
		DB_USER, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Printf("Failed to connect to database.\n%s\n", err)
		os.Exit(1)
	}
	defer db.Close()

	fmt.Print("username: ")
	fmt.Scan(&username)

	fmt.Print("password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Print("\n")
	if err != nil {
		log.Printf("Password read failed.\n%s\n", err)
		os.Exit(1)
	}

	query := fmt.Sprintf("SELECT id, password FROM users WHERE username='%s'", username)
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("User not found.\n%s\n", err)
		os.Exit(1)
	}

	for rows.Next() {
		err = rows.Scan(&uid, &password)
		if err != nil {
			log.Printf("Error retrieving password.\n%s\n", err)
			os.Exit(1)
		}
	}

	success = comparePassToHash(bytePassword, []byte(password))
	if success {
		fmt.Println("Log in successful")
		os.Setenv("USER_ID", string(uid))
	}
}

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log into your Zenith account",
	Long: `Log into your Zenith account. 
Currently this only sets an environment variable that does not persist.`,
	Run: func(cmd *cobra.Command, args []string) {
		if userId == "-1" {
			login()
		} else {
			fmt.Println("Please log out first before you attempt to log in again.")
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
