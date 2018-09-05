package mock

import (
	"net"
	"testing"
	"time"
)

type T1 struct {
	Fr float32 `mock:"regexp,\\d{2}\\.\\d"`
	F0 float32 `mock:"range"`
	F1 float64 `mock:"range,20"`
	F2 float64 `mock:"range,-20,1"`
	F3 float64 `mock:"range,-30.5,0,2"`

	Ir int   `mock:"regexp,\\d{2}"`
	I0 int8  `mock:"range"`
	I1 int16 `mock:"range,20"`
	I2 int32 `mock:"range,-20,1"`
	I3 int64 `mock:"range,-30,100,2"`

	Ur uint   `mock:"regexp,\\d{2}"`
	U0 uint8  `mock:"range"`
	U1 uint16 `mock:"range,20"`
	U2 uint32 `mock:"range,20,30"`
	U3 uint64 `mock:"range,30,100,2"`

	Email  string    `mock:"email"`
	Domain string    `mock:"domain"`
	Name   string    `mock:"name"`
	UUID   string    `mock:"uuid"`
	Time   time.Time `mock:"time"`
	IPv4   net.IP    `mock:"ipv4"`
	IPv6   net.IP    `mock:"ipv6"`
	URL    string    `mock:"url"`
	Info   string    `mock:"text"`
}

func TestMock(t *testing.T) {
	var tt = make([]T1, 4)
	_, err := Mock(&tt)
	if err != nil {
		t.Fatal(err)
	}
}
