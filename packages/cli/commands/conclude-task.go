package commands

import (
	"task-manager/packages/api/features"

	"github.com/spf13/cobra"
)

func ConcludeTaskCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "conclude-task",
		Short: "Conclude a task",
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) == 0 {
				cmd.Println("You must provide a task ID")
				return
			}

			features.ConcludeTask(args[0])
			cmd.Println("Task concluded")
		},
	}

	return cmd
}
