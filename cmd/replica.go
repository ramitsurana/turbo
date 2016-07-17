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
	"strings"	
	"github.com/spf13/cobra"
)

var replicaCmd = &cobra.Command{
	Use:   "replica",
	Short: "To create Replicas of your containers",
	Long: `To create Replicas of your containers`,
	Run: func(cmd *cobra.Command, args []string) {
		
	        fmt.Println("No. of Containers to create " + strings.Join(args, " "))
		
		//Will be Using loop in the next release
		fmt.Println("Creating 1 st container ...")
		cmd1 := exec.Command("docker", "create", "$2")
		err1 := cmd1.Start()
		if err1 != nil {
                	log.Fatal(err1)
			log.Printf("1 st container has failed ...")
        	}		
		fmt.Println("1st container created")
	
		fmt.Println("Creating 2 nd container ...")
		cmd2 := exec.Command("docker", "create", "$2")
		err2 := cmd2.Start()
		if err2 != nil {
                	log.Fatal(err2)
			log.Printf("2 nd container has failed ...")
        	}		
		fmt.Println("2nd container created")
	
		fmt.Println("Creating 3 rd container ...")
		cmd3 := exec.Command("docker", "create", "$2")
		err3 := cmd3.Start()
		if err3 != nil {
                	log.Fatal(err3)
			log.Printf("3 rd container has failed ...")
        	}		
		fmt.Println("All the containers are created.Here are there ID's:")
	},
}	

func init() {
	RootCmd.AddCommand(replicaCmd)		
}
