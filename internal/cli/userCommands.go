package cli

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	utils "github.com/anirudh97/GollabEdit/pkg"
	"github.com/spf13/cobra"
)

type userCredentials struct {
	email    string
	password string
}

type newAccount struct {
	username string
	password string
	email    string
}

// Creates the create account command
func createAccountCommand() *cobra.Command {
	var newAcc newAccount

	createAccountCmd := &cobra.Command{
		Use:   "create",
		Short: "Create an account by passing username, password and email",
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

			if !utils.ValidateEmail(newAcc.email) {
				fmt.Printf("%s is not an email format", newAcc.email)
				return
			}
			requestData := map[string]string{
				"username": newAcc.username,
				"password": newAcc.password,
				"email":    newAcc.email,
			}

			req := utils.DefaultRequestConfig()
			requestJSON, err := json.Marshal(requestData)
			if err != nil {
				fmt.Println("Error creating request data:", err)
				return
			}
			req.API = "/auth/create"
			req.Method = utils.POST
			req.Payload = requestJSON

			// Make the API request
			resp, statusCode, err := utils.MakeRequest(req)
			if err != nil {
				fmt.Println("Error making request:", err)
				return
			}
			var response utils.Response
			err = json.Unmarshal(resp, &response)
			if err != nil {
				fmt.Println("Error parsing response:", err)
				return
			}

			if statusCode != http.StatusCreated {
				fmt.Println("-------------------------------------------------------")
				fmt.Println("Failed to create account!")
				fmt.Println("Status Code: ", statusCode)
				fmt.Println("Error Message: ", response.Error)
				fmt.Println("-------------------------------------------------------")
				return
			}

			var responseData struct {
				Username string `json:"username"`
				Email    string `json:"email"`
			}
			err = json.Unmarshal(response.Data, &responseData)
			if err != nil {
				fmt.Println("Error parsing response data:", err)
				return
			}

			fmt.Println("-------------------------------------------------------")
			fmt.Printf("Account Created Successfully!\n")
			fmt.Printf("Username: %s\n", responseData.Username)
			fmt.Printf("Email: %s\n", responseData.Email)
			fmt.Println("-------------------------------------------------------")
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
			if user.email == "" {
				fmt.Println("Please provide an username using --email flag.")
				return
			}
			if user.password == "" {
				fmt.Println("Please provide password using --pwd flag.")
				return
			}

			requestData := map[string]string{
				"password": user.password,
				"email":    user.email,
			}

			req := utils.DefaultRequestConfig()
			requestJSON, err := json.Marshal(requestData)
			if err != nil {
				fmt.Println("Error creating request data:", err)
				return
			}
			req.API = "/auth/login"
			req.Method = utils.POST
			req.Payload = requestJSON

			// Make the API request
			resp, statusCode, err := utils.MakeRequest(req)
			if err != nil {
				fmt.Println("Error making request:", err)
				return
			}
			var response utils.Response
			err = json.Unmarshal(resp, &response)
			if err != nil {
				fmt.Println("Error parsing response:", err)
				return
			}

			if statusCode != http.StatusOK {
				fmt.Println("-------------------------------------------------------")
				fmt.Println("Login failed!")
				fmt.Println("Status Code: ", statusCode)
				fmt.Println("Error Message: ", response.Error)
				fmt.Println("-------------------------------------------------------")
				return
			}

			var responseData struct {
				Username string `json:"username"`
				Email    string `json:"email"`
				Token    string `json:"token"`
			}
			err = json.Unmarshal(response.Data, &responseData)
			if err != nil {
				fmt.Println("Error parsing response data:", err)
				return
			}

			fmt.Println("-------------------------------------------------------")
			fmt.Printf("Logged In Successfully!\n")
			fmt.Printf("Username: %s\n", responseData.Username)
			fmt.Printf("Email: %s\n", responseData.Email)
			fmt.Println("Token: ", responseData.Token)
			fmt.Println("-------------------------------------------------------")

			os.Setenv("GOLLABEDIT_USERNAME", responseData.Username)
			os.Setenv("GOLLABEDIT_TOKEN", responseData.Token)
			os.Setenv("GOLLABEDIT_EMAIL", responseData.Email)
		},
	}

	loginCmd.Flags().StringVar(&user.email, "email", "", "Email")
	loginCmd.Flags().StringVar(&user.password, "pwd", "", "Password")

	return loginCmd, user

}

// Creates the logout command
func createLogoutCommand() *cobra.Command {
	logoutCmd := &cobra.Command{
		Use:   "logout",
		Short: "Logout",
		Run: func(cmd *cobra.Command, args []string) {

			req := utils.DefaultRequestConfig()
			req.API = "/auth/logout"
			req.Method = utils.POST
			req.Payload = nil

			// Make the API request
			_, _, err := utils.MakeRequest(req)
			if err != nil {
				fmt.Println("Error making request:", err)
				return
			}
			fmt.Println("You have been successfully logged out.")
		},
	}

	os.Unsetenv("GOLLABEDIT_USERNAME")
	os.Unsetenv("GOLLABEDIT_TOKEN")
	os.Unsetenv("GOLLABEDIT_EMAIL")

	return logoutCmd
}
