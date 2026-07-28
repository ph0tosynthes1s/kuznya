package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewBuildCommand команда сборки для cli
func NewBuildCommand() *cobra.Command {
	var configPath string
	var outputPath string

	command := &cobra.Command{
		Use:   "build",
		Short: "Создать проект по YAML-конфигурации",
		Example: `  kuznya build
  kuznya build --config kuznya.yaml
  kuznya build --config templates/go-api.yaml --output ./my-api`,

		Args: cobra.NoArgs,

		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(
				cmd.OutOrStdout(),
				"Кузня запущена\nКонфигурация: %s\nРезультат: %s\n",
				configPath,
				outputPath,
			)

			return nil
		},
	}

	command.Flags().StringVarP(
		&configPath,
		"config",
		"c",
		"kuznya.yaml",
		"путь к YAML-конфигурации",
	)

	command.Flags().StringVarP(
		&outputPath,
		"output",
		"o",
		".",
		"директория для создаваемого проекта",
	)

	return command
}
