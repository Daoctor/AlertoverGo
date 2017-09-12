package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
)

func main() {
	args := os.Args[1:]
	// args = append(args, "test content")
	// args = append(args, "test title")

	var content string
	var title string
	sourceID := "xxx"
	receiverID := "xxx"

	switch argLen := len(args); argLen {
	case 1:
		content = args[0]
		title = "From Bot"
	case 2:
		content = args[0]
		title = args[1]
	default:
		title = "From Bot"

		fi, err := os.Stdin.Stat()

		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		if (fi.Mode()&os.ModeNamedPipe == 0) && (fi.Size() <= 0) {
			fmt.Print("Args Not right!\n")
			fmt.Print("Usage:\n")
			fmt.Print("    ./sendmsg CONTENT [TITLE]\n")
			fmt.Print("    echo hello |./sendmsg \n")
			fmt.Print("    ./sendmsg < file.txt")
			return
		}
		byteContent, _ := ioutil.ReadAll(os.Stdin)
		content = string(byteContent)

	}
	if len(content) > 1500 {
		fmt.Print("Content to large\n")
		os.Exit(1)
	}
	postURL := "https://api.alertover.com/v1/alert"
	// postURL := "https://httpbin.org/post"

	PostData := url.Values{"source": {sourceID},
		"receiver": {receiverID}, "content": {"c"}, "title": {"t"}}
	PostData.Set("title", title)
	PostData.Set("content", content)

	resp, err := http.PostForm(postURL, PostData)
	if err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))

	p, _ := regexp.Compile(`\"code\"\:\s{0,5}([\d\-]+)\,`)
	bodyData := string(body)
	r := p.FindStringSubmatch(bodyData)
	if len(r) < 2 {
		fmt.Print("Not found Response code, response detail:\n")
		fmt.Print(bodyData)
		os.Exit(1)
	} else {
		if r[1] != "0" {
			fmt.Print("Response Error, response_detail\n")
			fmt.Print(bodyData)
			os.Exit(1)
		} else {
			return
		}
	}

}
