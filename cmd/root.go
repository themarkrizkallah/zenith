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
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Flags
var (
	cfgFile, currency       string
	baseReserve, minBalance int
	isTest                  bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "zenith",
	Short: "A Go based Stellar wallet!",
	Long: `A Go based Stellar wallet!
Here's information on how to use it:`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.zenith.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&isTest, "tnet", "t", false, "use testnet instead of the real network")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetDefault("CURRENCY", "XLM")
	viper.SetDefault("BASE_RESERVE", 0.5)
	viper.SetDefault("MIN_BALANCE", 1)

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".zenith-cli" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/go/src/github.com/themarkrizkallah/zenith")
		viper.SetConfigName(".zenith")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		currency = viper.GetString("CURRENCY")
		baseReserve = viper.GetInt("BASE_RESERVE")
		minBalance = viper.GetInt("MIN_BALANCE")
	}
}
