package main

import (
	"errors"
	"flag"
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

var errUsage = errors.New("usage: kagi <query>")

func run(args []string) error {
	if len(args) == 0 {
		return errUsage
	}
	fs := flag.NewFlagSet(args[0], flag.ContinueOnError)
	var (
		kagiAPIKey = fs.String("kagi_api_key", os.Getenv("KAGI_API_KEY"), "API key to use with the Kagi FastGPT API")
	)
	if err := fs.Parse(args[1:]); err != nil {
		return fmt.Errorf("failed to parse flags: %w", err)
	}
	if *kagiAPIKey == "" {
		return errors.New("no KAGI_API_KEY env var or --kagi_api_key flag was set")
	}

	fArgs := fs.Args()
	if len(fArgs) == 0 {
		return errUsage
	}
	query := strings.Join(fArgs, " ")

	client := api.NewClient(*kagiAPIKey)

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
