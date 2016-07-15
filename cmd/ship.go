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
	"bufio"
	"os"	
	"github.com/spf13/cobra"
)

var shipCmd = &cobra.Command{
	Use:   "ship",
	Short: "Transfer your docker images over a remote i.p.",
	Long: `ship up your docker images `,
	Run: func(cmd *cobra.Command, args []string) {
		
		//Taking image name from the user
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Write the Image name: ")
		image, _ := reader.ReadString('\n')						
		
	        arg0 := "sudo"
                arg1 := "docker"
                arg2 := "save"
                arg3 := "-o"
                arg4 := "load"
                arg5 := "-i"
                arg6 := "$HOME/ship" //Path of the tar file 
		arg7 := "docker-machine"
		arg8 := "scp"
		arg9 := "-r"    
		
		//Converting images into tar files
		fmt.Println("Converting images as tar files...")
				
                cmd1 := exec.Command(arg0, arg1, arg2, arg3, arg4, arg6, image)
                err1 := cmd1.Start()

    		if err1 != nil {	 	                            
		      log.Fatal(err1)
		      log.Printf("Failed to create tar file")                      
    		}

		//Using scp to send the files
		fmt.Println("Setting up scp ...")
		
		cmd2 := exec.Command(arg7, arg8, arg9)
                err2 := cmd2.Start()

    		if err2 != nil {	 	                            
		      log.Fatal(err2)
		      log.Printf("\nFailed to setup scp")                      
    		}
		
		fmt.Println("Your Images have been shipped")
		fmt.Println("\nLoading your images ...")
		
		//Loading tar files as Image files
		cmd3 := exec.Command(arg0, arg1, arg4, arg5, arg6)
                err3 := cmd3.Start()

    		if err3 != nil {	 	                            
		      log.Fatal(err3)
		      log.Printf("\nFailed to convert tar files.")                      
    		}
		
		fmt.Println("Done")
		},
}

//Saves image as tar file		
//sudo docker save -o <save image to path> <image name>

//Loads image from tar file
//sudo docker load -i <path to image tar file>

//Using scp
//docker-machine scp 

func init() {
	RootCmd.AddCommand(shipCmd)		
}
