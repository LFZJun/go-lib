package htmlutil

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"testing"
)

var foo = `<html>
<body>

<table width="100%" border="1">
    <tr>
        <td rowspan="2">one</td>
        <td>two</td>
        <td>three</td>
    </tr>
    <tr>
        <td colspan="2">February</td>
    </tr>
</table>

</body>
</html>`

func BenchmarkParseTable(b *testing.B) {
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader([]byte(foo)))
	table := doc.Find("table")
	for i := 0; i < b.N; i++ {
		parseTable(table)
	}
}

func BenchmarkParseTable2(b *testing.B) {
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
