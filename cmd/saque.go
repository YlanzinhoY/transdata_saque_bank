/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ylanzinhoy/transdata_saque_bank/controller"
)

// saqueCmd represents the saque command
var saqueCmd = &cobra.Command{
	Use:   "saque",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var valor int
		fmt.Println("quanto deseja sacar?")
		fmt.Scan(&valor)
		controller.Controller(valor)
	},
}

func init() {
	rootCmd.AddCommand(saqueCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// saqueCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// saqueCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
