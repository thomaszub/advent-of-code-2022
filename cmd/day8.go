package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thomaszub/advent-of-code-2022/internal/day8"
)

var day8Cmd = &cobra.Command{
	Use:   "day8",
	Short: "Execute the solution for day 8",
	Long:  "Execute the solution for day 8",
	RunE: func(cmd *cobra.Command, args []string) error {
		return day8.Execute()
	},
}

func init() {
	rootCmd.AddCommand(day8Cmd)
}
