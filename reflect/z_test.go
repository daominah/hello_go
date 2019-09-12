package ref

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func Test01(t *testing.T) {
	var fields []string
	var err error
	fields, err = GetStructFieldNames(Merchant{})
	// fmt.Println("fields", fields, err)
	if len(fields) != 3 {
		t.Error(err)
	}
	m1 := 5
	fields, err = GetStructFieldNames(m1)
	if err != MapErr[ErrStructRequired] {
		t.Error(err)
	}
}

func Test2(t *testing.T) {
	a1 := &Merchant{}
	i := interface{}(a1)
	s2 := []*Merchant{a1, &Merchant{}}
	i2 := interface{}(s2)
	iV := reflect.ValueOf(i)
    i2V := reflect.ValueOf(i2)
	fmt.Println("value.Elem is dereference", iV.Elem())
    _ = i2V
	typ1 := reflect.TypeOf(i)
    fmt.Println("type.Elem typ1", typ1, typ1.Elem())
	typ2 := reflect.TypeOf(i2)
	fmt.Println("type.Elem typ2", typ2, typ2.Elem())
}

func Test03(t *testing.T) {
	//
	input := []map[string]string{
		map[string]string{"Id": "1", "Name": "Name1", "CreatedAt": "2019-01-01T12:23:34Z"},
		map[string]string{"Id": "2", "Namex": "Name2", "CreatedAt": "2019-01-01 12:23:34"},
		map[string]string{"Id": "3a", "Name": "Name", "CreatedAt": "2019-01-01T12:23:34Z"},
	}
	merchants := []*Merchant{}
	err, rowErrs := ReadData(input, &merchants)
	fmt.Println("err", err)
	for i, e := range rowErrs {
	    fmt.Println("err", i, e)
    }
	temp, _ := json.MarshalIndent(merchants, "", "    ")
	fmt.Println("merchants", string(temp))
}
