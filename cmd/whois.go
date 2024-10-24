/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/gucchi0421/goya/internal/whois"
	"github.com/spf13/cobra"
)

// whoisCmd represents the dns command
var whoisCmd = &cobra.Command{
	Use:   "whois [domain]",
	Short: "ドメインからWHOIS情報を取得",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		domain := args[0]
		whois.Search(domain)
	},
}

func init() {
	rootCmd.AddCommand(whoisCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dnsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dnsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
