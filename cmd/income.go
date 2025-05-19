/*
Copyright © 2025 bkmnst <bakemonsuta@gmail.com>
*/
package cmd

import (
	"fmt"
	godget "godget/pkg"

	"github.com/spf13/cobra"
)

// incomeCmd represents the income command
var incomeCmd = &cobra.Command{
	Use:   "income",
	Short: "Add income amount",
	Long:  `Add income to your transaction history.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		amount := args[0]
		godget.AddTransaction("+" + amount)
		fmt.Println("income called")
	},
}

func init() {
	rootCmd.AddCommand(incomeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// incomeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// incomeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
