package goserbench

import (
	"math/rand"
	"time"
)

//go:generate msgp -o msgp_gen.go -io=false -tests=false
//easyjson:json
type A struct {
	Name     string
	BirthDay time.Time
	Phone    string
	Siblings int
	Spouse   bool
	Money    float64
}

func NewA() interface{} {
	return &A{
		Name:     randString(16),
		BirthDay: time.Now(),
		Phone:    randString(10),
		Siblings: rand.Intn(5),
		Spouse:   rand.Intn(2) == 1,
		Money:    rand.Float64(),
	}
}
