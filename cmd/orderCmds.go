/******

*****/
package cmd

import (
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"ibxdev/webapi"
)

// Get list of accounts

var OrderCmd = &cobra.Command{
	Use:   "order",
	Short: "retrieve list of IB accounts",
	Long:  `A longer description that spans mult`,
	Run: func(cmd *cobra.Command, args []string) {
		data, _ := webapi.PlaceOrder(viper.GetString("file"))
		fmt.Println(string(data))

	},
}

func init() {
	rootCmd.AddCommand(OrderCmd)
	OrderCmd.Flags().StringP("file", "f", "", "specify filename")
	viper.BindPFlag("file", rootCmd.Flags().Lookup("file"))

}
