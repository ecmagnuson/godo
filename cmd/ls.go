package cmd

import (
	"fmt"
	"godo/utils"
	"strings"

	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//ls lists out all items with a specific location still left todo
//list out anything in db
func ls(str string) {

	db, err := gorm.Open(sqlite.Open(utils.TodoDBPath()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}

	var tasks []Task
	if len(str) > 0 { //list all from location or project
		if strings.Contains(str, "+") {
			//SELECT * FROM `tasks` WHERE project = str AND todo = 1 ORDER BY priority desc
			db.Order("priority desc").Where("project = ? AND todo = ?", str, 1).Find(&tasks)
		} else {
			//SELECT * FROM `tasks` WHERE location = str AND todo = 1 ORDER BY priority desc
			db.Order("priority desc").Where("location = ? AND todo = ?", str, 1).Find(&tasks)
		}
	} else { //list all
		//SELECT * FROM `tasks` WHERE `todo` = 1 ORDER BY priority desc
		db.Order("priority desc").Where("todo", 1).Find(&tasks)
	}

	/* 	db.Order("priority desc").Find(&tasks)
	   	// SELECT * FROM users ORDER BY age desc, name; */

	for _, task := range tasks {

		strTask := fmt.Sprintf("%d %s %s %s", task.ID, task.Priority, task.Task, task.Location)

		if task.Project != "+" {
			strTask = fmt.Sprintf(strTask+" %s", task.Project)
		}
		if task.Priority == "!" {
			colorReset := "\033[0m"
			colorBlue := "\033[34m"
			fmt.Println(string(colorBlue) + strTask + string(colorReset))
		} else {
			fmt.Println(strTask)
		}
	}
}

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list out the items in the todo database",
	Long:  "list out the items in the todo database",
	Run: func(cmd *cobra.Command, args []string) {
		ls(strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
