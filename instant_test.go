package moneinstant

import "testing"

func TestNewHomePage(t *testing.T) {
	page, err := NewHomePage()
	if err != nil {
		t.Error(err)
	}

	if len(selectInstants(page.Document).Nodes) == 0 {
		t.Error("Home page has no supported instants")
	}
}

func TestNewInstantPage(t *testing.T) {
	page, err := NewInstantPage("carlos")
	if err != nil {
		t.Error(err)
	}

	if len(selectInstants(page.Document).Nodes) == 0 {
		t.Error("Query page has no supported instants")
	}
}

func TestGetRandomInstant(t *testing.T) {
	page, err := NewInstantPage("carlos")
	if err != nil {
		t.Error(err)
	}

	instant, err := page.GetRandomInstant()
	if err != nil {
		t.Error(err)
	}

	if instant.URL == "" {
		t.Error("URL is empty")
	}
}

func TestGetAllInstants(t *testing.T) {
	page, err := NewInstantPage("carlos")
	if err != nil {
		t.Error(err)
	}

	instants, err := page.GetAllInstants()
	if err != nil {
		t.Error(err)
	}

	if len(instants) == 0 {
		t.Error("no supported instants found")
	}
}
