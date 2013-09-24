package main

import (
	"code.google.com/p/go.net/html"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Config struct {
	Tag           string
	AttributeName string
	IsURL         bool
}

func DownloadURL(url string) string {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(contents)
}

func main() {
	var cfg Config
	lines := []string{}
	if _, err := toml.DecodeFile(os.Args[1], &cfg); err != nil {
		log.Fatal(err)
		return
	}
	htmlData := DownloadURL(os.Args[2])
	doc, err := html.Parse(strings.NewReader(htmlData))
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == cfg.Tag {
			for _, tag := range n.Attr {
				if tag.Key == cfg.AttributeName {
					lines = append(lines, tag.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	if cfg.IsURL {
		u, err := url.Parse(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		for i := range lines {
			if strings.HasPrefix(lines[i], "#") {
				lines[i] = os.Args[2] + lines[i]
			} else if strings.HasPrefix(lines[i], "../") {
				lines[i] = os.Args[2] + "/" + lines[i]
			} else if strings.HasPrefix(lines[i], "//") {
				lines[i] = u.Scheme + ":" + lines[i]
			} else if strings.Index(lines[i], ":") == -1 { //this might not work all the time.
				lines[i] = u.Scheme + "://" + u.Host + "/" + lines[i]
			}
		}
	}
	for i := range lines {
		fmt.Println(lines[i])
	}

}
