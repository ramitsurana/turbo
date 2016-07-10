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
		fmt.Println("Monitoring images ...")
		cmd1 := exec.Command("docker", "rmi", "`docker", "images", "-a", "-q`")
		err1 := cmd1.Start()
		if err1 != nil {
                	log.Fatal(err1)
			log.Printf("Monitoring failed ...")
        	}		
		fmt.Println("")
	},
}

func init() {
	RootCmd.AddCommand(monitorCmd)		
}
