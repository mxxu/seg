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
	log.Print("segment string: " + s)
	// 先将string类型拆分成单字
	// go对unicode类型支持的很好，range就能自动搞定
	var words = splitStringToRunes(s)
	// 初始化辅助数据结构segflags
	var segflags = make([]SegFlag, len(*words))
	for i := range segflags {
		segflags[i].mincut = i+1
		segflags[i].head   = -1
	}
	// 
	for i, _ := range *words {
		var nd  = &(tree.root)
		for j := i; j < minInt(len(*words), tree.maxTokenLen+i); j++ {
			n, err := nd.SearchWord(&(*words)[j])
			if (err != nil) {
				//log.Printf("search word %c from %c failed", (*words)[j], nd.word)
				break
			}
			if (n.isEnd) {
				//log.Printf("token is end: %c, i=%d, j=%d", (*words)[i:j+1], i, j)
				if (segflags[j].head < 0) {
					segflags[j].head = i
				}
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
		// 分词词典中找不到，单字成词
		// TODO：多位的数字需要组合成一个词，比如“32场演唱会”
		if (segflags[i].head < 0) {
			segflags[i].head = i
		}
	}
	var numSegs = 0
	for i := len(segflags)-1; i >= 0; i = segflags[i].head - 1 {
		numSegs++
	}
	var ret = make([]string, numSegs)
	var j = numSegs-1
	for i := len(segflags)-1; i >= 0; i = segflags[i].head - 1 {
		ret[j] = fmt.Sprintf("%c", (*words)[segflags[i].head:i+1])
		j--
	}
	log.Print(ret)
	return nil
}