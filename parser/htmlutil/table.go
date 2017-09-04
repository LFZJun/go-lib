package htmlutil

import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

type (
	content struct {
		Text  string
		Value string
	}
	table [][]string
	kv map[string]string
)

func (k *kv) Set(key string, v string) {
	(*k)[key] = v
}

func (k *kv) Get(key string) string {
	if v, has := (*k)[key]; has {
		return v
	}
	return ""
}

func tx2Content(tx *goquery.Selection) content {
	value, _ := tx.Find("input").Attr("value")
	txt := strings.TrimSpace(tx.Text())
	return content{
		Text:  txt,
		Value: value,
	}
}

func (c content) String() string {
	if v := strings.TrimSpace(c.Value); v != "" {
		return v
	}
	return strings.TrimSpace(c.Text)
}

func (t table) FirstHeadKV() (p []kv) {
	if len(t) < 2 {
		return
	}
	headline := t[0]
	for i := 1; i < len(t); i++ {
		pp := kv{}
		for j := 0; j < len(t[i]); j++ {
			pp.Set(headline[j], t[i][j])
		}
		p = append(p, pp)
	}
	return
}

func ParseTable(table *goquery.Selection) (matrix table) {
	trs := table.Find("tr")
	switch length := trs.Length(); length {
	case 0:
		return
	case 1:
		matrix = make([][]string, 1)
		var rowOne []string
		// trs = tr
		trs.Children().Each(func(col int, tx *goquery.Selection) {
			c := tx2Content(tx)
			if colspan, has := tx.Attr("colspan"); has {
				colspanInt, _ := strconv.Atoi(colspan)
				for i := 0; i < colspanInt; i++ {
					rowOne = append(rowOne, c.String())
				}
			} else {
				rowOne = append(rowOne, c.String())
			}
		})
		matrix[0] = rowOne
	default:
		matrix = make([][]string, length)
		trs.Each(func(row int, tr *goquery.Selection) {
			tr.Children().Each(func(col int, tx *goquery.Selection) {
				rowspanInt, colspanInt := 1, 1
				c := tx2Content(tx)
				if rowspan, has := tx.Attr("rowspan"); has {
					rowspanInt, _ = strconv.Atoi(rowspan)
				}
				if colspan, has := tx.Attr("colspan"); has {
					colspanInt, _ = strconv.Atoi(colspan)
				}
				for i := rowspanInt - 1; i >= 0; i-- {
					it := row + i
					for j := 0; j < colspanInt; j++ {
						matrix[it] = append(matrix[it], c.String())
					}
				}
			})
		})
	}
	return
}
