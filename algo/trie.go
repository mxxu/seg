package algo

import fmt "fmt"
//import bytes "bytes"
import log "log"
import os "os"
import bufio "bufio"
import strings "strings"
import errors "errors"

type Word  rune
type Token []Word

type Trie struct {
	root  node
	numTokens int
	maxTokenLen int
}

type node struct {
	word  Word
	isEnd bool
	children []*node
}

func (n *node) Print(indent string) {
	fmt.Println(fmt.Sprintf("%sword: %#U, isEnd: %t, children: ", indent, n.word, n.isEnd))
	for _, child := range n.children {
		child.Print(indent + "    ")
	}
}

func (t *Trie) Print() {
	fmt.Println(fmt.Sprintf("numTokens: %d, maxTokenLen: %d, nodes:", t.numTokens, t.maxTokenLen))
	t.root.Print("")
}

func (t *Trie) Search(token *Token) (*node, error) {
	return t.root.Search(token)
}

func (n *node) Search(token *Token) (*node, error) {
	currnode := n
	for _, word := range *token {
		ok := false
		for _, child := range currnode.children {
			if child.word == word {
				ok = true
				currnode = child
				break
			}
		}
		if (!ok) {
			return nil, errors.New("not found")
		}
	}
	return currnode, nil
}

func (t *Trie) SearchWord(word *Word) (*node, error) {
	return t.root.SearchWord(word)
}

func (n *node) SearchWord(word *Word) (*node, error) {
	for _, child := range n.children {
		if child.word == *word {
			return child, nil
		}
	}
	return nil, errors.New("cannot find word")
}

func (t *Trie) AddToken(token *Token) {
	var newToken = false
	currnode := &t.root
	for _, word := range(*token) {
		ok := false
		for _, child := range(currnode.children) {
			if child.word == word {
				ok = true
				currnode = child
				break
			}
		}
		if (!ok) {
			child := &node{word, false, []*node{}}
			currnode.children = append(currnode.children, child)
			currnode = child
			newToken = true
		}
	}
	currnode.isEnd = true
	if (t.maxTokenLen < len(*token)) {
		t.maxTokenLen = len(*token)
	}
	if (newToken) {
		t.numTokens = t.numTokens + 1
	}
}

func splitStringToRunes(s string) *Token {
	var token Token
	for _, value := range s {
		token = append(token, Word(value))
	}
	return &token
}

func LoadDictToTrie(filepath string) *Trie {
	var tree Trie
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("open file failed: ", filepath)
	}
	bufReader := bufio.NewReader(file)
	for {
		line, err := bufReader.ReadString('\n')
		if err != nil {
			break
		}
		
		if line[0] == '#' {
			continue
		}
		line = strings.Trim(line, "\n")
		arr := strings.Split(line, " ")
		if tree.numTokens > 100000 {
			break
		}
		tree.AddToken(splitStringToRunes(arr[1]))
	}
	return &tree
}

func test() {
	/*
	var tree Trie
	const hanyu = "汉语文字"
	const hanyu2 = "汉字文学"
	tree.AddToken(splitStringToRunes(hanyu))
	tree.AddToken(splitStringToRunes(hanyu2))
	*/
	var dict_filepath = "/Users/xumaoxing/Downloads/cedict_ts.u8"
	tree := LoadDictToTrie(dict_filepath)
	//tree.Print()
	
	fmt.Println(tree.Search(splitStringToRunes("21kr")))
	n, _ := tree.Search(splitStringToRunes("USB"))
	n.Print("")
	n2, _ := n.Search(splitStringToRunes("记忆"))
	n2.Print("")
	
	//n3, _ := n.Search(splitStringToRunes("haha")); n3.Print("")
	
	//var n = node{Word('从'), true, nil}
	//log.Println(n.String())
}