package util_test

import (
	"math/rand"
	"strings"
)

const (
	Digits = "0123456789"
	Upper  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lower  = "abcdefghijklmnopqrstuvwxyz"
)

type RG struct {
	sb *strings.Builder
}

func NewRandomGenerator() *RG {
	return &RG{&strings.Builder{}}
}

func NewRandomGeneratorWithSeed(seed int64) *RG {
	rand.Seed(seed)
	return NewRandomGenerator()
}

func (r *RG) String() string {
	return r.sb.String()
}
func (r *RG) Space() {
	r.sb.WriteByte(' ')
}

func (r *RG) NewLine() {
	r.sb.WriteByte('\n')
}

func (r *RG) Byte(b byte) {
	r.sb.WriteByte(b)
}

func (r *RG) Bytes(s string) {
	r.sb.WriteString(s)
}

func (r *RG) One() {
	r.sb.WriteString("1\n")
}
