package cmd

import (
  "github.com/bovem/brag/utils"
	"github.com/spf13/cobra"
  "strings"
)

var model string
var prompt string

var summarizeCmd = &cobra.Command{
	Use:   "summarize",
	Short: "Summarizes the brags from the specified time period in a draft bragging document.",
	Long: `Summarizes brags from the specified time period using Large Language Models hosted on Ollama (https://ollama.ai/)
and generates a draft bragging document.

Requires following environment variable (assuming Ollama is already deployed and the model is downloaded).
$ export OLLAMA_URL="http://localhost:11434"

For Example:
brag summarize --model llama2:latest last-month
brag summarize --model llama2:latest 2-weeks

The time period has to be specified in any of the following formats
* <numeric-time>-<range-of-days> (brag about 2-months/brag about 3-years)
* today/yesterday/last-week/last-month (brag about last-week/brag about yesterday)
`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.Summarize(strings.Join(args, " "), model, prompt)
	},
}

func init() {
	rootCmd.AddCommand(summarizeCmd)
  summarizeCmd.PersistentFlags().StringVarP(&model, "model", "m", "llama2:latest", "LLM Model to be used while generating the bragging document")
  summarizeCmd.PersistentFlags().StringVarP(&prompt, "prompt", "p", utils.BragDocTemplate, "LLM Prompt for generating the bragging document")
}
