package algo

import (
	"fmt"
	"log"
)

type SegFlag struct {
	mincut int
	head   int
}

func minInt(a, b int) int {
	if (a < b) {
		return a
	}
	return b
}

func Segment(s string, tree *Trie) []string {
	var words = splitStringToRunes(s)
	var segflags = make([]SegFlag, len(*words))
	for i := range segflags {
		segflags[i].mincut = i+1
		segflags[i].head   = -1
	}
	for i, _ := range *words {
		var nd  = &(tree.root)
		for j := i; j < minInt(len(*words), tree.maxTokenLen+i); j++ {
			//fmt.Printf("%c\n", (*words)[j])
			if (i > 0 && segflags[i-1].head == -1) {
				log.Printf("%c's last word is not word-end, not going on", (*words)[j])
				break
			}
			n, err := nd.SearchWord(&(*words)[j])
			if (err != nil) {
				log.Printf("search word %c from %c failed", (*words)[j], nd.word)
				break
			}
			if (n.isEnd) {
				log.Printf("token is end: %c, i=%d, j=%d", (*words)[i:j+1], i, j)
				
				var newcut int
				if (i == 0) {
					newcut = 1
				} else {
					newcut = segflags[i-1].mincut + 1
				}
				if (i == 0 || segflags[j].mincut > newcut) {
					segflags[j].mincut = newcut
					segflags[j].head   = i
				}
			}
			nd = n
		}
	}
	fmt.Println(segflags)
	var numSegs = 0
	for i := len(segflags)-1; i >= 0; i = segflags[i].head - 1 {
		numSegs++
	}
	fmt.Println(numSegs)
	var ret = make([]string, numSegs)
	var j = numSegs-1
	for i := len(segflags)-1; i >= 0; i = segflags[i].head - 1 {
		ret[j] = fmt.Sprintf("%c", (*words)[segflags[i].head:i+1])
		j--
	}
	fmt.Println(ret)
	return nil
}