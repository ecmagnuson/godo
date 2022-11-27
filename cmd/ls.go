package cmd

import (
	"fmt"
	"godo/utils"
	"strings"

	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//lsLocation lists out all items with a specific location still left todo
func lsLocation(location string) {

	db, err := gorm.Open(sqlite.Open(utils.TodoDBPath()), &gorm.Config{
		// logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	var tasks []Task
	if len(location) > 0 { //list all from one location
		//SELECT * FROM `tasks` WHERE location = `location` AND todo = 1
		db.Where("location = ? AND todo = ?", location, 1).Find(&tasks)
	} else { //list all
		//SELECT * FROM `tasks` WHERE `todo` = 1
		db.Where("todo", 1).Find(&tasks)
	}

	for _, task := range tasks {
		if task.Project == "+" {
			fmt.Println(task.ID, task.Priority, task.Task, task.Location)
		} else {
			fmt.Println(task.ID, task.Priority, task.Task, task.Location, task.Project)
		}
	}
}

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list out the items in the todo database",
	Long:  "list out the items in the todo database",
	Run: func(cmd *cobra.Command, args []string) {
		lsLocation(strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
