package cmd

import (
	"godo/utils"
	"strconv"

	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func do(ids []int) {
	db, err := gorm.Open(sqlite.Open(utils.TodoDir()), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	var tasks []Task
	//UPDATE `tasks` SET `todo`=0 WHERE `id` IN (ids)
	db.Where("id", ids).Find(&tasks)
	db.Model(&tasks).Update("todo", 0)
}

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "`do` a task by entering its index.",
	Long:  "`do` a task by entering its index.",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, _ := strconv.Atoi(arg)
			ids = append(ids, id)
		}
		do(ids)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
