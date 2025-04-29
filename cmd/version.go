package cmd

import (
	"fmt"
	
	"github.com/prom3theu5/helm-values-diff/internal/meta"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("helm-values-diff version %s (commit %s, built at %s)\n", meta.Version, meta.Commit, meta.Date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
