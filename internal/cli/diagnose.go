// Copyright 2026 Pimmetjeoss and contributors. Licensed under Apache-2.0. See LICENSE.
// Novel command scaffold. Implement the RunE body before shipping.
// generate --force preserves implemented bodies; untouched TODO scaffolds may refresh.
// pp:data-source auto
// Supported strategies: auto, local, live, or computed. Change this default deliberately.

package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newNovelDiagnoseCmd(flags *rootFlags) *cobra.Command {
	var flagWorkspace string

	cmd := &cobra.Command{
		Use:         "diagnose",
		Short:       "Resolve a task name to one execution and its log",
		Example:     "  timextender-pp-cli diagnose --workspace finance --dry-run --json",
		Annotations: map[string]string{"mcp:read-only": "false", "pp:data-source": "auto", "pp:novel-scaffold": "true"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if dryRunOK(flags) {
				return writeDryRun(cmd.OutOrStdout(), flags, "diagnose")
			}
			// validate required flags here
			return fmt.Errorf("TODO: implement novel feature %q", "diagnose")
		},
	}
	cmd.Flags().StringVar(&flagWorkspace, "workspace", "", "TODO: describe --workspace")
	return cmd
}
