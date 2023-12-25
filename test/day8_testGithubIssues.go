package main

import (
	"fmt"
	"log"
	"month_1/test/Issues"
	"os"
)

func main() {
	result, err := Issues.SearchIssues(os.Args[1:])
	fmt.Println()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for i, item := range result.Items {
		if i < 5 {
			fmt.Printf("#%-5d %9.9s %.55s %9.9s\n",
				item.Number, item.User.Login, item.Title, item.Body)
		}

	}
}
