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
	"log"
	"os/exec"
	"github.com/spf13/cobra"
)


var kickstartCmd = &cobra.Command{
	Use:   "kickstart <W.I.P.>",
	Short: "restarts all your containers quickly",
	Long: `Restarts all your containers quickly`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Scanning Containers ...")
		arg0 := "docker"
                arg1 := "restart"
                arg2 := "$(docker"
                arg3 := "ps"
                arg4 := "-a"
                arg5 := "-q)"                 

		cmd1 := exec.Command(arg0, arg1, arg2, arg3, arg4, arg5)
                err1 := cmd1.Start()

    		if err1 != nil {	 	      
                      //println(err1.Error())
		      log.Fatal(err1)
		      log.Printf("\nFailed to restart containers ...")
                      //return
    		}
		
		fmt.Println("Your containers are up and running")
		},			
}

func init() {
	RootCmd.AddCommand(kickstartCmd)	
}
