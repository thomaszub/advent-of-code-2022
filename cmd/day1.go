package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thomaszub/advent-of-code-2022/internal/day1"
)

var day1Cmd = &cobra.Command{
	Use:   "day1",
	Short: "Execute the solution for day 1",
	Long:  "Execute the solution for day 1",
	Run: func(cmd *cobra.Command, args []string) {
		day1.Execute()
	},
}

func init() {
	rootCmd.AddCommand(day1Cmd)
}
