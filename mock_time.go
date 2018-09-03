package mock

import (
	"math/rand"
	"time"
)

// RandTime Returns a random time
func RandTime(format string) string {
	unix := time.Now().UnixNano()
	num := int64(rand.Int()) % unix
	return time.Unix(0, num).In(time.Local).Format(format)
}
