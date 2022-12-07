package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thomaszub/advent-of-code-2022/internal/day7"
)

var day7Cmd = &cobra.Command{
	Use:   "day7",
	Short: "Execute the solution for day 7",
	Long:  "Execute the solution for day 7",
	RunE: func(cmd *cobra.Command, args []string) error {
		return day7.Execute()
	},
}

func init() {
	rootCmd.AddCommand(day7Cmd)
}
