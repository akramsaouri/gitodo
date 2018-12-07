package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/shurcooL/graphql"
	"golang.org/x/oauth2"
)

func main() {

	// authenticate with github oauth
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	client := graphql.NewClient("https://api.github.com/graphql", httpClient)

	// build and query readmes for all owned repos under master branch
	var query struct {
		Viewer struct {
			Repositories struct {
				Edges []struct {
					Node struct {
						Name   graphql.String `graphql:"name"`
						URL    graphql.String `graphql:"url"`
						Object struct {
							Blob struct {
								Text graphql.String `graphql:"text"`
							} `graphql:"... on Blob"`
						} `graphql:"object(expression: \"master:README.md\")"`
					} `graphql:"node"`
				} `graphql:"edges"`
			} `graphql:"repositories(last: 100, isFork: false, affiliations: [OWNER]) "`
		} `graphql:"viewer"`
	}
	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	// parse readme lines and build todos
	for _, edge := range query.Viewer.Repositories.Edges {
		repo := edge.Node
		readme := string(repo.Object.Blob.Text)
		todoIndex := strings.LastIndex(readme, "## TODO")
		if todoIndex == -1 {
			// no matches were found
			continue
		}
		lines := strings.Split(readme[todoIndex:], "\n")
		if len(lines) > 1 {
			// todo section is not empty
			var todos []string
			// gather todos
			for i := 1; i < len(lines); i++ {
				line := lines[i]
				if line != "" {
					todos = append(todos, line)
				}
			}
			if len(todos) > 0 {
				fmt.Printf("# %s (%s)\n\n", repo.Name, repo.URL)
				fmt.Println(strings.Join(todos, "\n"))
				fmt.Print(strings.Repeat("-", 80) + "\n\n")
			}
		}
	}
}
