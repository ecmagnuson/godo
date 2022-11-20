package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

//make the zero value useful

//Tasks are the basic item in this program.
type Task struct {
	ID       int    // unique ID given to the Task
	task     string // "finish my CS hw before tonight"
	context  string // "@school"
	priority string // "p1", "p2", or "p3" (high to low) priority
	todo     bool   // true if still left to do, else false
}

//String is toString override of todoItem object
func (t Task) String() string {
	return fmt.Sprintf("%d %s %s %s", t.ID, t.task, t.context, t.priority)
}

//getTask returns a Task from a string
func getTask(argument string) Task {
	task := argument[:strings.IndexByte(argument, '@')-1] //collect everything before '@' char
	//splits the white space between '@' context and '+' priority
	contextPlusPriority := strings.Fields(argument[strings.IndexByte(argument, '@'):])
	context, priority := contextPlusPriority[0], contextPlusPriority[1]

	return Task{ID: 0, task: task, context: context, priority: priority, todo: true}
}

argument := "do something @home +1"
todo1 := getTask(argument)
fmt.Println(todo1)
// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a todo item to the database",
	Long:  "add a todo item to the database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
