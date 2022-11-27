package cmd

import (
	"bufio"
	"fmt"
	"godo/utils"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//Task is the database type in this program.
type Task struct {
	ID        int       // unique ID given to the Task
	Priority  string    // urgent task prefixed with "!"
	Task      string    // "finish my CS hw before tonight"
	Location  string    // "@school"
	Project   string    // "+p1", "+p2", or "+p3" (high to low) priority
	Created   time.Time // time.Now() called when Task created
	Completed time.Time // time.Time{} zero date of time.IsZero() until Todo is false
	Todo      bool      // true if still left to do, else false
}

//getTask returns a Task from a string
func getTask(text string) Task {
	priority := "0"                               //default priority is nothing showing
	task := text[:strings.IndexByte(text, '@')-1] //collect everything before '@' char

	if strings.HasPrefix(task, "!") {
		priority = "!"
		task = text[strings.IndexByte(text, '!')+2 : strings.IndexByte(text, '@')-1]
	}

	//splits the white space between '@' context and '+' priority
	contextPlusProject := strings.Fields(text[strings.IndexByte(text, '@'):])
	context, project := contextPlusProject[0], contextPlusProject[1]

	return Task{
		ID:        0,
		Task:      task,
		Priority:  priority,
		Location:  context,
		Project:   project,
		Created:   time.Now(),
		Completed: time.Time{}, //invoking zero date of time.IsZero()
		Todo:      true,
	}
}

//containsLocation checks for an "@" symbol in the text
func containsLocation(text string) bool {
	return strings.Contains(text, "@")
}

//containsProject returns true if the string has a project associated with it
func containsProject(text string) bool {
	return strings.Contains(text, "+")
}

//formatString readies the text so that it can be turned into a Task
func formatString(strTask string) string {
	if !containsLocation(strTask) {
		strTask = fmt.Sprintf("%s @unknown", strTask)
	}
	//no need to add project if there is not one given by user
	if !containsProject(strTask) {
		strTask = fmt.Sprintf("%s +", strTask)
	}

	return strTask
}

//insert adds a slice of Tasks to a given DBPath
func insert(task []Task, DBPath string) {

	db, err := gorm.Open(sqlite.Open(DBPath), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&task)
	//db.Create add all tasks to the db
	db.Create(&task)
}

//formatOneTask formats the args after the 'add' command into a Task
func formatOneTask(args []string) []Task {
	stringTask := formatString(strings.Join(args, " "))
	task := getTask(stringTask)
	var tasks = []Task{task}
	return tasks
}

//formatMultipleTasks formats stdin from a user into multiple Tasks
func formatMultipleTasks(args []string) []Task {
	var strTask []string
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		next, _ := reader.ReadString('\n')
		if next == "\n" {
			break
		}
		next = formatString(next)
		strTask = append(strTask, next)
	}

	//convert to task array here
	var tasks []Task
	for _, task := range strTask {
		tasks = append(tasks, getTask(task))
	}
	return tasks
}

//Add adds the incoming args from user to Task after passing it to helper functions
//that format the incoming args into Tasks.
//If no path specified then it will use home/.todo/todo.db
func Add(args []string, optionalDBPath ...string) {
	path := strings.Join(optionalDBPath, " ")
	if len(optionalDBPath) == 0 {
		path = utils.TodoDBPath()
	}

	if len(args) == 0 { //want to add multiple Tasks
		tasks := formatMultipleTasks(args)
		insert(tasks, path)
	} else { //add one Task
		tasks := formatOneTask(args)
		insert(tasks, path)
	}
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a todo item to the database",
	Long:  "add a todo item to the database",
	Run: func(cmd *cobra.Command, args []string) {
		Add(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
