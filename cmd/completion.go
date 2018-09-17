package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion <bash|zsh>",
	Short: "Generate shell completion scripts",
	Long: `To load completion run
	
. < (cobra-demo completion bash)

Or put in zshrc
echo "source <(cobra-demo completion zsh)" >> ~/.zshrc
`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		shell := args[0]
		switch shell {
		case "zsh":
			err = rootCmd.GenZshCompletion(os.Stdout)
		case "bash":
			err = rootCmd.GenBashCompletion(os.Stdout)
		default:
			err = fmt.Errorf("unknown shell type `%s`", shell)
		}
		return err
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
