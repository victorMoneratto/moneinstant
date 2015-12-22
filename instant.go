package moneinstant

import (
	"errors"
	"fmt"

	"github.com/victormoneratto/moneinstant/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
)

// Errors getting instants
var (
	ErrNoInstant = errors.New("no instant found on page")
	ErrNilPage   = errors.New("page is nil")
)

// Source is a website provider of instants
type Source interface {
	URL() string
	Home() *InstantsPage
	Trending() *InstantsPage
	Recent() *InstantsPage
	Query(name string) *InstantsPage
	SelectInstants(*goquery.Document) *goquery.Selection
	MakeInstants(*goquery.Selection) []Instant
}

// InstantsPage is a where we look for Instants
type InstantsPage struct {
	Source    Source
	Selection *goquery.Selection
	Error     error
}

// NumSelected returns the number of nodes that were
// selected as instants by SelectInstants
func (p *InstantsPage) NumSelected() int {
	if p == nil {
		return 0
	}

	return len(p.Selection.Nodes)
}

// NewInstantsPage return an instants page
func NewInstantsPage(source Source, relativeURL string) *InstantsPage {
	// fetch page
	url := fmt.Sprintf("%s/%s", source.URL(), relativeURL)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		// save the error on the page instance so we can see it later.
		// we won't return errors, that way we can make a "stream"
		// or "builder" pattern
		return &InstantsPage{Error: err}
	}
	// select all instants in page
	sel := source.SelectInstants(doc)
	if len(sel.Nodes) == 0 {
		// no instant node was selected
		return &InstantsPage{Source: source, Selection: sel, Error: ErrNoInstant}
	}

	return &InstantsPage{Source: source, Selection: sel}
}

// Instant contain url and name of instant
type Instant interface {
	Name() string
	AudioURL() string
}

// All returns a slice with all instants from page
func (p *InstantsPage) All() ([]Instant, error) {
	if p == nil {
		return nil, ErrNilPage
	}
	return instantsFromSelection(p, p.Selection)
}

// At returns the instant at an specific index
func (p *InstantsPage) At(index int) (Instant, error) {
	if p == nil {
		return nil, ErrNilPage
	}

	results, err := instantsFromSelection(p, p.Selection.Eq(index))
	if len(results) == 0 {
		return nil, err
	}
	return results[0], err
}

// First returns the first instant from page
func (p *InstantsPage) First() (Instant, error) {
	return p.At(0)
}

// instantsFromSelection returns a slice with all instants in selection
func instantsFromSelection(p *InstantsPage, sel *goquery.Selection) ([]Instant, error) {
	// look for errors while selecting instants
	if p.Error != nil {
		return nil, p.Error
	}

	// get instants in selection
	results := p.Source.MakeInstants(sel)
	if len(results) == 0 {
		return nil, ErrNoInstant
	}

	return results, nil
}
