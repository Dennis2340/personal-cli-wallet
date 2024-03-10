/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// transferCmd represents the transfer command
var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Transfer Sol to another wallet",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Please provide a single argument representing the private key as a byte array.")
			return
		}
		fmt.Println("Sender Private key: " + args[0])
		fmt.Println("Recepient address: " + args[1])
		fmt.Println("Amount to be sent: " + args[2])
		privateKey := []byte(args[0])
		amount, _ := strconv.ParseUint(args[2], 10, 64)
		txhash, _ := Transfer(privateKey, args[1], amount)
		fmt.Println("Transaction complete.\nTransaction hash: " + txhash)
	},
}

func init() {
	rootCmd.AddCommand(transferCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// transferCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// transferCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
