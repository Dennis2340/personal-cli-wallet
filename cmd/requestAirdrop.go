/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/portto/solana-go-sdk/client/rpc"
	"github.com/spf13/cobra"
)

// requestAirdropCmd represents the requestAirdrop command
var requestAirdropCmd = &cobra.Command{
	Use:   "requestAirdrop",
	Short: "Request airdrop in Solana",
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
		fmt.Println("Private key: " + args[0])
		fmt.Println("Amount to be sent: " + args[1])
		privateKey := []byte(args[0])
		wallet, _ := ImportOldWallet(privateKey, string(rpc.DevnetRPCEndpoint))
		fmt.Println("Requesting airdrop to: " + wallet.account.PublicKey.ToBase58())
		amount, _ := strconv.ParseUint(args[1], 10, 64)
		txhash, _ := RequestAirDrop(privateKey, amount)
		fmt.Println("Airdropped " + strconv.Itoa(int(amount)) + " SOL.\nTransaction hash: " + txhash)
	},
}

func init() {
	rootCmd.AddCommand(requestAirdropCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// requestAirdropCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// requestAirdropCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
