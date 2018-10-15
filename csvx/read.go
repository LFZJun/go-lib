package csvx

import (
	"encoding/csv"
	"github.com/kataras/iris/core/errors"
	"io"
	"reflect"
	"strconv"
)

const (
	Index string = "index"
	Head  string = "head"
)

type Reader struct {
	// 标题名称head和下表index
	nameHead map[string]int
	// 标题
	Head []string
	// csv Reader
	reader *csv.Reader
}

func NewReader(reader *csv.Reader) (*Reader, error) {
	r := Reader{
		reader: reader,
	}
	err := r.readHead()
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (r *Reader) readHead() error {
	record, err := r.reader.Read()
	if err != nil {
		return err
	}
	r.Head = record
	r.nameHead = make(map[string]int)
	for i := range record {
		r.nameHead[record[i]] = i
	}
	return nil
}

func (r *Reader) Read(dest interface{}) error {
	record, err := r.reader.Read()
	if err == io.EOF {
		return err
	}

	destValue := reflect.ValueOf(dest)
	if destValue.Kind() != reflect.Ptr {
		return errors.New("dest is not a Pointer")
	}
	destType := destValue.Type()
	classElem := destType.Elem()
	valueElem := destValue.Elem()
	for num := classElem.NumField() - 1; num >= 0; num-- {
		classInfo := classElem.Field(num)
		field := valueElem.Field(num)
		// 根据index tag来注入
		indexStr, ok := classInfo.Tag.Lookup(Index)
		if ok {
			err := r.fillIndexStr(field, record, indexStr)
			if err != nil {
				return err
			}
		}

		// 根据name tag来注入
		name, ok := classInfo.Tag.Lookup(Head)
		if ok {
			index, has := r.nameHead[name]
			if !has {
				continue
			}
			fillIndex(field, record, index)
		}
	}
	return nil
}

// 转换indexStr为int类型调用fillIndex
func (r *Reader) fillIndexStr(field reflect.Value, record []string, indexStr string) error {
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		return err
	}
	fillIndex(field, record, index)
	return nil
}

// 根据index提取value
func fillIndex(field reflect.Value, record []string, index int) {
	csvValue := record[index]
	field.Set(reflect.ValueOf(csvValue))
}
