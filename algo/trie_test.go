package algo

import (
	"fmt"
	"testing"
	"unicode"
	"unicode/utf8"
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
	
	fmt.Println(splitStringToRunes("我想1你234哈哈"))
	
	fmt.Println(unicode.IsLetter(rune('我')))
	r, size := utf8.DecodeRuneInString("我想")
	fmt.Println(unicode.IsDigit(r), size)
}