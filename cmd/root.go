package cmd

import (
	"os"
  "github.com/bovem/brag/utils"
	"github.com/spf13/cobra"
)

var bragComment string

var rootCmd = &cobra.Command{
	Use:   "brag",
	Short: "A command line tool to journal and build a bragging document.",
	Long: `A bragging or hype document is created to keep a record of achievements over time.

Inspired from Julia Evan's blog "Get your work recognized: write a brag document" 
(https://jvns.ca/blog/brag-documents/) this tool could be used to keep a record 
of your brags in Markdown Format that could be referred later to create a formatted 
bragging document.

You can add your brags using command
brag -c "YOUR TEXT HERE"

If BRAG_DOCS_REPO_SYNC environment variable is set to true then changes to bragging document
will be pushed to Git remote.

and review them later using subcommand "about"
brag about last-week`,
	Run: func(cmd *cobra.Command, args []string) { 
    utils.AddBrag(bragComment)
  },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&bragComment, "comment", "c", "", "Bragging Comment")
}
