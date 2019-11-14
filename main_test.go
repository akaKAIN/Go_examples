package main

import (
	"bufio"
	"examples/util"
	"strings"
	"bytes"
	"testing"
)

var testText = `1
2
3
4
4
5
6
7
`

var testTextOk = `1
2
3
4
5
6
7
`

var textFail = `1
2
1
`

func TestIt(t *testing.T){
	reader := bufio.NewReader(strings.NewReader(testText))
	writer := new(bytes.Buffer)
	err := util.Uniq(reader, writer)
	if err != nil {
		t.Error("test is Failed")
	}
	if writer.String() != testTextOk {
		t.Errorf("test is Failed -" +
			"result not correct\n%v - %v", writer.String(), testTextOk  )
	}
}

func TestError(t *testing.T){
	reader := bufio.NewReader(strings.NewReader(testText))
	writer := new(bytes.Buffer)
	if err := util.Uniq(reader, writer); err != nil {
		t.Error("test is Failed - error")
	}
}