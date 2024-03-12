package cli

import (
	"github.com/spf13/cobra"
)

func SetupCommands() *cobra.Command {

	rootCmd := &cobra.Command{Use: "gollabedit"}

	createAccountCmd := createAccountCommand()

	loginCmd := createLoginCommand()

	logoutCmd := createLogoutCommand()

	// User auth commands
	rootCmd.AddCommand(createAccountCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(logoutCmd)

	//File commands
	fileCmd := &cobra.Command{
		Use:   "file",
		Short: "File operations",
	}

	// File-related subcommands
	createFileCmd := createFileCommand()
	shareFileCmd := shareFileCommand()
	openFileCmd := openFileCommand()

	// Adding subcommands to fileCmd
	fileCmd.AddCommand(createFileCmd)
	fileCmd.AddCommand(shareFileCmd)
	fileCmd.AddCommand(openFileCmd)

	// Adding fileCmd as a subcommand of rootCmd
	rootCmd.AddCommand(fileCmd)

	return rootCmd
}
