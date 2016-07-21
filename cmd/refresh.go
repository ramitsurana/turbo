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

var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Completely removes and re-installs docker",
	Long: `Completely removes and re-installs docker `,
	Run: func(cmd *cobra.Command, args []string) {			
		arg0 := "sudo"
                arg1 := "apt-get"
		arg2 := "purge"
		arg3 := "docker-engine"		
		arg4 := "autoremove"
		arg5 := "--purge"
		arg6 := "rm"
		arg7 := "-rf"
		arg8 := "/var/lib/docker"
		arg9 := "install"		    
		
		//Removing Docker Engine 
		fmt.Println("Removing Docker Engine ...")
                cmd1 := exec.Command(arg0, arg1, arg2, arg3)
                stdout1, err1 := cmd1.Output()		

		if err1 != nil {	 	                            
		      log.Printf("\nFailed to remove docker-engine")
		      println(err1.Error())		      		     		      		  
                      return
    		}
		print(string(stdout1))

		cmd2 := exec.Command(arg0, arg1, arg4, arg5, arg3)
		stdout2, err2 := cmd2.Output()

		if err2 != nil {	 	                            
		      log.Printf("\nFailed to remove docker-engine")
		      println(err2.Error())		      		     		      		  
                      return
    		}     
		print(string(stdout2))
		fmt.Println("Docker uninstalled")		

		//Removing Docker Data
		cmd3 := exec.Command(arg6, arg7, arg8)
		stdout3, err3 := cmd3.Output()
		if err3 != nil {                	
			log.Printf("\nFailed to remove Docker data")
			println(err3.Error())
			return			
        		}
		print(string(stdout3))		
		fmt.Println("Docker removal completed")						

		//Installing Docker
		cmd4 := exec.Command(arg0, arg1, arg9, arg3)
		stdout4, err4 := cmd4.Output()

		if err4 != nil {		      	 	                            
		      log.Printf("\nFailed to install docker-engine")		      
         	      println(err4.Error())
		      return
    		}     
		print(string(stdout4))
		fmt.Println("Docker installed")
		},			
}

func init() {
	RootCmd.AddCommand(refreshCmd)		
}
