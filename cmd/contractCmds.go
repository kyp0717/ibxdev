/******

*****/
package cmd

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"ibxdev/webapi"
)

var ContractCmd = &cobra.Command{
	Use:   "contract",
	Short: "request contract for instruments",
	Long:  `A longer description that spans mult`,
	Run: func(cmd *cobra.Command, args []string) {
		stock, _ := cmd.Flags().GetString("stock")
		if stock != "" {
			data, _ := webapi.SearchSymbol(stock)
			fmt.Print(string(data))
		}
		conid, _ := cmd.Flags().GetString("conid")
		if conid != "" {
			data := webapi.GetCtxInfo(conid)
			spew.Dump(data)
		}

	},
}

func init() {
	RootCmd.AddCommand(ContractCmd)
	ContractCmd.Flags().StringP("stock", "s", "", "request stock price")
	ContractCmd.Flags().StringP("conid", "c", "", "request instrument by contract id")

}
