package cmd

import (
	"godo/utils"
	"strconv"
	"time"

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
	//UPDATE `tasks` SET `completed`="2022-11-24 22:50:48.546",`todo`=false WHERE `id` = 9
	db.Where("id", ids).Find(&tasks)
	db.Model(&tasks).Select("Todo", "Completed").Updates(Task{
		Todo:      false,
		Completed: time.Now(),
	})
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
