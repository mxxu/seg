package algo

import (
	//"fmt"
	"testing"
)

func TestSeg(t *testing.T) {
	var dict_filepath = "./cedict_ts.u8"
	tree := LoadDictToTrie(dict_filepath)
	//tree.Print()
	Segment("我们在野生动物园玩", tree)
	Segment("P民K歌", tree)
	Segment("中华人民共和国", tree)
	Segment("叔叔亲了我妈妈也亲了我", tree)
	var tests = []string {
		"这个把手该换了",
		"别把手放在我的肩膀上",
		"共同创造美好的新世纪",
		"我不喜欢日本和服",
		"工信处女干事每月经过下属科室都要亲口交代24口交换机等技术性器件的安装工作",
	}
	for _, test := range tests {
		Segment(test, tree)
	}
}