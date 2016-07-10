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

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "backups all your docker stuff",
	Long: `backups all your Docker images`,
	Run: func(cmd *cobra.Command, args []string) {		
		fmt.Println("checking your info on docker ...")
		cmd1 := exec.Command("docker","info")
		err1 := cmd1.Start()
		if err1 != nil {
                	log.Fatal(err1)
			log.Printf("Unable to get info on docker")
        	}
		fmt.Println("Copying data to docker-backup ...")
		cmd2 := exec.Command("mkdir","-p","$HOME/docker-backup")
		err2 := cmd2.Start()
		if err2 != nil {
                	log.Fatal(err2)
			log.Printf("Unable to make docker-backup")
        	}		
		fmt.Println("Copying data to docker-backup ...")		 
		cmd3 := exec.Command("sudo", "cp","/var/lib/docker","$HOME/docker-backup")
		err3 := cmd3.Start()
		if err3 != nil {
                	log.Fatal(err3)
			log.Printf("Unable to shift data into docker-backup")
        	}
		fmt.Println("Operation Successful.You can check the images at $HOME/docker-backup")
	},
}

func init() {
	RootCmd.AddCommand(backupCmd)
}
