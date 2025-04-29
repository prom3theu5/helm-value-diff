package cmd

import (
	"fmt"
	"os"

	"github.com/prom3theu5/helm-values-diff/internal/diff"
	"github.com/prom3theu5/helm-values-diff/internal/meta"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "helm-values-diff",
	Short:   "Generate minimal Helm values overrides",
	Example: "helm-values-diff base.yaml changed.yaml",
	Version: meta.Version,
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return diff.Run(args[0], args[1])
	},
}

func Execute() {
	rootCmd.SetVersionTemplate(fmt.Sprintf("helm-values-diff version %s (commit %s, built at %s)\n",
		meta.Version, meta.Commit, meta.Date))
	
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
