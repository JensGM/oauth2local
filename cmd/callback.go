package cmd
import (
	"bufio"
	"fmt"
	"os"

	"github.com/equinor/oauth2local/ipc"
	"github.com/spf13/cobra"
)

// callbackCmd represents the callback command
var callbackCmd = &cobra.Command{
	Use:   "callback",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli, err := ipc.NewClient()
		if err != nil {

			fmt.Println("Error...", err)
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			return
		}

		if len(args) != 1 {
			fmt.Println("Error args", args)
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			return
		}

		err = cli.SendCallback(args[0])
		if err != nil {

			fmt.Println("Error...", err)
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			return
		}

		fmt.Println("sent calback", args)
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	},
}

func init() {
	rootCmd.AddCommand(callbackCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// callbackCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// callbackCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
