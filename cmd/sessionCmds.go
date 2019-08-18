/******

*****/
package cmd

import (
	/* "fmt"*/
	/*"github.com/davecgh/go-spew/spew"*/
	"github.com/spf13/cobra"
	"ibxdev/webapi"
)

var LogOutCmd = &cobra.Command{
	Use:   "logout",
	Short: "log out of session",
	Long:  `A longer description that spans mult`,
	Run: func(cmd *cobra.Command, args []string) {
		webapi.Logout()
	},
}

var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "log out of session",
	Long:  `A longer description that spans mult`,
	Run: func(cmd *cobra.Command, args []string) {
		webapi.Status()
	},
}

func init() {
	RootCmd.AddCommand(LogOutCmd, StatusCmd)

}
