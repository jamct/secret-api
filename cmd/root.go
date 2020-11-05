package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "web-cli",
	Short: "A demo for secret management",
	Long: `store secrets`,
}

var cmdLogin = &cobra.Command{
	Use:   "login",
	Short: "Login",
	Long:  `Login`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		login(args[0])
	},
}

func Exec() {
	rootCmd.AddCommand(cmdLogin)
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
