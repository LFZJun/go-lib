package htmlutil

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"testing"
	"container/list"
)

var foo = `<html>
<body>

<table width="100%" border="1">
    <tr>
        <td rowspan="2">one</td>
        <td>two</td>
        <td rowspan="2">three</td>
    </tr>
    <tr>
        <td>February</td>
    </tr>
</table>

</body>
</html>`

func BenchmarkParseTable(b *testing.B) {
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader([]byte(foo)))
	table := doc.Find("table")
	for i := 0; i < b.N; i++ {
		ParseTable(table)
	}
}

func TestParseTable(t *testing.T) {
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader([]byte(foo)))
	table := doc.Find("table")
	fmt.Println(ParseTable(table))
}

func TestInsert(t *testing.T) {
	l := list.New()
	insert(l, 0, 1)
	insert(l, 1, 2)
	for ele := l.Front();ele!=nil;ele = ele.Next() {
		fmt.Println(ele.Value)
	}
}
