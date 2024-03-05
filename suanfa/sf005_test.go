package suanfa

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) error {
	t := reflect.TypeOf(str).Elem()
	if t.Kind() != reflect.Struct {
		return errors.New("不是结构体类型")
	}

	for i := 0; i < t.NumField(); i++ {
		tagString := t.Field(i).Tag.Get("info")
		fmt.Println(tagString)
		docTagString := t.Field(i).Tag.Get("doc	")
		fmt.Println(docTagString)
	}

	return nil
}

func TestSuanfa005(t *testing.T) {
	var u resume
	findTag(&u)
	return
}
