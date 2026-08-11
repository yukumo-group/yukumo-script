package win

import (
	"github.com/yukumo-group/yukumo-script/generator/tasks/singlesentence"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// tasksCMD contains subcommands that manage the tasks
var tasksCMD = &cobra.Command{
	Use:   "tasks",
	Short: "tasks cmd contains subcommands that manage the tasks",
	Long:  "tasks cmd contains subcommands that manage the tasks",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// showAllSingleSentenceTasksCMD shows all the tasks
var showAllSingleSentenceTasksCMD = &cobra.Command{
	Use:   "showAllSingleSentenceTasks",
	Short: "showAlllSingleSentenceTasks shows all the available tasks",
	Long:  "showAlllSingleSentenceTasks shows all the available tasks",
	Run: func(cmd *cobra.Command, args []string) {
		// Define the format of the texts
		title := color.New(color.FgGreen).Add(color.Bold)
		text := color.New(color.Italic)
		// Show all tasks
		allTasks := singlesentence.Manager.GetAllTasks()
		title.Println("Here are all the created tasks")
		for title := range allTasks {
			text.Println(title)
		}
	},
}
