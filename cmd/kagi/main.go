package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bcspragu/kagi/api"
)

func main() {
	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	var query string
	switch len(args) {
	case 0, 1:
		return errors.New("usage: kagi <query>")
	case 2:
		query = args[1]
	default:
		query = strings.Join(args[1:], " ")
	}
	if len(args) < 2 {
	}
	apiKey := os.Getenv("KAGI_API_KEY")
	if apiKey == "" {
		return errors.New("no KAGI_API_KEY was set")
	}

	client := api.NewClient(apiKey)

	resp, err := client.QueryFastGPT(query)
	if err != nil {
		return fmt.Errorf("error performing query: %w", err)
	}

	fmt.Println("===== OUTPUT =====")
	fmt.Println()
	fmt.Println(resp.Data.Output)
	fmt.Println()
	fmt.Println("===== REFERENCES =====")
	fmt.Println()

	for i, ref := range resp.Data.References {
		fmt.Printf("%d. %s - %s\n  - %s\n\n", i+1, ref.Title, ref.Link, ref.Snippet)
	}

	return nil
}
