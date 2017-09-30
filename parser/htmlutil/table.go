package htmlutil

import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
	"container/list"
)

type Ele struct {
	Row     int           `json:"row"`
	Col     int           `json:"col"`
	Text    string        `json:"text"`
	Rowspan int           `json:"rowspan"`
	Colspan int           `json:"colspan"`
	Index   *list.Element `json:"_"`
}

func (e Ele) shallowClone() Ele {
	return e
}

func (e Ele) IsRowspan() bool {
	return e.Rowspan > 1
}

func (e Ele) IsColspan() bool {
	return e.Colspan > 1
}

func insert(list2 *list.List, index int, value interface{}) *list.Element {
	if index < list2.Len() {
		ele := list2.Front()
		for i := 0; i < index; i++ {
			ele = ele.Next()
		}
		return list2.InsertBefore(value, ele)
	}
	return list2.PushBack(value)
}

func solveColspan(ele *Ele, m *[]*list.List) {
	currentRow := (*m)[ele.Row]
	cp := ele.shallowClone()
	cp.Colspan -= 1
	newEle := currentRow.InsertAfter(&cp, ele.Index)
	for c := newEle; c != nil; c = c.Next() {
		c.Value.(*Ele).Col += 1
	}
}

func solveRowspan(ele *Ele, m *[]*list.List) {
	offset := ele.Row + 1
	nextRow := (*m)[offset]
	cp := ele.shallowClone()
	cp.Row = offset
	cp.Rowspan -= 1
	cp.Colspan = 1
	newEle := insert(nextRow, cp.Col, &cp)
	for ; newEle != nil; newEle = newEle.Next() {
		newEle.Value.(*Ele).Col += 1
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
			ele := Ele{Row: row, Col: col, Text: tx.Text(), Rowspan: 1, Colspan: 1}
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
			ele := head.Value.(*Ele)
			if ele.IsColspan() {
				solveColspan(ele, &m)
			}
			if ele.IsRowspan() {
				solveRowspan(ele, &m)
			}
		}
	}
	matrix = make([][]string, len(m))
	for i := range m {
		matrix[i] = make([]string, 0, m[i].Len())
		for head := m[i].Front(); head != nil; head = head.Next() {
			matrix[i] = append(matrix[i], strings.TrimSpace(head.Value.(*Ele).Text))
		}
	}
	return
}
