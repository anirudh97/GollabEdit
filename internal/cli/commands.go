package cli

import (
	"github.com/spf13/cobra"
)

func SetupCommands() *cobra.Command {

	rootCmd := &cobra.Command{Use: "gollabedit"}

	createAccountCmd := createAccountCommand()

	loginCmd, _ := createLoginCommand()

	logoutCmd := createLogoutCommand()

	rootCmd.AddCommand(createAccountCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(logoutCmd)

	return rootCmd
}
