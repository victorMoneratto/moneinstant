package moneinstant

import (
	"fmt"

	"github.com/victormoneratto/moneinstant/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"strings"
)

// MyInstant a MyInstants instant
type MyInstant struct {
	name     string
	audioURL string
}

// Name returns the name of instant
func (instant MyInstant) Name() string {
	return instant.name
}

// AudioURL returns the instant audio url
func (instant MyInstant) AudioURL() string {
	return instant.audioURL
}

// MyInstants Source of instants
type MyInstants struct{}

// URL returns the url for instants page source
func (source *MyInstants) URL() string {
	return "http://www.myinstants.com"
}

// Home return the MyInstants homepage
func (source *MyInstants) Home() *InstantsPage {
	return NewInstantsPage(source, "")
}

// Trending return the MyInstants trending page
func (source *MyInstants) Trending() *InstantsPage {
	return NewInstantsPage(source, "trending")
}

// Recent returns the MyInstants recent page
func (source *MyInstants) Recent() *InstantsPage {
	return NewInstantsPage(source, "recent")
}

// Query returns an MyInstants query page
func (source *MyInstants) Query(name string) *InstantsPage {
	return NewInstantsPage(source, fmt.Sprintf("search/?name=%s", name))
}

// SelectInstants returns nodes to be considered for parsing Instants
func (source *MyInstants) SelectInstants(doc *goquery.Document) *goquery.Selection {
	// select all nodes of class "instant"
	return doc.Find(".instant")
}

// MakeInstants gets all instants inside collection
func (source *MyInstants) MakeInstants(sel *goquery.Selection) []Instant {
	instants := make([]Instant, 0, len(sel.Nodes))
	sel.Each(func(i int, sel *goquery.Selection) {
		instant := MyInstant{}

		// get title from link
		instant.name = sel.Find("a").Text()
		// use button's onclick inline script to get media location
		onclick, ok := sel.Find(".small-button").Attr("onclick")
		if ok {
			// remove surrounding play('')
			onclick = strings.TrimPrefix(onclick, "play('")
			onclick = strings.TrimSuffix(onclick, "')")
			// prefix media location with baseURL
			instant.audioURL = source.URL() + onclick

			// add to results
			instants = append(instants, instant)
		}
	})

	return instants
}
