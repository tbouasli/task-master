package main

import (
	"log"
	"task-manager/packages/cli/commands"
	"task-manager/packages/database"

	"github.com/spf13/cobra"
)

func Execute() {
	db := database.Connect()
	database.Migrate(db)

	var rootCmd = &cobra.Command{Use: "task-manager"}

	log.Println("Starting Task Manager CLI")

	rootCmd.AddCommand(
		commands.ListTasksCmd(),
		commands.CreateTaskCmd(),
		commands.StartTaskCmd(),
		commands.ConcludeTaskCmd(),
	)

	rootCmd.Execute()
}

func main() {
	Execute()
}
