package cmd

import (
	"fmt"
	"godo/utils"
	"log"

	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//lsAllLocations lists out all of the locations currently in use in the database
func lsAllLocations(args []string) {
	db, err := gorm.Open(sqlite.Open(utils.TodoDBPath()), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	context := args[0] // search by "@" or by "+"
	var tasks []Task
	if context == "@" {
		//SELECT DISTINCT `location` FROM `tasks` WHERE `todo` = 1 ORDER BY location asc
		db.Distinct("location").Order("location asc").Where("todo", 1).Find(&tasks)
		for _, task := range tasks {
			fmt.Println(task.Location)
		}
	} else if context == "+" {
		//SELECT DISTINCT `project` FROM `tasks` WHERE `todo` = 1 ORDER BY project asc
		db.Distinct("project").Order("project asc").Where("todo", 1).Find(&tasks)
		for _, task := range tasks {
			if task.Project == "+" {
				continue
			}
			fmt.Println(task.Project)
		}
	}
}

// allCmd represents the all subcommand
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "list all locations or projects currently in use",
	Long:  "list all locations currently in use",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalf("implementing listing all later.")
			//list out all combinations of @ and + togethor
		}
		lsAllLocations(args)
	},
}

//all is a subcommand of ls
func init() {
	lsCmd.AddCommand(allCmd)
}
