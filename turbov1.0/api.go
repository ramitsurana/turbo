// Copyright © 2016 Ramit Surana <ramitsurana@gmail.com>
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
	"os/exec"
	"log"
	"flag"
	"net/http"			
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Serves an GUI for turbo",
	Long: `Installs and Serves GUI for turbo `,
	Run: func(cmd *cobra.Command, args []string) {				
		fmt.Println(" Serving API...")
		var addr = flag.String("addr", ":9000", "website address")
		flag.Parse()
		mux := http.NewServeMux()
		mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("public"))))
		log.Println("Serving website at:", *addr)
		http.ListenAndServe(*addr, mux)					
	        },
}

func init() {
	RootCmd.AddCommand(apiCmd)		
}
