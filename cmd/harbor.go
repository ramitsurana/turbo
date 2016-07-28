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

var harborCmd = &cobra.Command{
	Use:   "harbor",
	Short: "installs and configures vmware harbor",
	Long: `Install and configures vmware harbor on your enviorment `,
	Run: func(cmd *cobra.Command, args []string) {				
		fmt.Println("Prerequisites : \nDocker-Compose: v1.6.0+\nDocker: v1.10.0+")
				
		fmt.Println("Cloning Harbor ..")
		cmd1 := exec.Command("git", "clone", "https://github.com/vmware/harbor")
		err1 := cmd1.Start()
		if err1 != nil {                	
			log.Printf("Cloning Failed")
			os.Exit(1)			
        		}		
		fmt.Println("Cloning Successfull ..")
		
		fmt.Println("Checking Harbor.cfg ...\n")
		fmt.Println("Opening Harbor.cfg ...\n")
		fmt.Println("Please input the configorations according to your system ...\n")

		fi, err2 := os.Open("harbor.cfg")
		if err2 != nil {
		        panic(err2)	
		}
    		// close fi on exit and check for its returned error
		defer func() {
	        if err3 := fi.Close(); err3 != nil {
        	    panic(err3)
        	}
		}()							
								
		
		cmd2 := exec.Command("cd", "Deploy")
		err4 := cmd2.Start()
		if err4 != nil {
                	log.Fatal(err4)
			log.Printf("Cannot find Deploy")
			os.Exit(1)
        		}				
		fmt.Println("Running Prepare.sh ... ")					
		
		cmd3 := exec.Command("./prepare")
		err5 := cmd3.Start()
		if err5 != nil {
                	log.Fatal(err5)
			log.Printf("Cannot run prepare.sh")
			os.Exit(1)
        		}				
		fmt.Println("Using Docker compose up ...")					
		
		cmd4 := exec.Command("sudo", "docker-compose", "up", "-d")
		err6 := cmd4.Start()
		if err6 != nil {
                	log.Fatal(err6)
			log.Printf("Failed to use docker-compose up.")
			os.Exit(1)
        		}				
		fmt.Println("Harbor completely installed & configured to use.Please open the hostname as url in your browser.")		
		},
}

func init() {
	RootCmd.AddCommand(harborCmd)		
}
