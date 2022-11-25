package cmd

import (
	"fmt"
	"godo/utils"

	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//lsAllLocations lists out all of the locations currently in use in the database
func lsAllLocations() {
	db, err := gorm.Open(sqlite.Open(utils.TodoDir()), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	var tasks []Task

	//SELECT DISTINCT `location` FROM `tasks` WHERE `todo` = 1 ORDER BY location asc
	db.Distinct("location").Order("location asc").Where("todo", 1).Find(&tasks)

	for _, task := range tasks {
		fmt.Println(task.Location)
	}
}

// lscCmd represents the lsc command
var lscCmd = &cobra.Command{
	Use:   "lsc",
	Short: "list all locations currently in use",
	Long:  "list all locations currently in use",
	Run: func(cmd *cobra.Command, args []string) {
		lsAllLocations()
	},
}

func init() {
	rootCmd.AddCommand(lscCmd)
}
