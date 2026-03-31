/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/RevDau-PallaviShinde/Go-Cli-App/handlers"
	"github.com/spf13/cobra"
)

var keyword string

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "pallavi-search",
	Short: "Search tasks",
	Long:  `Search tasks based on keywords in title or description.`,
	Run: func(cmd *cobra.Command, args []string) {

		tasks, err := handlers.SearchTasks(keyword)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("No matching tasks found.")
			return
		}

		for _, task := range tasks {
			fmt.Println("ID:", task.ID)
			fmt.Println("Title:", task.Title)
			fmt.Println("Description:", task.Description)
			fmt.Println("Status:", task.Status)
			fmt.Println("Priority:", task.Priority)
			fmt.Println("---------------------------")
		}

	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringVar(&keyword, "keyword", "", "Keyword to search (required)")
	_ = searchCmd.MarkFlagRequired("keyword")

}
