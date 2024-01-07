package cmd

import (
  "github.com/bovem/brag/utils"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the bragging document directory specified by $BRAG_DOCS_LOC",
	Long: `Creates the bragging document directory with a subdirectory for the current
year. 

Requires BRAG_DOCS_LOC environment variable.
$ export BRAG_DOCS_LOC=XXXXXXXXXXXXXXXXXXXX

For example:
If BRAG_DOCS_LOC=/work-docs then init command will create /work-docs directory and 
/work-docs/2024/ subdirectory.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.Initialize()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
