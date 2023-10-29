package cli

import (
	"fmt"
	"os"

	utils "github.com/anirudh97/GollabEdit/pkg"
	"github.com/spf13/cobra"
)

type userCredentials struct {
	username string
	password string
}

type newAccount struct {
	username string
	password string
	email    string
}

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

// Creates the create account command
func createAccountCommand() *cobra.Command {
	var newAcc newAccount

	createAccountCmd := &cobra.Command{
		Use:   "create",
		Short: "Create an account with by passing username, password and email",
		Run: func(cmd *cobra.Command, args []string) {
			if newAcc.username == "" {
				fmt.Println("Please provide an username using --username flag.")
				return
			}
			if newAcc.password == "" {
				fmt.Println("Please provide password using --pwd flag.")
				return
			}
			if newAcc.email == "" {
				fmt.Println("Please provide email using --email flag.")
				return
			}

			if utils.ValidateEmail(newAcc.email) == false {
				fmt.Printf("%s is not an email format", newAcc.email)
				return
			}

		},
	}

	createAccountCmd.Flags().StringVar(&newAcc.username, "username", "", "Username")
	createAccountCmd.Flags().StringVar(&newAcc.password, "pwd", "", "Password")
	createAccountCmd.Flags().StringVar(&newAcc.email, "email", "", "Email")

	return createAccountCmd
}

// Creates the login command
func createLoginCommand() (*cobra.Command, userCredentials) {
	var user userCredentials

	loginCmd := &cobra.Command{
		Use:   "login",
		Short: "Authenticate with username and password",
		Run: func(cmd *cobra.Command, args []string) {
			if user.username == "" {
				fmt.Println("Please provide an username using --username flag.")
				return
			}
			if user.password == "" {
				fmt.Println("Please provide password using --pwd flag.")
				return
			}
			os.Setenv("GOLLABEDIT_USERNAME", user.username)
			os.Setenv("GOLLABEDIT_PWD", user.password)
		},
	}

	loginCmd.Flags().StringVar(&user.username, "username", "", "Username")
	loginCmd.Flags().StringVar(&user.password, "pwd", "", "Password")

	return loginCmd, user

}

// Creates the logout command
func createLogoutCommand() *cobra.Command {
	logoutCmd := &cobra.Command{
		Use:   "logout",
		Short: "Authenticate with username and password",
		Run: func(cmd *cobra.Command, args []string) {
			os.Unsetenv("GOLLABEDIT_USERNAME")
			os.Unsetenv("GOLLABEDIT_PWD")
		},
	}

	return logoutCmd
}
