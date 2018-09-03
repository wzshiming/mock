package mock

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/wzshiming/crun"
)

const tagFlag = "mock"

// Mock Inject mock data into the structure
func Mock(v interface{}) (interface{}, error) {
	ret, err := mock("", reflect.ValueOf(v))
	if err != nil {
		return nil, err
	}
	return ret.Interface(), nil
}

func mock(tag string, val reflect.Value) (reflect.Value, error) {
	typ := val.Type()
	kind := val.Kind()
	switch kind {
	case reflect.Ptr:
		if val.IsNil() {
			return mock(tag, reflect.New(typ.Elem()))
		}
		newVal, err := mock(tag, val.Elem())
		if err != nil {
			return reflect.Value{}, err
		}
		val.Elem().Set(newVal)
		return val, nil
	case reflect.String:
		newVal := reflect.New(typ).Elem()
		err := assignString(tag, newVal)
		if err != nil {
			return reflect.Value{}, err
		}
		return newVal, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		newVal := reflect.New(typ).Elem()
		bits := typ.Bits()
		err := assignInt(tag, bits, newVal)
		if err != nil {
			return reflect.Value{}, err
		}
		return newVal, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		newVal := reflect.New(typ).Elem()
		bits := typ.Bits()
		err := assignUint(tag, bits, newVal)
		if err != nil {
			return reflect.Value{}, err
		}
		return newVal, nil
	case reflect.Float32, reflect.Float64:
		newVal := reflect.New(typ).Elem()
		bits := typ.Bits()
		err := assignFloat(tag, bits, newVal)
		if err != nil {
			return reflect.Value{}, err
		}
		return newVal, nil
	case reflect.Struct:
		newVal := reflect.New(typ).Elem()
		num := typ.NumField()
		for i := 0; i != num; i++ {
			v := val.Field(i)
			f := typ.Field(i)
			tagVal := f.Tag.Get(tagFlag)
			if tagVal == "" {
				continue
			}
			data, err := mock(tagVal, v)
			if err != nil {
				return reflect.Value{}, err
			}
			newVal.Field(i).Set(data)
		}
		return newVal, nil
	case reflect.Array:
		newVal := reflect.New(typ).Elem()
		num := val.Len()
		for i := 0; i != num; i++ {
			v := val.Index(i)
			data, err := mock(tag, v)
			if err != nil {
				return reflect.Value{}, err
			}
			newVal.Index(i).Set(data)
		}
		return newVal, nil
	case reflect.Slice:
		num := val.Len()
		newVal := reflect.MakeSlice(typ, num, val.Cap())
		for i := 0; i != num; i++ {
			v := val.Index(i)
			data, err := mock(tag, v)
			if err != nil {
				return reflect.Value{}, err
			}
			newVal.Index(i).Set(data)
		}
		return newVal, nil
	}
	return reflect.Value{}, fmt.Errorf("Error: There are unsupported kinds: %s", kind.String())
}

func assignString(tag string, val reflect.Value) error {

	reg, err := crun.Compile(tag)
	if err != nil {
		return err
	}
	ret := reg.Rand()
	val.SetString(ret)
	return nil
}

func assignFloat(tag string, bits int, val reflect.Value) error {
	min := minFloat(bits)
	max := maxFloat(bits)
	data := strings.Split(tag, ",")
	dataFloat := make([]float64, 0, len(data))
	for _, v := range data {
		d, err := strconv.ParseFloat(v, bits)
		if err != nil {
			return err
		}
		dataFloat = append(dataFloat, d)
	}

	var ret float64
	switch len(dataFloat) {
	case 1:
		max, _ = compareFloat(dataFloat[0], max)
		ret = RandFloat(0, max)
	case 2:
		min0, max0 := compareFloat(dataFloat[0], dataFloat[1])
		_, min = compareFloat(max0, min)
		max, _ = compareFloat(min0, max)
		ret = RandFloat(min, max)
	case 3:
		min0, max0 := compareFloat(dataFloat[0], dataFloat[1])
		_, min = compareFloat(max0, min)
		max, _ = compareFloat(min0, max)
		ret = RandFloatStep(min, max, dataFloat[2])
	default:
		return fmt.Errorf("Error: wrong number of arguments:%s", tag)
	}
	val.SetFloat(ret)
	return nil
}

func assignInt(tag string, bits int, val reflect.Value) error {
	min := minInt(bits)
	max := maxInt(bits)
	data := strings.Split(tag, ",")
	dataInt := make([]int64, 0, len(data))
	for _, v := range data {
		d, err := strconv.ParseInt(v, 0, 0)
		if err != nil {
			return err
		}
		dataInt = append(dataInt, d)
	}

	var ret int64
	switch len(dataInt) {
	case 1:
		max, _ = compareInt(dataInt[0], max)
		ret = RandInt(0, max)
	case 2:
		min0, max0 := compareInt(dataInt[0], dataInt[1])
		_, min = compareInt(min0, min)
		max, _ = compareInt(max0, max)
		ret = RandInt(min, max)
	case 3:
		min0, max0 := compareInt(dataInt[0], dataInt[1])
		_, min = compareInt(min0, min)
		max, _ = compareInt(max0, max)
		ret = RandIntStep(min, max, dataInt[2])
	default:
		return fmt.Errorf("Error: wrong number of arguments:%s", tag)
	}
	val.SetInt(ret)
	return nil
}

func assignUint(tag string, bits int, val reflect.Value) error {
	max := maxUint(bits)
	var min uint64
	data := strings.Split(tag, ",")
	dataUint := make([]uint64, 0, len(data))
	for _, v := range data {
		d, err := strconv.ParseUint(v, 0, 0)
		if err != nil {
			return err
		}
		dataUint = append(dataUint, d)
	}

	var ret uint64
	switch len(dataUint) {
	case 1:
		max, _ = compareUint(dataUint[0], max)
		ret = RandUint(0, max)
	case 2:
		min0, max0 := compareUint(dataUint[0], dataUint[1])
		_, min = compareUint(min0, min)
		max, _ = compareUint(max0, max)
		ret = RandUint(min, max)
	case 3:
		_, min = compareUint(dataUint[0], min)
		max, _ = compareUint(dataUint[1], max)
		ret = RandUintStep(min, max, dataUint[2])
	default:
		return fmt.Errorf("Error: wrong number of arguments:%s", tag)
	}
	val.SetUint(ret)
	return nil
}
