package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

//parse string to make Task object
//convert Task to JSON
//Put JSON in sqlite database

//Task is the basic item in this program.
type Task struct {
	ID        int       // unique ID given to the Task
	Task      string    // "finish my CS hw before tonight"
	Location  string    // "@school"
	Priority  string    // "p1", "p2", or "p3" (high to low) priority
	Created   time.Time // time.Now() called when Task created
	Completed time.Time // time.Time{} zero date of time.IsZero() until Todo is false
	Todo      bool      // true if still left to do, else false
}

//ToJSON converts a Task struct to JSON
func (t Task) ToJSON() string {
	b, err := json.Marshal(t)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

//String is toString override of todoItem object
func (t Task) String() string {
	return fmt.Sprintf("%d %s %s %s %t", t.ID, t.Task, t.Location, t.Priority, t.Todo)
}

//WriteFile writes text to a file
func WriteFile(file string, text string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}

//GetTask returns a Task from a string
func GetTask(text string) Task {
	task := text[:strings.IndexByte(text, '@')-1] //collect everything before '@' char
	//splits the white space between '@' context and '+' priority
	contextPlusPriority := strings.Fields(text[strings.IndexByte(text, '@'):])
	context, priority := contextPlusPriority[0], contextPlusPriority[1]

	return Task{
		ID:        0,
		Task:      task,
		Location:  context,
		Priority:  priority,
		Created:   time.Now(),
		Completed: time.Time{}, //invoking zero date of time.IsZero()
		Todo:      true,
	}
}

//ContainsLocation checks for an "@" symbol in the text
func ContainsLocation(text string) bool {
	return strings.Contains(text, "@")
}

//ContainsPriority returns true if the string has a priority
func ContainsPriority(task string) bool {
	var priorities = []string{"+p1", "+p2", "+p3"}
	return slices.Contains(priorities, task)
}

//Format readies the text so that it can be turned into a Task
func Format(task string) string {
	if !ContainsLocation(task) {
		task = fmt.Sprintf("%s @unknown", task)
	}
	if !ContainsPriority(task) {
		task = fmt.Sprintf("%s +p3", task)
	}

	return task
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a todo item to the database",
	Long:  "add a todo item to the database",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			panic("supporting 1 arg for now")
		}

		stringTask := Format(strings.Join(args, " "))
		task := GetTask(stringTask)

		JSONTask := task.ToJSON()
		path := "/home/elliott/.sourcecode/godo/todo.txt"
		WriteFile(path, JSONTask)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
