/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/gucchi0421/goya/internal/newpass"
	"github.com/spf13/cobra"
)

var (
	length int
	count int
)

// newpassCmd represents the newpass command
var newpassCmd = &cobra.Command{
	Use:   "newpass",
	Short: "ランダムなパスワードを生成します",
	Run: func(cmd *cobra.Command, args []string) {
		if  err := newpass.Generate(length, count);err != nil {
			fmt.Printf("error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	newpassCmd.Flags().IntVarP(&length, "len", "l", 12, "パスワードの長さを指定できます")
	newpassCmd.Flags().IntVarP(&count, "count", "c", 1, "生成するパスワードの数を指定できます")
	rootCmd.AddCommand(newpassCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newpassCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newpassCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
