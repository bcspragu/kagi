# Kagi FastGPT CLI

[![GoDoc](https://pkg.go.dev/badge/github.com/bcspragu/kagi?status.svg)](https://pkg.go.dev/github.com/bcspragu/kagi?tab=doc)
[![CI Workflow](https://github.com/bcspragu/kagi/actions/workflows/test.yml/badge.svg)](https://github.com/bcspragu/kagi/actions?query=branch%3Amain)



`kagi` is a simple CLI for querying the [Kagi search engine](https://kagi.com/) with their [FastGPT API](https://help.kagi.com/kagi/api/fastgpt.html)

## Installation

First, build + install the `kagi` CLI with:

```
go install github.com/bcspragu/kagi/cmd/kagi@latest
```

Then, create an API key + add API credits, following [the official Kagi instructions](https://help.kagi.com/kagi/api/fastgpt.html#quick-start). Add the API key to your `~/.bashrc` or equivalent with something like:

```
export KAGI_API_KEY=...
```

If you don't want to expose sensitive credentials to all applications running in your shell, you can wrap the `kagi` CLI in a shell script, e.g.:

```bash
#!/bin/bash

# Maybe this lives in ~/.local/bin, which has a higher precedence than your $GOPATH/bin dir

KAGI_API_KEY=$(some CLI password manager or just hardcode it) $GOPATH/bin/kagi "$@"

# or, if you prefer flags:
# $GOPATH/bin/kagi --kagi_api_key=... "$@"
```

And you should be good to go! Try running `kagi <query>` to test it out.

## Usage with Glow

Since the output of the Kagi CLI is Markdown ([H/T @lanej](https://github.com/bcspragu/kagi/pull/2)), you can pipe it to a CLI markdown rendering tool like [Glow](https://github.com/charmbracelet/glow), e.g:

```bash
# Assuming the KAGI_API_KEY is specified somewhere
WIDTH="$(stty size |  awk '{ print $2 }' | xargs -I {} echo '{}-10' | bc)"
$GOPATH/bin/kagi "$@" | glow -p -w "$WIDTH"
```

## Example Output

```
$ kagi net/http golang
# net/http golang
Here is a summary of the key information about the ```net/http``` package in Go based on the provided sources:
- The ```net/http``` package in Go provides functions and types for building HTTP servers and clients. **【1】【2】**
- The ```http.ResponseWriter``` is used to assemble the HTTP server's response, and writing to it sends data to the HTTP client. **【2】**
- The ```http.Request``` represents the client HTTP request, and ```r.URL.Path``` contains the path component of the request URL. **【2】**
- The ```http.Handler``` interface is a key component of the ```net/http``` package, as it defines the contract for handling HTTP requests. **【3】**
- The ```http.DetectContentType``` function is used to detect the MIME type of a given byte slice, but it has somelimitations in detecting certain types of content. **【4】**
- Common errors that can occur when using the ```net/http``` package include ```ErrBodyNotAllowed```, which is returned when the HTTP method or response code does not permit a body, and ```ErrHijacked```, which is returned when the underlying connection has been hijacked. **【1】**
- The ```http.RemoteAddr``` field in the ```http.Request``` struct allows HTTP servers to record the network address that sent the request, usually for logging purposes. **【5】**
In summary, the ```net/http``` package provides a powerful and flexible way to build HTTP servers and clients in Go, with a well-designed set of types and functions for handling common web application tasks.

# References
1. net/http - Go Packages - https://pkg.go.dev/net/http - View Source var ( // ErrBodyNotAllowed is returned by ResponseWriter.Write calls // when the HTTP method or response code does not permit a // body. ErrBodyNotAllowed = errors. New("http: request method or response status code does not allow body") // ErrHijacked is returned by ResponseWriter.Write calls when // the underlying connection has been hijacked using the // Hijacker interface.
2. Writing Web Applications - The Go Programming Language - https://go.dev/doc/articles/wiki/ - An http.ResponseWriter value assembles the HTTP server's response; by writing to it, we send data to the HTTP client. An http.Requestis a data structure that represents the client HTTP request. r.URL.Path is the path component of the request URL.
3. The net/http package | Building Web Apps with Go - https://codegangsta.gitbooks.io/building-web-apps-with-go/content/http_basics/index.html - The Browser makes an HTTP request with some information, the Server then processes that request and returns a Response. This pattern of request-response is one of the key focal points in building webapplications in Go. In fact, the net/http package's most important piece is the http.Handler Interface. The http.Handler Interface
4. go - Golang “net/http” DetectContentType error - https://stackoverflow.com/questions/40601725/golang-net-http-detectcontenttype-error - The standard library's code is only supposed to detect certain types (like HTML that contains one of a few common tags) according to a certain standardized algorithm--more in https://golang.org/src/net/http/sniff.go . ... How can I check specific golang net/http error code? ... simple client on golang returns " net/http: HTTP/1.x transport connection broken: unexpected EOF "
5. Golang-web-dev/017_understanding-net-http-package/README.md... - https://github.com/GoesToEleven/golang-web-dev/blob/master/017_understanding-net-http-package/README.md - PostForm url.Values MultipartForm *multipart.Form // RemoteAddr allows HTTP servers and other software to record // the network address that sent the request, usually for// logging.The HTTP client ignores PostForm and uses Body instead.

```

## Troubleshooting

If the `kagi` tool isn't working for you, make sure:

- You've reloaded your `~/.bashrc` or equivalent after adding your API key, e.g. with `source ~/.bashrc` or opening a new shell.
  - Run `echo $KAGI_API_KEY` to confirm it's set in the current shell
- You've 'topped up' your API credits
  - Confirm there's a non-zero 'Credit remaining' balance in [the Kagi billing UI](https://kagi.com/settings?p=billing_api)

If you're still having issues after that, feel free to file an issue, including any error output.
