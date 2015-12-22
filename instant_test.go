package moneinstant

import "testing"
import "github.com/victormoneratto/moneinstant/Godeps/_workspace/src/github.com/stretchr/testify/assert"
import "math/rand"

func TestInstants(t *testing.T) {
	myInstants := MyInstants{}

	// pages know to be valid
	validPages := []*InstantsPage{
		myInstants.Home(),
		myInstants.Trending(),
		myInstants.Recent(),
		myInstants.Query("ceccon")}

	for _, page := range validPages {
		// test First
		first, err := page.First()
		assert.NoError(t, err, "error on First()")
		assert.NotNil(t, first, "First() returns nil")

		// test At (random)
		index := rand.Intn(page.NumSelected())
		randInstant, err := page.At(index)
		assert.NoError(t, err, "error on At()")
		assert.NotNil(t, randInstant, "At() returns nil")

		// test All
		instants, err := page.All()
		assert.NoError(t, err, "error on All()")
		assert.NotEmpty(t, instants, "no instants found")
		assert.Equal(t, instants[0], first, "instants[0] and first are different")
	}

	var invalidPage *InstantsPage
	// test invalid First
	first, err := invalidPage.First()
	assert.Error(t, err, "no error for invalid page")
	assert.Nil(t, first, "invalid first should be nil")

	// test invalid At
	at0, err := invalidPage.At(0)
	assert.Error(t, err, "no error for invalid page")
	assert.Nil(t, at0, "invalid At(0) should be nil")

	// test invalid All
	all, err := invalidPage.All()
	assert.Error(t, err, "no error for invalid page")
	assert.Empty(t, all, "invalid All() should be empty")
}
