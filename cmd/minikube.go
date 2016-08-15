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
	"runtime"
)

var minikubeCmd = &cobra.Command{
	Use:   "minikube",
	Short: "Installs and configure minikube for your system",
	Long: `Install and configures minikube on your enviorment `,
	Run: func(cmd *cobra.Command, args []string) {
		
		_, err0 := exec.LookPath("minikube")
		if err0 == nil {			
			fmt.Println("Minikube is already installed.Exiting ...")		
			os.Exit(1)
			}		
		fmt.Printf("Minikube is not installed.Installing ...")

		arg0 := "curl"
                arg1 := "-Lo"
                arg2 := "minikube"                
                arg3 := "&&"
                arg4 := "chmod"
                arg5 := "+x"
		arg6 := "sudo"
		arg7 := "mv"
		arg8 := "/usr/local/bin"
		arg9 := "kubectl"
		arg10 := "https://storage.googleapis.com/minikube/releases/v0.7.0/minikube-linux-amd64"
		arg11 := "https://storage.googleapis.com/minikube/releases/v0.7.0/minikube-darwin-amd64"
		arg12 := "http://storage.googleapis.com/kubernetes-release/release/v1.3.0/bin/linux/amd64/kubectl"
		arg13 := "http://storage.googleapis.com/kubernetes-release/release/v1.3.0/bin/darwin/amd64/kubectl"
		arg14 := "start"
		
		//For Windows
		if runtime.GOOS == "windows" {
    		fmt.Println("Downloading minikube for Windows ...")
		fmt.Println("Sorry minikube is still under experimental phase for this.")
		os.Exit(1)
		}

		//For Linux
		if runtime.GOOS == "linux" {
        	fmt.Println("Downloading minikubev0.7 for Linux...")				
		cmd1 := exec.Command(arg0, arg1, arg2, arg10, arg3, arg4, arg5, arg2, arg3, arg6, arg7, arg2, arg8)
		err1 := cmd1.Start()
		if err1 != nil {
                	log.Fatal(err1)
			log.Printf("Downloading failed")
			os.Exit(1)			
        		}
		err1 = cmd1.Wait()		
		fmt.Println("Downloading Successfull")

		fmt.Println("Checking for kubectl ...")
		_, err2 := exec.LookPath("kubectl")
		if err2 != nil {
		log.Fatal(err2)
			log.Printf("Kubectl is not installed.")		
			}		
		fmt.Printf("Installing kubectl ...")
		
		cmd2 := exec.Command(arg0, arg1, arg9, arg12, arg3, arg4, arg5, arg9, arg3, arg6, arg7, arg9, arg8)
		err3 := cmd2.Start()
		if err3 != nil {
                	log.Fatal(err3)
			log.Printf("Unable to install kubectl")
        		}
		err3 = cmd2.Wait()				
		fmt.Println("Kubectl installed.")
				
		cmd3 := exec.Command(arg2,arg13)
		err4 := cmd3.Start()
		if err4 != nil {
                	log.Fatal(err4)
			log.Printf("Failed to start minikube")
			os.Exit(1)
        		}
		err4 = cmd3.Wait()				
		fmt.Println("Minikube is installed and set for use")
		}			
		

		//For Mac
		if runtime.GOOS == "osx" {        	
		fmt.Println("Downloading minikubev0.7 for Mac...")				
		cmd4 := exec.Command(arg0, arg1, arg2, arg11, arg3, arg4, arg5, arg2, arg3, arg6, arg7, arg2, arg8)
		err5 := cmd4.Start()
		if err5 != nil {
                	log.Fatal(err5)
			log.Printf("Downloading failed")
			os.Exit(1)			
        		}
		err5 = cmd4.Wait()		
		fmt.Println("Downloading Successfull")

		fmt.Println("Checking for kubectl ...")
		_, err6 := exec.LookPath("kubectl")
		if err6 != nil {
			log.Fatal(err6)
			log.Printf("Kubectl is not installed.")
			os.Exit(1)
		}
		fmt.Printf("Installing kubectl ...")		
		cmd6 := exec.Command(arg0, arg1, arg9, arg13, arg3, arg4, arg5, arg9, arg3, arg6, arg7, arg9, arg8)
		err7 := cmd6.Start()
		if err7 != nil {
                	log.Fatal(err7)
			log.Printf("Unable to Untar the file")
        		}
		err7 = cmd6.Wait()				
		fmt.Println("Setting path ...")
		
		
		cmd7 := exec.Command(arg2,arg14)
		err8 := cmd7.Start()
		if err8 != nil {
                	log.Fatal(err8)
			log.Printf("Failed to start minikube")
			os.Exit(1)
        		}				
		fmt.Println("minikube is installed and set for use")			
		}        										
	},
}

func init() {
	RootCmd.AddCommand(minikubeCmd)		
}
