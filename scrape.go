package main

import (
	"code.google.com/p/go.net/html"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Config struct {
	Tag           string
	AttributeName string
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
	os.Create(os.Args[3])
	write, _ := os.OpenFile(os.Args[3], os.O_RDWR|os.O_APPEND, 0666)
	//write.WriteString(os.Args[2] + "\n") This line will need to be uncommented when I make a script to deal with things that are linked like ../file.ext, with out the whole URL. Unless it works already.
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
					write.WriteString(tag.Val + "\n")
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}
