/*
Copyright © 2025 bkmnst <bakemonsuta@gmail.com>
*/
package cmd

import (
	"fmt"
	godget "godget/pkg"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Shows recent transactions",
	Long:  `Shows a list of transactions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(godget.ListTransactions())
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
