package moneinstant

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"github.com/victormoneratto/moneinstant/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
)

// MyInstants base url
var MyInstantsURL = "http://www.myinstants.com"

// Page to look for instants
type Page struct {
	Document *goquery.Document
}

// Returns the MyInstants home page
func NewHomePage() (*Page, error) {
	return NewInstantPage("")
}

// Returns MyInstants page with query by name,
// or home page if name is empty
func NewInstantPage(name string) (*Page, error) {
	var url string
	if name == "" {
		url = MyInstantsURL
	} else {
		url = fmt.Sprintf("%s/search/?name=%s", MyInstantsURL, name)
	}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	return &Page{Document: doc}, nil
}

// URL and Name of am instant
type InstantResult struct {
	URL, Name string
}

// Errors getting instants
var (
	ErrNoInstant    = errors.New("no instant found on page")
	ErrParseInstant = errors.New("couldn't parse instant")
)

// Get random instant from page
func (p *Page) GetRandomInstant() (*InstantResult, error) {
	instants := selectInstants(p.Document)
	if len(instants.Nodes) == 0 {
		return nil, ErrNoInstant
	}

	index := rand.Intn(len(instants.Nodes))
	results := makeInstantResults(instants.Eq(index))
	if len(results) == 0 {
		return nil, ErrParseInstant
	}

	return results[0], nil
}

// Get All instants in page
func (p *Page) GetAllInstants() ([]*InstantResult, error) {
	instants := selectInstants(p.Document)
	if len(instants.Nodes) == 0 {
		return nil, ErrNoInstant
	}

	results := makeInstantResults(instants)
	if len(results) == 0 {
		return nil, ErrParseInstant
	}

	return results, nil
}

// Select all nodes to be considered for parsing Instants
func selectInstants(doc *goquery.Document) *goquery.Selection {
	return doc.Find(".instant")
}

// Parse instants in selection
func makeInstantResults(sel *goquery.Selection) []*InstantResult {
	results := make([]*InstantResult, 0, len(sel.Nodes))
	valid := make([]bool, len(sel.Nodes))

	sel.Find(".small-button").Each(func(i int, sel *goquery.Selection) {
		if url, ok := sel.Attr("onclick"); ok {
			valid[i] = true
			url = strings.Replace(url, "play('", MyInstantsURL+"/", 1)
			if strings.HasSuffix(url, "')") {
				url = url[:len(url)-2]
			}
			results = append(results, &InstantResult{URL: url})
		}
	})
	resultIndex := 0
	sel.Find("a").Each(func(i int, sel *goquery.Selection) {
		if valid[i] {
			results[resultIndex].Name = sel.Text()
			resultIndex++
		}
	})

	return results
}
