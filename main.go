package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"io"
	"net/http"
	"regexp"
	"sort"
)

type WebsiteResponse struct {
	URL          string
	ResponseSize int64
}

func main() {
	var d []WebsiteResponse
	c := make(chan WebsiteResponse)

	// https://uibakery.io/regex-library/url

	var args struct {
		Urls []string `arg:"required"`
	}
	arg.MustParse(&args)

	for _, url := range args.Urls {
		go handleRequestsConcurrently(url, c)
	}

	for range args.Urls {
		d = append(d, <-c)
	}

	sortResponses(d)

	for _, r := range d {
		fmt.Printf("%s\t%v bytes\n", r.URL, r.ResponseSize)
	}
}

func handleRequestsConcurrently(url string, c chan<- WebsiteResponse) {
	rHttpUrl := regexp.MustCompile("^https?:\\/\\/(?:www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b(?:[-a-zA-Z0-9()@:%_\\+.~#?&\\/=]*)$")
	rUrl := regexp.MustCompile("(?:www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b(?:[-a-zA-Z0-9()@:%_\\+.~#?&\\/=]*)$")
	if rUrl.MatchString(url) {
		if !rHttpUrl.MatchString(url) {
			url = "https://" + url
		}
		responseSize := handleRequest(url)

		c <- WebsiteResponse{
			URL:          url,
			ResponseSize: responseSize,
		}
	} else {
		c <- WebsiteResponse{
			URL:          url,
			ResponseSize: 0,
		}
	}
}

// https://www.notebookcast.com/en/blog/golang-sort-array-struct
func sortResponses(rl []WebsiteResponse) {
	sort.Slice(rl, func(i, j int) bool { return rl[i].ResponseSize > rl[j].ResponseSize })
}

func handleRequest(url string) int64 {
	resp, err := http.Get(url)
	if err != nil {
		return 0
	}
	b := getResponseBodySize(resp)
	return b
}

func getResponseBodySize(resp *http.Response) int64 {
	b, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		return 0
	}
	return b
}
