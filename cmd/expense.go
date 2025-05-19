/*
Copyright © 2025 bkmnst <bakemonsuta@gmail.com>
*/
package cmd

import (
	"fmt"
	godget "godget/pkg"

	"github.com/spf13/cobra"
)

// expenseCmd represents the expense command
var expenseCmd = &cobra.Command{
	Use:   "expense",
	Short: "Add expense amount",
	Long:  `Add expense to your transaction history.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		amount := args[0]
		godget.AddTransaction("-" + amount)
		fmt.Println("expense called")
	},
}

func init() {
	rootCmd.AddCommand(expenseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// expenseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// expenseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
