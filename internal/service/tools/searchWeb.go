package tools

import (
	"fmt"
	"net/url"
	"os/exec"
)

func SearchWeb(query string) (string, error) {
	q := url.QueryEscape(query)

	searchUrl := fmt.Sprintf("https://www.google.com/search?q=%s", q)

	cmd := exec.Command("open", searchUrl)

	return "Successfuly opened", cmd.Run()
}

func SearchYoutube(query string) (string, error) {
	q := url.QueryEscape(query)

	searchUrl := fmt.Sprintf("https://www.youtube.com/results?search_query=%s", q)

	cmd := exec.Command("open", searchUrl)

	return "Successfuly opened", cmd.Run()
}
