/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/portto/solana-go-sdk/client/rpc"
	"github.com/spf13/cobra"
)

// createWalletCmd represents the createWallet command
var createWalletCmd = &cobra.Command{
	Use:   "createWallet",
	Short: "Create a new Wallet",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating new Wallet")
		wallet := CreateNewWallet(rpc.DevnetRPCEndpoint)
		fmt.Println("Public Key: " + wallet.account.PublicKey.ToBase58())
		fmt.Println("Public Key saved in 'data' file")
	},
}

func init() {
	rootCmd.AddCommand(createWalletCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createWalletCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createWalletCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
