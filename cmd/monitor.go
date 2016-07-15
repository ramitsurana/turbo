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

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "To monitor your containers",
	Long: `It monitors your docker containers`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Checking ...")

		arg0 := "glances"
                arg1 := "-w"
                arg2 := "$2"    
				 		
                cmd1 := exec.Command(arg2)
                stdout1, err1 := cmd1.Output()

		if err1 != nil {	 	                            
		      log.Printf("\nFailed to search.Please check your Internet connection.")
		      println(err1.Error())		      		     		      		  
                      return
    		}    
		print(string(stdout1))

		fmt.Println("Starting monitoring ...")
		cmd2 := exec.Command(arg0, arg1)
		err2 := cmd2.Start()
		if err2 != nil {
                	log.Fatal(err2)
			log.Printf("Monitoring failed ...")
        	}		
		fmt.Println("You can view the monitor at http://<your-ip>/")
	},
}

func init() {
	RootCmd.AddCommand(monitorCmd)		
}
