package fetcher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFetcher(t *testing.T) {
	const n = 2

	//given a fetcher with capacity of n
	NewFetcher(n)
	fetcher := singleton

	//count of elements in the full fetcher must be n
	fillTheChan()
	assert.Len(t, singleton.sem, n)

	//call the constructor one more time with different capacity
	NewFetcher(3)
	//object must be the same
	assert.Equal(t, fetcher, singleton)
}

func fillTheChan() {
	for {
		select {
		case singleton.sem <- struct{}{}:
		default:
			return
		}
	}
}
