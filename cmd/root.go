// Copyright © 2017 Robert Anthony Bellamy  rabellamy@gmail.com
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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// Represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "random",
	Short: "proceeding, made, or occurring without definite aim, reason, or pattern",
	Long: `adjective
    1. proceeding, made, or occurring without definite aim, reason, or pattern:
       the random selection of numbers.
    2. Statistics. of or characterizing a process of selection in which each item of
       a set has an equal probability of being chosen.`,
}

// Adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.random.yaml)")

	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enables ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".ramndom") // name of config file (without extension)
	viper.AddConfigPath("$HOME")    // Adds home directory as first search path
	viper.AutomaticEnv()            // Reads in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
