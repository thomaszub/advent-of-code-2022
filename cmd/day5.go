package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thomaszub/advent-of-code-2022/internal/day5"
)

var day5Cmd = &cobra.Command{
	Use:   "day5",
	Short: "Execute the solution for day 5",
	Long:  "Execute the solution for day 5",
	RunE: func(cmd *cobra.Command, args []string) error {
		return day5.Execute()
	},
}

func init() {
	rootCmd.AddCommand(day5Cmd)
}
