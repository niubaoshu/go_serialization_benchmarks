package goserbench

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

const num = 1000

var comparer = NewComparer()

func init() {
	comparer.Add(reflect.TypeOf((*ProtoBufA)(nil)).Elem(), func(i interface{}, i2 interface{}) bool {
		ia := i.(ProtoBufA)
		ib := i2.(ProtoBufA)
		return *ia.BirthDay == *ib.BirthDay && *ia.Name == *ib.Name && *ia.Phone == *ib.Phone && *ia.Money == *ib.Money && *ia.Siblings == *ib.Siblings
	})
}

func randString(l int) string {
	buf := make([]byte, l)
	for i := 0; i < (l+1)/2; i++ {
		buf[i] = byte(rand.Intn(256))
	}
	return fmt.Sprintf("%x", buf)[:l]
}

func generate(new func() interface{}, num int) []interface{} {
	a := make([]interface{}, num)
	for i := 0; i < num; i++ {
		a[i] = new()
	}
	return a
}

func bench(s Serializer, randNew func() interface{}, validate bool) *serializeBenchResault {
	ret := &serializeBenchResault{
		serializer: s,
	}
	data := generate(randNew, num)
	ser := make([][]byte, num)
	var sumlen int
	for i, d := range data {
		var buf []byte
		buf, ret.merr = s.Marshal(d)
		if validate && ret.merr != nil {
			return ret
		}
		t := make([]byte, len(buf))
		sumlen += len(buf)
		copy(t, buf)
		ser[i] = t
	}
	ret.length = sumlen / num
	ret.marshalResult = testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s.Marshal(data[i%num])
		}
	})
	g := randNew()
	ret.unmarshalResult = testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := s.Unmarshal(ser[i%num], g)
			if err != nil && ret.unmerr == nil {
				ret.unmerr = fmt.Errorf("%s failed to unmarshal: %s", s, err.Error())
				return
			}
			if validate {
				e := data[i%num]
				if !comparer.DeepEqual(e, g) && ret.unmerr == nil {
					e, g := indirect(e), indirect(g)
					ret.unmerr = fmt.Errorf("\n exp type = %T; value = %+v;\n got type = %T; value = %+v; \n", e, e, g, g)
					b.StartTimer()
					return
				}
			}
		}
	})
	return ret
}

func indirect(i interface{}) interface{} {
	v := reflect.ValueOf(i)
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	return v.Interface()
}
