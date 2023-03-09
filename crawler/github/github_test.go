package github

import (
	"flag"
	"fmt"
	"os"
)

func Example() {
	gh := New("chengui/zypp-garden")

	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	searchFilter := searchCmd.String("filter", "", "search filter")
	searchState := searchCmd.String("state", "", "search state")
	searchSort := searchCmd.String("sort", "", "search sort")

	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	createTitle := createCmd.String("title", "", "issue title")
	createBody := createCmd.String("body", "", "issue body")

	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getId := getCmd.Int("id", 0, "issue id")

	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateId := updateCmd.Int("id", 0, "issue id")
	updateTitle := updateCmd.String("title", "", "issue title")
	updateBody := updateCmd.String("body", "", "issue body")

	closeCmd := flag.NewFlagSet("close", flag.ExitOnError)
	closeId := closeCmd.Int("id", 0, "issue id")

	switch os.Args[1] {
	case "search":
		searchCmd.Parse(os.Args[2:])
		issues, err := gh.SearchIssues(map[string]string{
			"filter": *searchFilter,
			"state":  *searchState,
			"sort":   *searchSort,
		})
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(2)
		}
		fmt.Println(issues)
	case "create":
		createCmd.Parse(os.Args[2:])
		_, err := gh.CreateIssue(&NewIssue{
			Title: *createTitle,
			Body:  *createBody,
		})
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(2)
		}
		fmt.Println("create successfully")
	case "get":
		getCmd.Parse(os.Args[2:])
		issue, err := gh.GetIssue(*getId)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(2)
		}
		fmt.Println(issue)
	case "update":
		updateCmd.Parse(os.Args[2:])
		_, err := gh.UpdateIssue(*updateId, &NewIssue{
			Title: *updateTitle,
			Body:  *updateBody,
		})
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(2)
		}
		fmt.Printf("update %v successfully.\n", updateId)
	case "close":
		closeCmd.Parse(os.Args[2:])
		_, err := gh.CloseIssue(*closeId)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(2)
		}
		fmt.Printf("close %v successfully.\n", closeId)
	default:
		flag.Usage()
	}
}
