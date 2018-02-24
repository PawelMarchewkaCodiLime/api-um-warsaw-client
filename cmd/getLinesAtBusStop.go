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

	"github.com/kaweue/api-um-warsaw-client/lib/authenticator"
	"github.com/kaweue/api-um-warsaw-client/lib/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"encoding/json"
)

// getLinesAtBusStopCmd represents the getLinesAtBusStop command
var getLinesAtBusStopCmd = &cobra.Command{
	Use:   "getLinesAtBusStop",
	Short: "It gets information about lines at a bus stop",
	Long: ` It gets information about lines at a bus stop

Command requires following arguments: BusStopId BusStopNr`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("api-key")
		auth := authenticator.NewAuthenticator(apiKey)
		client := umwarsawclient.NewClient("https://api.um.warszawa.pl/api/", auth, nil)
		lines, err := client.GetLinesOnBusStop(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		linesJson, _ := json.MarshalIndent(lines, "", "  ")
		fmt.Println(string(linesJson))
	},
}

func init() {
	rootCmd.AddCommand(getLinesAtBusStopCmd)
}
