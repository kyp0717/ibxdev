/******

*****/
package cmd

import (
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	ib "ibxdev/webapi"
)

// Get list of accounts

var AccountCmd = &cobra.Command{
	Use:   "account",
	Short: "retrieve list of IB accounts",
	Long:  `A longer description that spans mult`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var getAcctListCmd = &cobra.Command{
	Use:   "list",
	Short: "retrieve list of IB accounts",
	Long:  `A longer description that spans mult`,
	Run: func(cmd *cobra.Command, args []string) {
		ib.GetAcctList(ib.UrlAcctList)

	},
}

// Get an object containing PnL for selected account
var getPnLCmd = &cobra.Command{
	Use:   "pnl",
	Short: "retrieve pnl",
	Long:  `A longer description that spans mult`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("call pnl cmd:  ibx accounts pnl")

	},
}

func init() {
	rootCmd.AddCommand(AccountCmd)
	AccountCmd.AddCommand(getAcctListCmd)
}