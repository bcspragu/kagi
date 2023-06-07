# Kagi FastGPT CLI

[![GoDoc](https://pkg.go.dev/badge/github.com/bcspragu/kagi?status.svg)](https://pkg.go.dev/github.com/bcspragu/kagi?tab=doc)
[![CI Workflow](https://github.com/bcspragu/kagi/actions/workflows/test.yml/badge.svg)](https://github.com/bcspragu/kagi/actions?query=branch%3Amain)



`kagi` is a quick-and-dirty CLI for querying the [Kagi search engine](https://kagi.com/) with their [FastGPT API](https://help.kagi.com/kagi/api/fastgpt.html)

## Installation

First, build + install the `kagi` CLI with:

```
go install github.com/bcspragu/kagi/cmd/kagi@latest
```

Then, create an API key + add API credits, following [the official Kagi instructions](https://help.kagi.com/kagi/api/fastgpt.html#quick-start). Add the API key to your `~/.bashrc` or equivalent with something like:

```
export KAGI_API_KEY=...
```

And you should be good to go! Try running `kagi <query>` to test it out.

As an aside, I'm **really** not a fan of storing sensitive credentials in accessible-to-everything-all-the-time environment variables, and if anyone has a good + ergonomic alternative (e.g. involving `pass` or credential helpers), I'm all ears.

## Example Output

```
$ kagi net/http golang
===== OUTPUT =====

The net/http package is part of the Go standard library and provides HTTP client and server functionality. [1][2][3]

===== REFERENCES =====

1. http package - net/http - Go Packages - https://pkg.go.dev/net/http
  - Package http provides HTTP client and server implementations. ... Manually configuring HTTP/2 via the golang.org/x/net/http2 package takes precedence over...

2. Writing Web Applications - The Go Programming Language - https://go.dev/doc/articles/wiki/
  - Covered in this tutorial: Creating a data structure with load and save methods; Using the net/http package to build web applications; Using the html/...

3. How To Make an HTTP Server in Go | DigitalOcean - https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go
  - Apr 21, 2022 ... The net/http package not only includes the ability to make HTTP requests, but also provides an HTTP server you can use to handle those requests.
```

## Troubleshooting

If the `kagi` tool isn't working for you, make sure:

- You've reloaded your `~/.bashrc` or equivalent after adding your API key, e.g. with `source ~/.bashrc` or opening a new shell.
  - Run `echo $KAGI_API_KEY` to confirm it's set in the current shell
- You've 'topped up' your API credits
  - Confirm there's a non-zero 'Credit remaining' balance in [the Kagi billing UI](https://kagi.com/settings?p=billing_api)

If you're still having issues after that, feel free to file an issue, including any error output.
