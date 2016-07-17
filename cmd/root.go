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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string


var RootCmd = &cobra.Command{
	Use:   "Turbo",
	Short: "Simple and Powerfull utility for Docker",
	Long: `Turbo:
  Simple and Powerfull utility for Docker`,
}


func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
		
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.turbo.yaml)")		
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func initConfig() {
	if cfgFile != "" { 
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".turbo") 
	viper.AddConfigPath("$HOME")  
	viper.AutomaticEnv()          

	
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
