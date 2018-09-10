package goserbench

import (
	"math/rand"
	"time"
)

//go:generate -command genxdr go run ../../calmh/xdr/cmd/genxdr/main.go
//go:generate genxdr -o structdefxdr_generated.go structdefxdr.go
type XDRA struct {
	Name     string
	BirthDay int64
	Phone    string
	Siblings int
	Spouse   bool
	Money    uint64
}

func NewXDRA() interface{} {
	return &XDRA{
		Name:     randString(16),
		BirthDay: time.Now().Unix(),
		Phone:    randString(10),
		Siblings: rand.Intn(5),
		Spouse:   rand.Intn(2) == 1,
		Money:    rand.Uint64(),
	}
}
