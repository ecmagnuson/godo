package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func lsAllLocations() {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}

	var tasks []Task

	//SELECT DISTINCT `location` FROM `tasks` ORDER BY location asc
	db.Distinct("location").Order("location asc").Find(&tasks)

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
