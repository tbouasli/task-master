package commands

import (
	"task-manager/packages/api/features"

	"github.com/spf13/cobra"
)

func StartTaskCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "start-task",
		Short: "Start a task",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Println("You must provide a task ID")
				return
			}

			features.StartTask(args[0])
			cmd.Println("Task started")
		},
	}

	return cmd
}
