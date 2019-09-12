package ref

import (
	"errors"
	"reflect"
	"strconv"
	"time"
)

type ErrCode int

const (
	ErrPointerToSliceRequired  = 1
	ErrSliceRequired           = 2
	ErrPointerToStructRequired = 3
	ErrStructRequired          = 4
	ErrTimeRequired            = 5
	ErrNilPointer              = 6
	ErrUnsupportedFieldType    = 7
	ErrUnexpected              = 8
)

var MapErr map[ErrCode]error = map[ErrCode]error{
	ErrPointerToSliceRequired:  errors.New(`A pointer to slice is required (*[]*Struct). `),
	ErrSliceRequired:           errors.New(`A slice is required ([]*Struct). `),
	ErrPointerToStructRequired: errors.New(`A pointer to struct is required (*Struct). `),
	ErrStructRequired:          errors.New(`A struct is required. `),
	ErrNilPointer:              errors.New(`Nil pointer dereference. `),
	ErrTimeRequired:            errors.New(`A time in format "1/2/2006 3:04:05 PM" is required`),
	ErrUnsupportedFieldType:    errors.New(`Unsupported field data type. `),
	ErrUnexpected:              errors.New(`Unexpected error. `),
}

type Merchant struct {
	Id        int64
	Name      string
	CreatedAt time.Time
}

// input a struct object (not pointer)
func GetStructFieldNames(obj interface{}) ([]string, error) {
	result := []string{}
	objType := reflect.TypeOf(obj)
	if objType.Kind() != reflect.Struct {
		return nil, MapErr[ErrStructRequired]
	}
	for i := 0; i < objType.NumField(); i++ {
		result = append(result, objType.Field(i).Name)
	}
	return result, nil
}

// args: destination must be *[]*struct
// return: bigErr, rowErrs
func ReadData(input []map[string]string, destination interface{}) (
	error, []error) {
	// destV is a value *[]*struct
	destV := reflect.ValueOf(destination)
	if destV.Kind() != reflect.Ptr {
		return MapErr[ErrPointerToSliceRequired], nil
	}
	if destV.IsNil() {
		return MapErr[ErrNilPointer], nil
	}
	// dest is a value []*struct
	dest := destV.Elem()
	// slice is a type []*struct
	slice := dest.Type()
	if slice.Kind() != reflect.Slice {
		return MapErr[ErrSliceRequired], nil
	}
	baseStructPtr := slice.Elem()
	if baseStructPtr.Kind() != reflect.Ptr {
		return MapErr[ErrPointerToStructRequired], nil
	}
	// base is ref.type Struct
	baseStruct := baseStructPtr.Elem()
	if baseStruct.Kind() != reflect.Struct {
		return MapErr[ErrStructRequired], nil
	}
	rowErrors := []error{}
	for _, row := range input {
		structPtr := reflect.New(baseStruct)
		var rowErr error
		for inputFieldName, inputFieldValue := range row {
			if inputFieldValue == "" {
				continue
			}
			field := structPtr.Elem().FieldByName(inputFieldName)
			if !field.IsValid() {
				continue
			}
			rowErr = ParseValueString(field, inputFieldValue)
			if rowErr != nil {
				break
			}
		}
		dest.Set(reflect.Append(dest, structPtr))
		rowErrors = append(rowErrors, rowErr)
	}
	return nil, rowErrors
}

func ParseValueString(field reflect.Value, inputS string) error {
	switch field.Kind() {
	case reflect.String:
		field.SetString(inputS)
		return nil
	case reflect.Int64:
		int64Value, err := strconv.ParseInt(inputS, 10, 64)
		if err != nil {
			float64Value, err := strconv.ParseFloat(inputS, 64)
			if err != nil {
				return err
			}
			field.SetInt(int64(float64Value))
			return nil
		}
		field.SetInt(int64Value)
		return nil
	case reflect.Float64:
		float64Value, err := strconv.ParseFloat(inputS, 64)
		if err != nil {
			return err
		}
		field.SetFloat(float64Value)
		return nil
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(inputS)
		if err != nil {
			return err
		}
		field.SetBool(boolValue)
		return nil
	case reflect.Struct: // time.Time field
		switch field.Interface().(type) {
		case time.Time:
			timeValue, err := time.Parse(time.RFC3339, inputS)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(timeValue))
			return nil
		default:
			return MapErr[ErrTimeRequired]
		}
	default:
		return MapErr[ErrUnsupportedFieldType]
	}
}
