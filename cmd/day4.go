package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thomaszub/advent-of-code-2022/internal/day4"
)

var day4Cmd = &cobra.Command{
	Use:   "day4",
	Short: "Execute the solution for day 4",
	Long:  "Execute the solution for day 4",
	RunE: func(cmd *cobra.Command, args []string) error {
		return day4.Execute()
	},
}

func init() {
	rootCmd.AddCommand(day4Cmd)
}
