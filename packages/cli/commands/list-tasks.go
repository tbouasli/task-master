package commands

import (
	"task-manager/packages/api/features"

	"github.com/spf13/cobra"
)

func ListTasksCmd() *cobra.Command {
	var input features.ListTasksInput

	var cmd = &cobra.Command{
		Use:   "list-tasks",
		Short: "List tasks",
		Run: func(cmd *cobra.Command, args []string) {
			tasks := features.ListTasks(input)

			for _, task := range tasks {
				cmd.Println(task.ID, task.Name, task.Status)
			}
		},
	}

	cmd.Flags().StringVarP(&input.Status, "status", "s", "", "Task status")

	return cmd
}
