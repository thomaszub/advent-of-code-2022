package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thomaszub/advent-of-code-2022/internal/day3"
)

var day3Cmd = &cobra.Command{
	Use:   "day3",
	Short: "Execute the solution for day 3",
	Long:  "Execute the solution for day 3",
	RunE: func(cmd *cobra.Command, args []string) error {
		return day3.Execute()
	},
}

func init() {
	rootCmd.AddCommand(day3Cmd)
}
