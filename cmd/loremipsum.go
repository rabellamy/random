// Copyright © 2017 Robert Anthony Bellamy rabellamy@gmail.com
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
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// loremipsumCmd represents the loremipsum command
var loremipsumCmd = &cobra.Command{
	Use:   "loremipsum",
	Short: "Get some random lorem ipsum.",
	Long: `Random loremipsum’

Form the Loripsum.net, the ‘lorem ipsum’ generator that doesn't suck.`,
	Run: getLoremIpsum,
}

func init() {
	RootCmd.AddCommand(loremipsumCmd)
}

func getLoremIpsum(cmd *cobra.Command, args []string) {
	response, err := http.Get("http://loripsum.net/api/plaintext")

	// Defer the closing of the body
	defer response.Body.Close()

	if err != nil {
		log.Println("err happened", err)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))
}
