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
		command1 := exec.Command("git", "clone", "https://github.com/MapleLeafTeam/Hikari-Core.git")
		err1 := command1.Run()
		if err1 != nil {
			log.Fatalf("command1.Run() failed with %s\n", err1)
		}
		fmt.Printf("clone the Core susesfully")

		command2 := exec.Command("pip", "install", "poetry")
		err2 := command2.Run()
		if err2 != nil {
			log.Fatalf("command2.Run() failed with %s\n", err2)
		}
		fmt.Printf("installed Poetry susesfully")

		command3 := exec.Command("cd", "Hikari-Core")
		err3 := command3.Run()
		if err3 != nil {
			log.Fatalf("command3.Run() failed with %s\n", err3)
		}

		command4 := exec.Command("poetry", "install")
		err4 := command4.Run()
		if err4 != nil {
			log.Fatalf("command2.Run() failed with %s\n", err4)
		}

		command5 := exec.Command("cd Hikari-Core && python init.py && poetry run uvicorn main:app --host 0.0.0.0 --port 8080")
		err5 := command5.Run()
		if err5 != nil {
			log.Fatalf("command5.Run() failed with %s\n", err5)
		}
		fmt.Printf("runing now!")
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
