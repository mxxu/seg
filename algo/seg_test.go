package algo

import (
	//"fmt"
	"testing"
)

func TestSeg(t *testing.T) {
	fd := FileDict {
		filepath: "./dict.txt",
		linenum: 0,
		maxLine: -1,
		processor: &SogouInput_Processor{},
	}
	tree := New(fd.LoadDict())
	
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
		"结婚的和尚未结婚的",
		"梁启超生前住在这里",
		"费孝通向人大常委会提交书面报告",
		"粮食不卖给八路军",
		"阿拉斯加遭强暴风雪袭击致xx人死亡",
		"把手抬起来",
		"他说的确实在理",
	}
	for _, test := range tests {
		Segment(test, tree)
	}
}