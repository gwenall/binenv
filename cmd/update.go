package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/devopsworks/tools/binenv/internal/app"
)

// localCmd represents the local command
func updateCmd() *cobra.Command {
	app := app.New()
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update software definitions",
		Long:  `Updates installable software`,
		// Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				return app.Update(args[0])
			}
			return app.Update("")
		},
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			// Remove already selected distributions from completion
			list := app.GetPackagesListWithPrefix(toComplete)
			list = removeFromSlice(list, args)
			return list, cobra.ShellCompDirectiveNoFileComp
		},
	}

	// cmd.Flags().IntVarP(&a.Params.MinLength, "min-length", "m", 16, "Specify minimum password length, must not be less than 8")
	return cmd
}

func removeFromSlice(orig, remove []string) []string {
	for i := 0; i < len(orig); i++ {
		url := orig[i]
		for _, rem := range remove {
			if url == rem {
				orig = append(orig[:i], orig[i+1:]...)
				i--
				break
			}
		}
	}

	return orig
}
