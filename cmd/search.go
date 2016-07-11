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

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search images from multiple registries",
	Long: `Search command searches for the image
from various registries and prints the result.Currently using Docker registry,quay.io and gcr.io`,
	Run: func(cmd *cobra.Command, args []string) {				
		arg0 := "docker"
                arg1 := "search"
                arg2 := "$2"    
		
		//Searching via Docker Registry 
		fmt.Println("Searching your image via Docker registry")
                cmd1 := exec.Command(arg0, arg1, arg2)
                stdout1, err1 := cmd1.Output()

		if err1 != nil {	 	                            
		      log.Printf("\nFailed to search via Docker Registry.Please check your Internet connection.")
		      println(err1.Error())		      		     		      		  
                      return
    		}    
		print(string(stdout1))		
		

		//Searching via Quay.io
		fmt.Println("Searching your image via Quay.io registry")
		cmd2 := exec.Command(arg0, arg1, arg2)
                stdout2, err2 := cmd2.Output()
		

		if err2 != nil {	 	      
                      println(err2.Error())		      
		      log.Printf("Failed to search via Quay.io.Please check your Internet connection.")
		      return                      		      
    		}

		print(string(stdout2))    		     		
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)
}
