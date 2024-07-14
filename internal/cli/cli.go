package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "easysub",
	Short: "get subs",
	Long:  "",
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return fmt.Errorf(":%w", err)
	}

	return nil
}
