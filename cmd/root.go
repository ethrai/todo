package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"

	"github.com/ethrai/todo/internal/todo"
)

var list *todo.List

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A small todo app",
	Long: `Todo is a CLI tool for managing tasks and time in terminal.

Usage:
  todo [command]

Available Commands:
  add         Add a task
  rm          Remove a task
  done/undo   Toggle task completion
  list        List task in formatted manner

Flags:
  -f, --file   string    file name where tasks will be stored in json format (default is $HOME/.tododata.json)
      --config string    config file (default is $HOME/.todoconf.json)
  -h, --help             help for todo

Use "todo [command] --help" for more information about a command.
  `,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// TODO: init is kinda messy here, probably should move list initialization
// somewhere else!
func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	userHome, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defaultStorePath := path.Join(userHome, ".tododata.json")

	storePath := rootCmd.Flags().
		StringP("file", "f", defaultStorePath,
			"Filepath where tasks will be stored")

	list = todo.New(*storePath)
	if err := list.Load(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
