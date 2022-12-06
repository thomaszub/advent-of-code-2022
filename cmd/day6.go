package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thomaszub/advent-of-code-2022/internal/day6"
)

var day6Cmd = &cobra.Command{
	Use:   "day6",
	Short: "Execute the solution for day 6",
	Long:  "Execute the solution for day 6",
	RunE: func(cmd *cobra.Command, args []string) error {
		return day6.Execute()
	},
}

func init() {
	rootCmd.AddCommand(day6Cmd)
}
