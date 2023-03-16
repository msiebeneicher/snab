// generating docs for all commands
// see https://github.com/spf13/cobra/blob/main/doc/README.md for further information
package origin

import (
	"path/filepath"
	"snab/pkg/logger"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var DocsCmd = &cobra.Command{
	Use:   "docs",
	Short: "generate docs",
	Long:  `generate docs`,
}

var DocsGenCmd *cobra.Command

// InitDocsGenCmd get the rootCmd and init a cobra.Command
func InitDocsGenCmd(rootCmd *cobra.Command) *cobra.Command {
	DocsGenCmd = &cobra.Command{
		Use:     "generate",
		Aliases: []string{"gen"},
		Short:   "Generate markdown docs",
		Long:    `Generating markdown docs for all commands`,
		Args:    cobra.ExactArgs(1),
		Example: `app snab docs generate /tmp`,
		Run: func(cmd *cobra.Command, args []string) {
			p, err := filepath.Abs(args[0])
			if err != nil {
				logger.WithField("err", err).Fatalf("error get filepath for`%s`: %s\n", args[0], p)
			}

			err = doc.GenMarkdownTree(rootCmd, p)
			if err != nil {
				logger.Fatal(err)
			}
		},
	}

	return DocsGenCmd
}
