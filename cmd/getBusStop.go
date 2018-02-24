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

// getBusStopCmd represents the getBusStop command
var getBusStopCmd = &cobra.Command{
	Use:   "getBusStop",
	Short: "It gets bus stop id",
	Long: `It gets bus stop id

Command requires following arguments: BusName`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("api-key")
		auth := authenticator.NewAuthenticator(apiKey)
		client := umwarsawclient.NewClient("https://api.um.warszawa.pl/api/", auth, nil)
		busStop, err := client.GetBusStop(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		busStopJson, _ := json.MarshalIndent(busStop, "", "  ")
		fmt.Println(string(busStopJson))
	},
}

func init() {
	rootCmd.AddCommand(getBusStopCmd)
}
