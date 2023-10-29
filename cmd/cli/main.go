package main

import (
	"fmt"
	"os"

	"github.com/anirudh97/GollabEdit/internal/cli"
)

func main() {

	cmd := cli.SetupCommands()

	err := cmd.Execute()
	if err != nil {
		fmt.Println("Some error occured! \n", err)
		os.Exit(1)
	}

}
