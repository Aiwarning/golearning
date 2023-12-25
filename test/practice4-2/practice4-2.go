package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

const (
	owner      = "Aiwarning"
	repository = "golearning"
)

var (
	client *github.Client
	token  = "your-Issues-access-token"
)

func init() {
	ctx := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	))
	client = github.NewClient(ctx)
}

func main() {
	fmt.Println("GitHub Issue Tool")

	for {
		fmt.Println("\nOptions:")
		fmt.Println("1. Create Issue")
		fmt.Println("2. Read Issues")
		fmt.Println("3. Update Issue")
		fmt.Println("4. Close Issue")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice (1-5): ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			createIssue()
		case 2:
			readIssues()
		case 3:
			updateIssue()
		case 4:
			closeIssue()
		case 5:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 5.")
		}
	}
}

func createIssue() {
	fmt.Println("\nCreating Issue:")
	fmt.Print("Title: ")
	title := getUserInput()

	fmt.Print("Body: ")
	body := getUserInput()

	issue := &github.IssueRequest{
		Title: &title,
		Body:  &body,
	}

	newIssue, _, err := client.Issues.CreateIssue(owner, repository, issue)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Issue created. Number: %d\n", newIssue.GetNumber())
}

func readIssues() {
	fmt.Println("\nReading Issues:")
	issues, _, err := client.Issues.ListByRepo(owner, repository, nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, issue := range issues {
		fmt.Printf("Issue #%d: %s\n", issue.GetNumber(), issue.GetTitle())
	}
}

func updateIssue() {
	fmt.Println("\nUpdating Issue:")
	fmt.Print("Enter issue number to update: ")
	var issueNumber int
	fmt.Scan(&issueNumber)

	issue, _, err := client.Issues.Get(owner, repository, issueNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Current Title: %s\n", issue.GetTitle())
	fmt.Print("Enter new title: ")
	title := getUserInput()

	fmt.Printf("Current Body: %s\n", issue.GetBody())
	fmt.Print("Enter new body: ")
	body := getUserInput()

	issueUpdate := &github.IssueRequest{
		Title: &title,
		Body:  &body,
	}

	updatedIssue, _, err := client.Issues.Edit(owner, repository, issueNumber, issueUpdate)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Issue #%d updated.\n", updatedIssue.GetNumber())
}

func closeIssue() {
	fmt.Println("\nClosing Issue:")
	fmt.Print("Enter issue number to close: ")
	var issueNumber int
	fmt.Scan(&issueNumber)

	_, err := client.Issues.Edit(owner, repository, issueNumber, &github.IssueRequest{State: github.String("closed")})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Issue #%d closed.\n", issueNumber)
}

func getUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func openEditor(filename, initialContent string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return fmt.Errorf("EDITOR environment variable not set")
	}

	cmd := exec.Command(editor, filename)
	cmd.Stdin = strings.NewReader(initialContent)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
