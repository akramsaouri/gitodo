package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/shurcooL/graphql"
	"golang.org/x/oauth2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	client := graphql.NewClient("https://api.github.com/graphql", httpClient)
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
	err = client.Query(context.Background(), &query, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, edge := range query.Viewer.Repositories.Edges {
		node := edge.Node
		text := string(node.Object.Blob.Text)
		lastIndex := strings.LastIndex(text, "## TODO")
		if lastIndex == -1 {
			continue
		}
		fmt.Printf("# %s (%s)\n\n", node.Name, node.URL)
		lines := strings.Split(text[lastIndex:], "\n")
		for index, line := range lines {
			if index != 0 && line != "" {
				fmt.Println(line)
			}
		}
		fmt.Print(strings.Repeat("-", 80) + "\n\n")
	}
}
