package algo

import fmt "fmt"
import errors "errors"
import "unicode"

type Word  []rune
type Token []Word

func (w *Word) Print() {
	for _, r := range *w {
		fmt.Printf("%c", r)
	}
}

func (t *Token) Print(sep string) {
	for _, w := range *t {
		w.Print()
		fmt.Printf(sep)
	}
}

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
			if isWordEqual(child.word, word) {
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

func (t *Trie) SearchWord(word Word) (*node, error) {
	return t.root.SearchWord(word)
}

func (n *node) SearchWord(word Word) (*node, error) {
	for _, child := range n.children {
		if isWordEqual(child.word, word) {
			return child, nil
		}
	}
	return nil, errors.New("cannot find word")
}

func isWordEqual(a, b Word) bool {
	if (len(a) != len(b)) {
		return false
	}
	for i, v := range a {
		if (v != b[i]) {
			return false
		}
	}
	return true
}

func (t *Trie) AddToken(token *Token) {
	var newToken = false
	currnode := &t.root
	for _, word := range(*token) {
		ok := false
		for _, child := range(currnode.children) {
			if isWordEqual(child.word, word) {
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

func bothAlphaNum(a, b rune) bool {
	if (unicode.IsLetter(a) && unicode.IsLetter(b)) {
		return true
	} else if (unicode.IsDigit(a) && unicode.IsDigit(b)) {
		return true
	}
	return false
}

func isAlphaNum(a rune) bool {
	return unicode.IsNumber(a) || (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z')
}

func splitStringToRunes(s string) *Token {
	var token Token
	var inAlphaNum bool = false
	var word Word
	for i, value := range s {
		isBegin := false
		if (isAlphaNum(value)) {
			//fmt.Printf("value %c is alpha num\n", value)
			if (!inAlphaNum) {
				inAlphaNum = true
				isBegin = true
			}
		} else {
			isBegin = true
			inAlphaNum = false
		}
		if (isBegin && i > 0) {
			//fmt.Printf("%c", token)
			token = append(token, word)
			word = nil
		}
		word = append(word, value)
		//fmt.Printf("%d, %c, %c\n", i, value, word)
	}
	token = append(token, word)
	return &token
}

func New(ch chan string) *Trie {
	var tree Trie
	for phrase := range ch {
		tree.AddToken(splitStringToRunes(phrase))
	}
	return &tree
}

func (t *Trie) Update(ch chan string) {
	for phrase := range ch {
		t.AddToken(splitStringToRunes(phrase))
	}
}