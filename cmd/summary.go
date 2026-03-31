/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/RevDau-PallaviShinde/Go-Cli-App/handlers"
	"github.com/spf13/cobra"
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "my-task-summary",
	Short: "my task summary",
	Long:  `app summary`,
	Run: func(cmd *cobra.Command, args []string) {

		total, completed, pending, err := handlers.GetSummary()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Task Summary")
		fmt.Println("----------------")
		fmt.Println("Total Tasks     :", total)
		fmt.Println("Completed Tasks :", completed)
		fmt.Println("Pending Tasks   :", pending)

	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// summaryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// summaryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
