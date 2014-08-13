package algo

import (
	//"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	word_chan := make(chan string)
	go func() {
		word_chan <- "as"
		word_chan <- "df"
		close(word_chan)
	} ()
	tree := New(word_chan)
	tree.Print()
}