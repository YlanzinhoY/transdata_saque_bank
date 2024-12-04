/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/ylanzinhoy/transdata_saque_bank/controller"
)

// saqueCmd represents the saque command
var saqueCmd = &cobra.Command{
	Use:   "saque",
	Short: "Simples comando de sacar seu dinheiro com a menor quantia de notas",
	Run: func(cmd *cobra.Command, args []string) {
		var valor int
		fmt.Println("Olá Sr.Victor! Quanto deseja sacar?")
		_, err := fmt.Scan(&valor)
		if err != nil || valor < 1 {
			log.Fatal("entrada invalida ou saque pelo menos 1 real")
			return
		}
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
