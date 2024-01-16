package cmd

import (
  "github.com/bovem/brag/utils"
	"github.com/spf13/cobra"
  "strings"
  "fmt"
)

var timeFrame string

var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "Describes the brags from the specified time period.",
	Long: `Presents brags from the specified time period with date and time. For Example:
brag about today
brag about last-week
brag about 2-months

The time period has to be specified in any of the following formats
* <numeric-time>-<range-of-days> (brag about 2-months/brag about 3-years)
* today/yesterday/last-week/last-month (brag about last-week/brag about yesterday)
`,
	Run: func(cmd *cobra.Command, args []string) {
    brags := utils.Bragging(strings.Join(args, " "))
    fmt.Println(brags)
	},
}

func init() {
	rootCmd.AddCommand(aboutCmd)
	aboutCmd.PersistentFlags().StringVarP(&timeFrame, "timeframe", "t", "today", "Time Frame for Bragging")
}
