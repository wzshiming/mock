package mock

import (
	"encoding"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/wzshiming/crun"
)

const (
	tagFlag       = "mock"
	keyRangeFlag  = "range"
	keyRegexpFlag = "regexp"
	keyDomainFlag = "domain"
	keyURLFlag    = "url"
	keyEmailFlag  = "email"
	keyUUIDFlag   = "uuid"
	keyNameFlag   = "name"
	keyIPv4Flag   = "ipv4"
	keyIPv6Flag   = "ipv6"
	keyTimeFlag   = "time"
	keyWordFlag   = "word"
	keyTextFlag   = "text"
)

// Mock Inject mock data into the structure
func Mock(v interface{}) (interface{}, error) {
	return MockTag(v, "")
}

// MockTag Inject mock data into the structure by tag
func MockTag(v interface{}, tag string) (interface{}, error) {
	ret, err := mock(tag, reflect.ValueOf(v))
	if err != nil {
		return nil, err
	}
	return ret.Interface(), nil
}

var typTextUnmarshaler = reflect.TypeOf(new(encoding.TextUnmarshaler)).Elem()

func mock(tag string, val reflect.Value) (reflect.Value, error) {
	typ := val.Type()

	if typ.Implements(typTextUnmarshaler) {
		if val.IsNil() {
			return mock(tag, reflect.New(typ.Elem()))
		}
		err := assignImplements(tag, val)
		if err != nil {
			return reflect.Value{}, err
		}
		return val, nil
	}
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
			data, err := mock(tagVal, v.Addr())
			if err != nil {
				return reflect.Value{}, err
			}
			newVal.Field(i).Set(data.Elem())
		}
		return newVal, nil
	case reflect.Array:
		newVal := reflect.New(typ).Elem()
		num := val.Len()
		for i := 0; i != num; i++ {
			v := val.Index(i)
			data, err := mock(tag, v.Addr())
			if err != nil {
				return reflect.Value{}, err
			}
			newVal.Index(i).Set(data.Elem())
		}
		return newVal, nil
	case reflect.Slice:
		num := val.Len()
		newVal := reflect.MakeSlice(typ, num, val.Cap())
		for i := 0; i != num; i++ {
			v := val.Index(i)
			data, err := mock(tag, v.Addr())
			if err != nil {
				return reflect.Value{}, err
			}
			newVal.Index(i).Set(data.Elem())
		}
		return newVal, nil
	}
	return reflect.Value{}, fmt.Errorf("Error: There are unsupported kinds: %s", kind.String())
}

func assignImplements(tag string, val reflect.Value) error {
	ret, err := tagString(tag)
	if err != nil {
		return err
	}

	v, _ := val.Interface().(encoding.TextUnmarshaler)
	return v.UnmarshalText([]byte(ret))
}

func randString(tag string) (string, error) {
	reg, err := crun.Compile(tag)
	if err != nil {
		return "", err
	}
	ret := reg.Rand()
	return ret, nil
}

func tagString(tag string) (string, error) {
	data := strings.SplitN(tag, ",", 2)
	method := data[0]
	data = data[1:]
	switch method {
	case keyRegexpFlag:
		if len(data) == 0 {
			return "", nil
		}
		tag = data[0]
		return randString(tag)
	case keyDomainFlag:
		return RandDomain(), nil
	case keyURLFlag:
		return RandURL(), nil
	case keyEmailFlag:
		return RandEmail(), nil
	case keyUUIDFlag:
		return RandUUID(), nil
	case keyNameFlag:
		return RandName(), nil
	case keyWordFlag:
		return RandWord(), nil
	case keyTextFlag:
		return RandText(), nil
	case keyIPv4Flag:
		return RandIPv4(), nil
	case keyIPv6Flag:
		return RandIPv6(), nil
	case keyTimeFlag:
		return RandTime(time.RFC3339), nil
	case keyRangeFlag:
		// No action
	}
	return "", fmt.Errorf("Error: There are unsupported tag: %s", tag)
}

func assignString(tag string, val reflect.Value) error {
	ret, err := tagString(tag)
	if err != nil {
		return err
	}
	val.SetString(ret)
	return nil
}

func assignFloat(tag string, bits int, val reflect.Value) error {
	min := minFloat(bits)
	max := maxFloat(bits)
	data := strings.Split(tag, ",")
	method := data[0]
	data = data[1:]
	switch method {
	case keyRegexpFlag:
		if len(data) == 0 {
			return nil
		}
		tag = data[0]
		d, err := randString(tag)
		if err != nil {
			return err
		}
		ret, err := strconv.ParseFloat(d, bits)
		if err != nil {
			return err
		}
		val.SetFloat(ret)
		return nil
	case keyRangeFlag:
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
		case 0:
			ret = RandFloat(0, 1)
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
	return fmt.Errorf("Error: There are unsupported tag: %s", tag)
}

func assignInt(tag string, bits int, val reflect.Value) error {
	min := minInt(bits)
	max := maxInt(bits)
	data := strings.Split(tag, ",")
	method := data[0]
	data = data[1:]
	switch method {
	case keyRegexpFlag:
		if len(data) == 0 {
			return nil
		}
		tag = data[0]
		d, err := randString(tag)
		if err != nil {
			return err
		}
		ret, err := strconv.ParseInt(d, 0, 0)
		if err != nil {
			return err
		}
		val.SetInt(ret)
		return nil
	case keyRangeFlag:
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
		case 0:
			ret = RandInt(0, max)
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
	return fmt.Errorf("Error: There are unsupported tag: %s", tag)
}

func assignUint(tag string, bits int, val reflect.Value) error {
	max := maxUint(bits)
	var min uint64
	data := strings.Split(tag, ",")
	method := data[0]
	data = data[1:]
	switch method {
	case keyRegexpFlag:
		if len(data) == 0 {
			return nil
		}
		tag = data[0]
		d, err := randString(tag)
		if err != nil {
			return err
		}
		ret, err := strconv.ParseUint(d, 0, 0)
		if err != nil {
			return err
		}
		val.SetUint(ret)
		return nil
	case keyRangeFlag:
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
		case 0:
			ret = RandUint(0, max)
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
	return fmt.Errorf("Error: There are unsupported tag: %s", tag)
}
