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
	"os"			
	"github.com/spf13/cobra"
)

var ldapCmd = &cobra.Command{
	Use:   "ldap",
	Short: "Uses to install and configure Openldap",
	Long: `Uses to install and configure Openldap `,
	Run: func(cmd *cobra.Command, args []string) {				
		fmt.Println("Downloading Openldap ...")		
		cmd1 := exec.Command("docker", "pull", "--name", "openldap", "--detach", "osixia/openldap:1.1.5")
		err1 := cmd1.Start()
		if err1 != nil {
                	log.Fatal(err1)
			log.Printf("Downloading failed")
			os.Exit(1)			
        		}		
		fmt.Println("Pulling Successfull")

		fmt.Println("Running Openldap ...")		
		cmd2 := exec.Command("docker", "run", "--name", "openldap", "--detach", "osixia/openldap:1.1.5")
		err2 := cmd2.Start()
		if err2 != nil {
                	log.Fatal(err2)
			log.Printf("Unable to run the container")
			os.Exit(1)
        		}				
		fmt.Println("Running Openldap ...")									
	},
}

func init() {
	RootCmd.AddCommand(ldapCmd)		
}
