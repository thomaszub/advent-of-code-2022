package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thomaszub/advent-of-code-2022/internal/day2"
)

var day2Cmd = &cobra.Command{
	Use:   "day2",
	Short: "Execute the solution for day 2",
	Long:  "Execute the solution for day 2",
	Run: func(cmd *cobra.Command, args []string) {
		day2.Execute()
	},
}

func init() {
	rootCmd.AddCommand(day2Cmd)
}
