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



var rktCmd = &cobra.Command{
	Use:   "rkt",
	Short: "Installs and configures rkt",
	Long: `Install and configures rkt on your enviorment `,
	Run: func(cmd *cobra.Command, args []string) {
		//Downloading rkt				
		fmt.Println("Downloading rkt ...")		
		cmd1 := exec.Command("wget", "https://github.com/coreos/rkt/releases/download/v1.8.0/rkt-v1.10.0.tar.gz")
		err1 := cmd1.Start()
		if err1 != nil {
                	log.Fatal(err1)			
			log.Printf("Downloading failed.Have you got wget installed ?")
			os.Exit(1)								
        		}		
		fmt.Println("Downloading Successfull")

		//Installing rkt
		fmt.Println("Installing rkt ...")		
		cmd2 := exec.Command("tar", "xzvf", "rkt-v1.10.0.tar.gz")
		err2 := cmd2.Start()
		if err2 != nil {
                	log.Fatal(err2)
			log.Printf("Unable to Untar the file")
			os.Exit(1)
        		}				
		fmt.Println("Changing directory ...")

		//Setting the path
		cmd3 := exec.Command("cd","rkt-v1.10.0/")
		err3 := cmd3.Start()
		if err3 != nil {
                	log.Fatal(err3)
			log.Printf("Cannot get into the dir")
			os.Exit(1)
        		}				
		fmt.Println("rkt is installed and configured.Try using ./rkt")			
		},			
}

func init() {
	RootCmd.AddCommand(rktCmd)		
}
