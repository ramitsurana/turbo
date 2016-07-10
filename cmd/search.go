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
	Short: "Search images from registries",
	Long: `Search command searches for the image
from various registries and prints the result.Currently using Docker registry,quay.io and gcr.io`,
	Run: func(cmd *cobra.Command, args []string) {		
		fmt.Println("Searching your image via Docker registry")
		cmd1 := exec.Command("docker","search","$")
		err1 := cmd1.Start()
		//Prints the result
		if err1 != nil {
                	log.Fatal(err1)
			log.Printf("Unable to get info from docker registry")
        	}
		fmt.Println("Searching your image via Quay.io registry")
		cmd2 := exec.Command("docker","search","$")
		err2 := cmd2.Start()
		if err2 != nil {
                	log.Fatal(err2)
			log.Printf("Unable to make docker-backup")
        	}		
		fmt.Println("Searching your image via Gcr.io registry")		 
		cmd3 := exec.Command("docker","search","$")
		err3 := cmd3.Start()
		if err3 != nil {
                	log.Fatal(err3)
			log.Printf("Unable to shift data into docker-backup")
        	}
		fmt.Println("Operation Successful.You can check the images at $HOME/docker-backup")
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)
}
