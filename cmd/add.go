/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"net/url"
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
		targetUrl := "http://127.0.0.1:8080/plugins/anime"

		payload := url.Values{"key":{"value"}, "id": {"123"}}

		response, err := http.PostForm(targetUrl, payload)

		if err != nil:
			
	},
}

func init() {
	addCmd.PersistentFlags().String("id", "", "write a id for this")
	addCmd.PersistentFlags().String("typr", "", "write a id for this")
	addCmd.PersistentFlags().String("description", "", "write a id for this")
	addCmd.PersistentFlags().String("status", "", "write a id for this")
	addCmd.PersistentFlags().String("playback_link", "", "write a id for this")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
