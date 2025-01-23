package commands

import (
	"task-manager/packages/api/features"

	"github.com/spf13/cobra"
)

func CreateTaskCmd() *cobra.Command {
	var input features.CreateTaskInput

	var cmd = &cobra.Command{
		Use:   "create-task",
		Short: "Create a task",
		Run: func(cmd *cobra.Command, args []string) {
			features.CreateTask(input)
			cmd.Println("Task created")
		},
	}

	cmd.Flags().StringVarP(&input.Name, "name", "n", "", "Task name")
	cmd.Flags().StringVarP(&input.Description, "description", "d", "", "Task description")

	return cmd
}
