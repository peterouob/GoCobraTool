package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

var cusArgsCheckCmd = &cobra.Command{
	Use: "cusargs",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("至少輸入一個參數")
		}
		if len(args) > 2 {
			return errors.New("最多兩個參數")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run cusArgs cmd begin")
		fmt.Println(args)
		fmt.Println("run cusArgs cmd end")
	},
}

// 默認參數設定
var argsCheckCmd = &cobra.Command{
	Use:  "args",
	Args: cobra.OnlyValidArgs,
	// 規定範圍：3選1
	ValidArgs: []string{"123", "abc", "nick"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run cusArgs cmd begin")
		fmt.Println(args)
		fmt.Println("run cusArgs cmd end")
	},
}

func init() {
	rootCmd.AddCommand(cusArgsCheckCmd)
	rootCmd.AddCommand(argsCheckCmd)
}
