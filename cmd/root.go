package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"

	"github.com/ethrai/todo/internal/todo"
)

var (
	cfgFile   string
	storeFile string
	list      *todo.List
)

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

func init() {
	cobra.OnInitialize(initStore)
	rootCmd.PersistentFlags().
		StringVarP(&storeFile, "file", "f", "", "store file (default is $HOME/.tododata.json)")
}

func initStore() {
	if storeFile == "" {
		userHome, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		storeFile = path.Join(userHome, ".tododata.json")
	}

	list = todo.New(storeFile)
	if err := list.Load(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
