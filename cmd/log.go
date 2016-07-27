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
	"github.com/spf13/cobra"
	"os"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Uses logspout to collect your docker logs",
	Long: `Runs and configures logspout on your enviorment `,
	Run: func(cmd *cobra.Command, args []string) {				
		fmt.Println("Running logspout ...")
		cmd1 := exec.Command("docker", "run", "-d",
    		"--volume=/var/run/docker.sock:/var/run/docker.sock",
		"--publish=127.0.0.1:8000:80",
   		 "gliderlabs/logspout")
		err1 := cmd1.Start()
		if err1 != nil {
                	log.Fatal(err1)
			log.Printf("Failed to run logspout")
			os.Exit(1)			
        		}		
		fmt.Println("Logspout has successfully collected logs.Try running curl http://127.0.0.1:8000/logs")					
		},		
}

func init() {
	RootCmd.AddCommand(logCmd)		
}
