package commands

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// NewRootCommand корневая команда по имени cli-библиотеки
func NewRootCommand() *cobra.Command {
	command := &cobra.Command{
		Use:          "kuznya",
		Short:        "Генератор структуры проектов",
		Long:         `Kuznya создаёт структуру проекта и файлы на основании YAML-конфигурации.`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	command.AddCommand(NewBuildCommand())

	return command
}

func Execute() {
	command := NewRootCommand()

	if err := command.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка:", err)
		os.Exit(1)
	}
}

func ExecuteWithWriter(writer io.Writer) error {
	command := NewRootCommand()
	command.SetOut(writer)
	command.SetErr(writer)

	return command.Execute()
}
