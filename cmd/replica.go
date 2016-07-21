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
	//"strings"	
	"github.com/spf13/cobra"
)

var replicaCmd = &cobra.Command{
	Use:   "replica",
	Short: "To create Replicas of your containers",
	Long: `To create Replicas of your containers`,
	Run: func(cmd *cobra.Command, args []string) {
		var x int
		var y string
		
		fmt.Println("No. of Containers to create:")
		fmt.Scanf("%d",&x)

		fmt.Println("Name of the image:")
		fmt.Scanf("%s",&y)	        

		for i := 0; i < x; i++ {	        
		cmd1 := exec.Command("docker", "create", "&y")
		err1 := cmd1.Start()
		if err1 != nil {
                	log.Fatal(err1)
			log.Printf("Creating Containers has failed ...")
        	}		
		fmt.Println(i, "st Container created")
    		}
	},				
}	

func init() {
	RootCmd.AddCommand(replicaCmd)		
}
