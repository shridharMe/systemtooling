/*
Copyright © 2020 SHRIDHAR PATIL shridhar.patil01@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
"os/exec"
	"github.com/spf13/cobra"

)
var (
	  action string
)
var terraCmd = &cobra.Command{
	Use:   "terra",
	Short: "run terraform commands",
	Run: func(cmd *cobra.Command, args []string) {	 
	 	if out, err := exec.Command("terraform",  action).Output(); err != nil {
			fmt.Println("error ", err)
		}else{
			
			fmt.Println("Output" , string(out))
		}
		
	},
}

func init() {
	terraCmd.Flags().StringVarP(&action, "action", "a", "", "terraform command")
	
}