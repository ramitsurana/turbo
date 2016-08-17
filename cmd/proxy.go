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
	"os"
	"log"			
	"github.com/spf13/cobra"	
)

var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Installs and configure squid3 proxy for your system",
	Long: `Install and configures squid3 proxy on your enviorment `,
	Run: func(cmd *cobra.Command, args []string) {
		_, err0 := exec.LookPath("squid")
		if err0 == nil {
			log.Fatal(err0)
			log.Printf("squid is already installed.")		
			os.Exit(1)
			}		
		fmt.Printf("squid is not installed.Installing ...")
				
        	fmt.Println("Downloading & Running squid3 proxy for your system ...")				
		cmd1 := exec.Command("docker", "run", "--net", "host", "-d", "jpetazzo/squid-in-a-can", "-v", "/home/user/persistent_squid_cache:/var/cache/squid3")
		err1 := cmd1.Start()
		if err1 != nil {
                	log.Fatal(err1)
			log.Printf("Downloading failed")
			os.Exit(1)			
        		}
		err1 = cmd1.Wait()		
		fmt.Println("Downloading Successfull")
				
		fmt.Printf("Configuring Iptables ...")				
		cmd2 := exec.Command("iptables", "-t", "nat", "-A", "PREROUTING", "-p", "tcp", "--dport", "80", "-j", "REDIRECT", "--to", "3129", "-w")
		err2 := cmd2.Start()
		if err2 != nil {
                	log.Fatal(err2)
			log.Printf("Failed to start squid")
			os.Exit(1)
        		}				
		fmt.Println("squid is installed")			
		

		fmt.Printf("Clearing Iptables configurations...")				
		cmd3 := exec.Command("iptables", "-t", "nat", "-D", "PREROUTING", "-p", "tcp", "--dport", "80", "-j", "REDIRECT", "--to", "3129", "-w")
		err3 := cmd3.Start()
		if err3 != nil {
                	log.Fatal(err3)
			log.Printf("Failed to start squid")
			os.Exit(1)
        		}				
		fmt.Println("squid is installed and set for use")					       										
	},
}

func init() {
	RootCmd.AddCommand(proxyCmd)		
}
