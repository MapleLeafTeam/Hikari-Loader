/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {
	var id string
	var typr string
	var description string
	var status string
	var playback_link string
	// 参数分别是：标志结果、长选项、短选项、默认值、描述
	addCmd.Flags().StringVarP(&id, "id", "i", "new_id", "desc")
	addCmd.Flags().StringVarP(&typr, "id", "i", "new_id", "desc")
	addCmd.Flags().StringVarP(&description, "id", "i", "new_id", "desc")
	addCmd.Flags().StringVarP(&status, "id", "i", "new_id", "desc")
	addCmd.Flags().StringVarP(&playback_link, "id", "i", "new_id", "desc")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
