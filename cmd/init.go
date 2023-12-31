/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "install the Hikari-CMS",
	Long:  `use this command to install the Hikari-CMS in your server`,
	Run: func(cmd *cobra.Command, args []string) {
		command1 := exec.Command("/bin/bash", "./script/install.sh")
		err1 := command1.Run()
		if err1 != nil {
			log.Fatalf("command1.Run() failed with %s\n", err1)
		}
		fmt.Printf("clone the Core susesfully")

	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
