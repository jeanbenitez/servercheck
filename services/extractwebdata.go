package services

import (
	"fmt"
	"net/http"
	"regexp"

	"golang.org/x/net/html"
)

// ExtractWebData analyze given a url and a basurl, recoursively scans the page
// following all the links and fills the `visited` map
func ExtractWebData(domain string) (title string, logo string) {
	url := "http://" + domain
	page, err := parse(url)
	if err != nil {
		fmt.Printf("Error getting page %s %s\n", url, err)
		return
	}
	title = pageTitle(page)
	logo = pageLogo(page)
	if logo != "" {
		globalURL, _ := regexp.MatchString("^http", logo)
		if !globalURL {
			if logo[0] != 47 {
				logo = "/" + logo
			}
			logo = url + logo
		}
	}
	return
}

// parse given a string pointing to a URL will fetch and parse it
// returning an html.Node pointer
func parse(url string) (*html.Node, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Cannot get page")
	}
	b, err := html.Parse(r.Body)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse page")
	}
	return b, err
}

// pageTitle finds the title tag
func pageTitle(n *html.Node) string {
	var title string
	if n.Type == html.ElementNode && n.Data == "title" {
		return n.FirstChild.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		title = pageTitle(c)
		if title != "" {
			break
		}
	}
	return title
}

// pageLogo finds og:image in meta tags
func pageLogo(n *html.Node) string {
	var logo string
	if n.Type == html.ElementNode && (n.Data == "meta" || n.Data == "link") {
		key := "content"
		if n.Data == "link" {
			key = "href"
		}
		ok := false
		for _, attr := range n.Attr {
			if n.Data == "meta" && (attr.Key == "property" && attr.Val == "og:image") || (attr.Key == "itemprop" && attr.Val == "image") {
				ok = true
			} else if n.Data == "link" && attr.Key == "rel" {
				icon, _ := regexp.MatchString("icon", attr.Val)
				if icon {
					ok = true
				}
			}

			if attr.Key == key {
				logo = attr.Val
			}
		}

		if ok && logo != "" {
			return logo
		}

		logo = ""
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		logo = pageLogo(c)
		if logo != "" {
			break
		}
	}
	return logo
}
