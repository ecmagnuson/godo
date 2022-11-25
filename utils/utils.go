package utils

import (
	"log"
	"os"
	"path/filepath"
)

//TodoDir returns the filepath of the todo.db (OS agnostic) - home/.todo/todo.db
func TodoDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(home, ".todo", "todo.db")
}
