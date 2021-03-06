package mock

import (
	"crypto/rand"
	"encoding/binary"
	"math"
)

// RandInt Returns an int64 between min and max.
func RandInt(min, max int64) int64 {
	off := min
	size := max - min
	return int64(randUint64()%uint64(size)) + off
}

// RandIntStep Returns an int64 whose step distance between min and max is step.
func RandIntStep(min, max, step int64) int64 {
	off := min
	sub := max - min
	size := sub / step
	return int64(randUint64()%uint64(size))*step + off
}

// RandUint Returns an uint64 between min and max.
func RandUint(min, max uint64) uint64 {
	off := min
	size := max - min
	return randUint64()%size + off
}

// RandUintStep Returns an uint64 whose step distance between min and max is step.
func RandUintStep(min, max, step uint64) uint64 {
	off := min
	sub := max - min
	size := sub / step
	return (randUint64()%size)*step + off
}

// RandFloat Returns an float64 between min and max.
func RandFloat(min, max float64) float64 {
	off := min
	size := max - min
	return randFloat64()*size + off
}

// RandFloatStep Returns an float64 whose step distance between min and max is step.
func RandFloatStep(min, max, step float64) float64 {
	off := min
	sub := max - min
	size := int64(sub / step)
	return float64(randUint64()%uint64(size))*step + off
}

func randUint64() uint64 {
	var buf [8]byte
	rand.Read(buf[:])
	return binary.BigEndian.Uint64(buf[:])
}

func randFloat64() float64 {
	f := float64(randUint64()>>1) / (1 << 63)
	if f == 1 {
		return randFloat64()
	}
	return f
}

func compareInt(a, b int64) (min, max int64) {
	if a < b {
		return a, b
	}
	return b, a
}

func compareUint(a, b uint64) (min, max uint64) {
	if a < b {
		return a, b
	}
	return b, a
}

func compareFloat(a, b float64) (min, max float64) {
	if a < b {
		return a, b
	}
	return b, a
}

func maxUint(bit int) uint64 {
	switch bit {
	default:
		return 0
	case 8:
		return math.MaxUint8
	case 16:
		return math.MaxUint16
	case 32:
		return math.MaxUint32
	case 64:
		return math.MaxUint64
	}
}

func maxInt(bit int) int64 {
	switch bit {
	default:
		return 0
	case 8:
		return math.MaxInt8
	case 16:
		return math.MaxInt16
	case 32:
		return math.MaxInt32
	case 64:
		return math.MaxInt64
	}
}

func minInt(bit int) int64 {
	switch bit {
	default:
		return 0
	case 8:
		return math.MinInt8
	case 16:
		return math.MinInt16
	case 32:
		return math.MinInt32
	case 64:
		return math.MinInt64
	}
}

func maxFloat(bit int) float64 {
	switch bit {
	default:
		return 0
	case 32:
		return math.MaxFloat32
	case 64:
		return math.MaxFloat64
	}
}

func minFloat(bit int) float64 {
	switch bit {
	default:
		return 0
	case 32:
		return -math.MaxFloat32
	case 64:
		return -math.MaxFloat64
	}
}
