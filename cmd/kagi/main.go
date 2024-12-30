package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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

	response, err := respond(resp, query)
	if err != nil {
		return fmt.Errorf("failed to build response: %w", err)
	}
	fmt.Print(response)

	return nil
}

func respond(resp *api.FastGPTResponse, query string) (string, error) {
	var buf bytes.Buffer
	buf.WriteString("# ")
	buf.WriteString(query)
	buf.WriteRune('\n')
	if err := streamAndRemoveDoubleNewlines(resp.Data.Output, &buf); err != nil {
		return "", fmt.Errorf("failed to remove double newlines: %w", err)
	}
	buf.WriteRune('\n')

	// If there are no references, return early
	if len(resp.Data.References) == 0 {
		return buf.String(), nil
	}

	buf.WriteString("\n# References\n")

	for i, ref := range resp.Data.References {
		// fmt.Sprintf("%d. %s - %s  - %s\n", i+1, ref.Title, ref.Link, ref.Snippet)
		buf.WriteString(strconv.Itoa(i + 1))
		buf.WriteString(". ")
		buf.WriteString(ref.Title)
		buf.WriteString(" - ")
		buf.WriteString(ref.Link)
		buf.WriteString(" - ")
		buf.WriteString(ref.Snippet)
		buf.WriteRune('\n')
	}

	return buf.String(), nil
}

// Remove all repeated newlines or empty lines from the given string
func streamAndRemoveDoubleNewlines(inp string, buf *bytes.Buffer) error {
	r := strings.NewReader(inp)

	sc := bufio.NewScanner(r)

	first := true
	for sc.Scan() {
		if sc.Text() == "" {
			continue
		}
		if !first {
			buf.WriteRune('\n')
		}
		first = false
		buf.WriteString(sc.Text())
	}

	if err := sc.Err(); err != nil {
		return fmt.Errorf("error reading input: %w", err)
	}

	return nil
}
