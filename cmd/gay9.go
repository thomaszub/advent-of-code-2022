package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thomaszub/advent-of-code-2022/internal/day9"
)

var day9Cmd = &cobra.Command{
	Use:   "day9",
	Short: "Execute the solution for day 9",
	Long:  "Execute the solution for day 9",
	RunE: func(cmd *cobra.Command, args []string) error {
		return day9.Execute()
	},
}

func init() {
	rootCmd.AddCommand(day9Cmd)
}
