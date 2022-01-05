package main

import (
	"encoding/json"
	"fmt"
	"github.com/jessevdk/go-flags"
	"net/url"
	"os"
	"text/template"
)

var options struct {
	Format string `short:"f" long:"format" name:"output format" description:"format output with text/template syntax"`
	Args   struct {
		Input string `positional-arg-name:"input"`
	} `positional-args:"true" required:"1"`
}

type urlObject struct {
	*url.URL
	Query url.Values `json:"Query"`
}

func logAndExit(err error) {
	if _, pErr := fmt.Fprintln(os.Stderr, err); pErr != nil {
		panic(pErr)
	}

	os.Exit(1)
}

func main() {
	parser := flags.NewParser(&options, flags.Default)

	if _, err := parser.Parse(); err != nil {
		os.Exit(1)
	}

	parsedURL, err := url.Parse(options.Args.Input)
	if err != nil {
		logAndExit(err)
	}

	output, err := json.Marshal(urlObject{
		URL:   parsedURL,
		Query: parsedURL.Query(),
	})
	if err != nil {
		logAndExit(err)
	}

	if options.Format != "" {
		tpl, err := template.New("gurl").Parse(options.Format + "\n")
		if err != nil {
			logAndExit(err)
		}

		if err := tpl.Execute(os.Stdout, parsedURL); err != nil {
			logAndExit(err)
		}

		return
	}

	if _, err := fmt.Fprintln(os.Stdout, string(output)); err != nil {
		logAndExit(err)
	}
}
