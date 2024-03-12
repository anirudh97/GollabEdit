package cli

import (
	"encoding/json"
	"fmt"
	"net/http"

	utils "github.com/anirudh97/GollabEdit/pkg"
	"github.com/spf13/cobra"
)

type file struct {
	filename string
	location string
	owner    string
}

type shareFile struct {
	filename        string
	location        string
	sharedByEmail   string
	sharedWithEmail string
	permission      string
}

func createFileCommand() *cobra.Command {
	var f file

	createFileCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a file by passing filename and location",
		Run: func(cmd *cobra.Command, args []string) {
			config, err := utils.ReadConfigJson()
			if err != nil {
				fmt.Println("You are not Logged In. Please login.")
				return
			}

			f.owner = config.Email
			// Filename validation
			if !utils.ValidateFilename(f.filename) {
				fmt.Println("Invalid file format. Only txt or text formats allowed.")
				return
			}

			requestData := map[string]string{
				"filename": f.filename,
				"location": f.location,
				"owner":    f.owner,
			}

			req := utils.DefaultRequestConfig()
			requestJSON, err := json.Marshal(requestData)
			if err != nil {
				fmt.Println("Error creating request data:", err)
				return
			}
			req.API = "/file"
			req.Method = utils.POST
			req.Payload = requestJSON
			req.Bearer = config.Token

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
				fmt.Println("Failed to create file!")
				fmt.Println("Status Code: ", statusCode)
				fmt.Println("Error Message: ", response.Error)
				fmt.Println("-------------------------------------------------------")
				return
			}

			var responseData struct {
				Filename  string `json:"filename"`
				Location  string `json:"location"`
				CreatedAt string `json:"createdAt"`
			}
			err = json.Unmarshal(response.Data, &responseData)
			if err != nil {
				fmt.Println("Error parsing response data:", err)
				return
			}

			fmt.Println("-------------------------------------------------------")
			fmt.Printf("File Created Successfully!\n")
			fmt.Printf("Filename: %s\n", responseData.Filename)
			fmt.Printf("Location: %s\n", responseData.Location)
			fmt.Println("Timestamp in UTC: ", responseData.CreatedAt)
			fmt.Println("-------------------------------------------------------")
		},
	}

	createFileCmd.Flags().StringVar(&f.filename, "filename", "", "Filename")
	createFileCmd.Flags().StringVar(&f.location, "location", "", "Location")

	return createFileCmd
}

func shareFileCommand() *cobra.Command {
	var f shareFile

	shareFileCmd := &cobra.Command{
		Use:   "share",
		Short: "Share a file by passing filename, location, email of user whom you want to share, permission",
		Run: func(cmd *cobra.Command, args []string) {
			config, err := utils.ReadConfigJson()
			if err != nil {
				fmt.Println("You are not Logged In. Please login.")
				return
			}

			f.sharedByEmail = config.Email

			// Filename validation
			if !utils.ValidateFilename(f.filename) {
				fmt.Println("Invalid file format. Only txt or text formats allowed.")
				return
			}

			requestData := map[string]string{
				"filename":        f.filename,
				"location":        f.location,
				"permission":      f.permission,
				"sharedWithEmail": f.sharedWithEmail,
				"sharedByEmail":   f.sharedByEmail,
			}

			req := utils.DefaultRequestConfig()
			requestJSON, err := json.Marshal(requestData)
			if err != nil {
				fmt.Println("Error creating request data:", err)
				return
			}
			req.API = "/file/share"
			req.Method = utils.POST
			req.Payload = requestJSON
			req.Bearer = config.Token

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
				fmt.Println("Failed to share file!")
				fmt.Println("Status Code: ", statusCode)
				fmt.Println("Error Message: ", response.Error)
				fmt.Println("-------------------------------------------------------")
				return
			}

			var responseData struct {
				Filename        string `json:"filename"`
				Location        string `json:"location"`
				Permission      string `json:"permission"`
				SharedWithEmail string `json:"sharedWithEmail"`
			}
			err = json.Unmarshal(response.Data, &responseData)
			if err != nil {
				fmt.Println("Error parsing response data:", err)
				return
			}

			fmt.Println("-------------------------------------------------------")
			fmt.Printf("File Shared Successfully!\n")
			fmt.Printf("Filename: %s\n", responseData.Filename)
			fmt.Printf("Location: %s\n", responseData.Location)
			fmt.Println("Permission: ", responseData.Permission)
			fmt.Println("Shared With: ", responseData.SharedWithEmail)
			fmt.Println("-------------------------------------------------------")
		},
	}

	shareFileCmd.Flags().StringVar(&f.filename, "filename", "", "Filename")
	shareFileCmd.Flags().StringVar(&f.location, "location", "", "Location")
	shareFileCmd.Flags().StringVar(&f.sharedWithEmail, "shareWithEmail", "", "ShareWithEmail")
	shareFileCmd.Flags().StringVar(&f.permission, "permission", "", "Permission")

	return shareFileCmd
}

func openFileCommand() *cobra.Command {
	var f file

	openFileCmd := &cobra.Command{
		Use:   "open",
		Short: "Open a file by passing filename and location",
		Run: func(cmd *cobra.Command, args []string) {
			config, err := utils.ReadConfigJson()
			if err != nil {
				fmt.Println("You are not Logged In. Please login.")
				return
			}

			f.owner = config.Email
			// Filename validation
			if !utils.ValidateFilename(f.filename) {
				fmt.Println("Invalid file format. Only txt or text formats allowed.")
				return
			}

			requestData := map[string]string{
				"filename": f.filename,
				"location": f.location,
				"owner":    f.owner,
			}

			req := utils.DefaultRequestConfig()
			requestJSON, err := json.Marshal(requestData)
			if err != nil {
				fmt.Println("Error creating request data:", err)
				return
			}
			req.API = "/file/open"
			req.Method = utils.POST
			req.Payload = requestJSON
			req.Bearer = config.Token

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
				fmt.Println("Failed to open file!")
				fmt.Println("Status Code: ", statusCode)
				fmt.Println("Error Message: ", response.Error)
				fmt.Println("-------------------------------------------------------")
				return
			}

			// var responseData struct {
			// 	Filename  string `json:"filename"`
			// 	Location  string `json:"location"`
			// 	CreatedAt string `json:"createdAt"`
			// }
			// err = json.Unmarshal(response.Data, &responseData)
			// if err != nil {
			// 	fmt.Println("Error parsing response data:", err)
			// 	return
			// }

			fmt.Println("-------------------------------------------------------")
			fmt.Printf("File Opened Successfully!\n")
			fmt.Printf("Filename: %s\n", f.filename)
			fmt.Printf("Location: %s\n", f.location)
			fmt.Println("-------------------------------------------------------")
		},
	}

	openFileCmd.Flags().StringVar(&f.filename, "filename", "", "Filename")
	openFileCmd.Flags().StringVar(&f.location, "location", "", "Location")

	return openFileCmd
}
