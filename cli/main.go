package main

import (
	"flag"
	"fmt"
	"os"
)

func HandleGet(getCmd *flag.FlagSet, all *bool, id *string) {
	getCmd.Parse(os.Args[2:])
	if !*all && *id == "" {
		fmt.Println("id is required or specify -- all for all videos")
		getCmd.PrintDefaults()
		os.Exit(1)
	}
	if *all {
		fmt.Println("Get all videos")
		return
	}
	if *id != "" {
		fmt.Println("Get video with ID", id)
	}
}

func HandleAdd(getCmd *flag.FlagSet, title *string) {
	if *title == "" {
		fmt.Println("titile is required to create video")
		getCmd.PrintDefaults()
	} else {
		fmt.Println("Video ", title, "will be created")
	}
}

func main() {

	// get command
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	getAll := getCmd.Bool("all", false, "Get all videos")
	getId := getCmd.String("id", "", "Get video by id")

	// add command
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addTitle := getCmd.String("title", "", "Get video by id")

	if len(os.Args) < 2 {
		fmt.Println("expected 'get' or 'add' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		HandleGet(getCmd, getAll, getId)
	case "add":
		HandleAdd(addCmd, addTitle)
	default:
		fmt.Println("Option not found")
	}

}
