package algo

import (
	"testing"
	//"fmt"
)

func TestDict(t *testing.T) {
	/*
	fd := FileDict{filepath: "./cedict_ts.u8", linenum: 0, maxLine: 10, processor: &CEDICT_Processor{}}
	tree := New(fd.LoadDict())
	
	fd2 := FileDict{filepath: "./SogouHanyuUtf8.txt", linenum: 0, maxLine: 10, processor: &SogouInput_Processor{}}
	tree.Update(fd2.LoadDict())
	
	fd3 := FileDict{filepath: "Freq/SogouLabDicUtf8.dic", linenum: 0, maxLine: 10, processor: &SogouLab_Processor{}}
	tree.Update(fd3.LoadDict())
	tree.Print()
	*/
	fd := FileDict {
		filepath: "./dict.txt",
		linenum: 0,
		maxLine: 10,
		processor: &SogouInput_Processor{},
	}
	New(fd.LoadDict()).Print()
}