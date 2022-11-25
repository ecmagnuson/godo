package cmd

import (
	"fmt"
	"godo/utils"
	"strings"

	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//ls lists out all of the items in todo.db still left to do.
func ls() {
	db, err := gorm.Open(sqlite.Open(utils.TodoDir()), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	var tasks []Task
	//SELECT * FROM `tasks` WHERE `todo` = 1
	db.Where("todo", 1).Find(&tasks)

	for _, task := range tasks {
		if task.Priority == "+p3" {
			fmt.Println(task.ID, task.Task, task.Location)
		} else {
			fmt.Println(task.ID, task.Task, task.Location, task.Priority)
		}
	}
}

//lsLocation lists out all items with a specific location still left todo
func lsLocation(location string) {

	db, err := gorm.Open(sqlite.Open(utils.TodoDir()), &gorm.Config{
		// logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	var tasks []Task
	//SELECT * FROM `tasks` WHERE location = `location` AND todo = 1
	db.Where("location = ? AND todo = ?", location, 1).Find(&tasks)

	for _, task := range tasks {
		if task.Priority == "+p3" {
			fmt.Println(task.ID, task.Task, task.Location)
		} else {
			fmt.Println(task.ID, task.Task, task.Location, task.Priority)
		}
	}
}

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list out the items in the todo database",
	Long:  "list out the items in the todo database",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			ls()
		} else if len(args) == 1 {
			lsLocation(strings.Join(args, " "))
		} else {
			panic("no args for entire list and 1 arg for specific location.")
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
