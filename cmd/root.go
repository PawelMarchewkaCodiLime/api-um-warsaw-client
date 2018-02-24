// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "api-um-warsaw-client",
	Short: "This is simple cli wrapper on umwarsawclient library",
	Long: `It allows user to perform simple queries to Warsaw UM Api and get pieces of information:

Example calls:
api-um-warsaw-client.exe --api-key xxxx-xxxx-xxxx-xxxx-xxx getBusStop znana
api-um-warsaw-client.exe --api-key xxxx-xxxx-xxxx-xxxx-xxx getLinesAtBusStop 5104 01
api-um-warsaw-client.exe --api-key xxxx-xxxx-xxxx-xxxx-xxx getTimeTable 5104 01 155`,

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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().String("api-key", "", "UM api key")
	viper.BindPFlag("api-key", rootCmd.PersistentFlags().Lookup("api-key"))
}
