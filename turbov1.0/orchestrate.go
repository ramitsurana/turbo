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
)

var orchestrateCmd = &cobra.Command{
	Use:   "orchestrate",
	Short: "Installs and configure kubernetes",
	Long: `Install and configures kubernetes on your enviorment `,
	Run: func(cmd *cobra.Command, args []string) {				
		fmt.Println("Downloading kubernetes ...")		
		cmd1 := exec.Command("wget", "https://github.com/kubernetes/kubernetes/releases/download/v1.3.0/kubernetes.tar.gz")
		err1 := cmd1.Start()
		if err1 != nil {
                	log.Fatal(err1)
			log.Printf("Downloading failed")
			exit()			
        		}		
		fmt.Println("Downloading Successfull")

		fmt.Println("Installing kubernetes ...")		
		cmd2 := exec.Command("tar", "xzvf", "rkt-v1.10.0.tar.gz")
		err2 := cmd2.Start()
		if err2 != nil {
                	log.Fatal(err2)
			log.Printf("Unable to Untar the file")
        		}				
		fmt.Println("Setting path ...")

		cmd3 := exec.Command("alias", "rkt=",""sudo", "'${PWD}/rkt-v1.8.0/rkt'"")
		err3 := cmd3.Start()
		if err3 != nil {
                	log.Fatal(err3)
			log.Printf("Path not set.Try using ./rkt")
        		}				
		fmt.Println("rkt is installed and configured")			
		}		
	},
}

func init() {
	RootCmd.AddCommand(orchestrateCmd)		
}
