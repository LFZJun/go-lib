package htmlutil

import (
	"container/list"
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

type Element struct {
	Row     int           `json:"row"`
	Col     int           `json:"col"`
	Text    string        `json:"text"`
	Rowspan int           `json:"rowspan"`
	Colspan int           `json:"colspan"`
	Index   *list.Element `json:"_"`
}

func (e Element) IsRowspan() bool {
	return e.Rowspan > 1
}

func (e Element) IsColspan() bool {
	return e.Colspan > 1
}

var solveColspan = func(ele *Element, m *[]*list.List) {
	newEle := (*m)[ele.Row].InsertAfter(&Element{Row: ele.Row, Col: ele.Col, Text: ele.Text, Rowspan: ele.Rowspan, Colspan: ele.Colspan - 1}, ele.Index)
	for c := newEle; c != nil; c = c.Next() {
		c.Value.(*Element).Col += 1
	}
}

var solveRowspan = func(ele *Element, m *[]*list.List) {
	offset := ele.Row + 1
	newEle := (*m)[offset].InsertBefore(&Element{Row: ele.Row, Col: ele.Col, Text: ele.Text, Rowspan: ele.Rowspan - 1, Colspan: 1}, ele.Index)
	for c := newEle; c != nil; c = c.Next() {
		c.Value.(*Element).Col += 1
	}
}

func ParseTable(table *goquery.Selection) (matrix [][]string) {
	trs := table.Find("tr")
	if trs.Length() == 0 {
		return
	}
	var m []*list.List
	trs.Each(func(row int, tr *goquery.Selection) {
		it := list.New()
		tr.Children().Each(func(col int, tx *goquery.Selection) {
			ele := Element{Row: row, Col: col, Text: tx.Text(), Rowspan: 1, Colspan: 1}
			if rowspan, has := tx.Attr("rowspan"); has {
				ele.Rowspan, _ = strconv.Atoi(rowspan)
			}
			if colspan, has := tx.Attr("colspan"); has {
				ele.Colspan, _ = strconv.Atoi(colspan)
			}
			ele.Index = it.PushBack(&ele)
		})
		m = append(m, it)
	})
	for _, v := range m {
		for head := v.Front(); head != nil; head = head.Next() {
			ele := head.Value.(*Element)
			if ele.IsRowspan() {
				solveRowspan(ele, &m)
			}
			if ele.IsColspan() {
				solveColspan(ele, &m)
			}
		}
	}
	matrix = make([][]string, len(m))
	for i := range m {
		matrix[i] = make([]string, 0, m[i].Len())
		for head := m[i].Front(); head != nil; head = head.Next() {
			matrix[i] = append(matrix[i], head.Value.(*Element).Text)
		}
	}
	return
}
