/*
Copyright © 2025 bkmnst <bakemonsuta@gmail.com>
*/
package cmd

import (
	"fmt"
	godget "godget/pkg"

	"github.com/spf13/cobra"
)

// balanceCmd represents the balance command
var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Shows current balance",
	Long:  `Displays current balance, read from osmewhere, idk now`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Current balance: %.2f\n", godget.CalculateBalance())
	},
}

func init() {
	rootCmd.AddCommand(balanceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// balanceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// balanceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
