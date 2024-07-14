package cli

import "github.com/spf13/cobra"

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "find all movies files",
	Run: func(cmd *cobra.Command, args []string) {

		var path string

		if len(args) != 0 {
			path = args[0]
		}

		PerformSearch(path)
	},
}

func init() {
	rootCmd.AddCommand(findCmd)
}
