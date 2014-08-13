package algo

import (
	"os"
	"log"
	"bufio"
	"strings"
)

type DictLoader interface {
	LoadDict() chan string
}

// line processor interface
type LineProcessor interface {
	ProcessLine(line string) string
}

// 具体的行处理类型

/*
格式：#开头的是注释行，否则：繁体 简体 音标 英文释义，空格分割

一日千里 一日千里 [yi1 ri4 qian1 li3] /lit. one day, a thousand miles (idiom); rapid progress/
一日為師，終身為父 一日为师，终身为父 [yi1 ri4 wei2 shi1 , zhong1 shen1 wei2 fu4] /lit. teacher for one day, father for ever (idiom)/
一旦 一旦 [yi1 dan4] /in case (sth happens)/if/once (sth happens, then...)/when/in a short time/in one day/

*/
type CEDICT_Processor struct {}
func (p *CEDICT_Processor) ProcessLine(line string) string {
	if line[0] == '#' {
		return ""
	}
	line = strings.Trim(line, "\n")
	arr := strings.Split(line, " ")
	return arr[1]
}

/*
格式：词 \t 词频 \t 词性（可选）
一个	818357166	
我们	770027797	PRON,
*/
type SogouLab_Processor struct {}
func (p *SogouLab_Processor) ProcessLine(line string) string {
	line = strings.Trim(line, "\n")
	arr := strings.Split(line, "\t")
	return arr[0]
}

/*
格式：一行一个词 
吖吖
吖啶
阿Q正传
*/
type SogouInput_Processor struct {}
func (p *SogouInput_Processor) ProcessLine(line string) string {
	return strings.Trim(line, "\n")
}

// file based dict, need a line processor
type FileDict struct {
	filepath  string
	linenum   int
	maxLine   int
	processor LineProcessor
}

func (fd *FileDict) LoadDict() chan string {
	ch := make(chan string)
	
	file, err := os.Open(fd.filepath)
	if err != nil {
		log.Fatal("open file failed: ", fd.filepath)
	}
	
	bufReader := bufio.NewReader(file)
	
	go func(chan string) {
		for {
			line, err := bufReader.ReadString('\n')
			if err != nil {
				break
			}
			word := fd.processor.ProcessLine(line)
			if (len(word) > 0) {
				fd.linenum++
				if (fd.maxLine > 0 && fd.linenum > fd.maxLine) {
					break
				}
				ch <- word
			}
		}
		close(ch)
	} (ch)
	return ch
}