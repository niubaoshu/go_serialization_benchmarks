package goserbench

import (
	"math/rand"
	"time"
)

type NoTimeA struct {
	Name     string
	BirthDay int64
	Phone    string
	Siblings int
	Spouse   bool
	Money    float64
}

func NewNoTimeA() interface{} {
	return &NoTimeA{
		Name:     randString(16),
		BirthDay: time.Now().UnixNano(),
		Phone:    randString(10),
		Siblings: rand.Intn(5),
		Spouse:   rand.Intn(2) == 1,
		Money:    rand.Float64(),
	}
}
