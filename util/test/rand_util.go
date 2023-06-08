package util_test

import (
	"math/rand"
	"sort"
	"strconv"
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

func (r *RG) _int(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func (r *RG) IntOnly(min, max int) int {
	return r._int(min, max)
}

// Int generates a random int in range [min, max]
func (r *RG) Int(min, max int) int {
	v := r._int(min, max)
	r.sb.WriteString(strconv.Itoa(v))
	r.Space()
	return v
}

// Float generates a random float in range [min, max] with a fixed precision
func (r *RG) Float(min, max float64, precision int) float64 {
	v := min + rand.Float64()*(max-min)
	r.sb.WriteString(strconv.FormatFloat(v, 'f', precision, 64))
	r.Space()
	return v
}

// Str generates a random string with length in range [minLen, maxLen] and its chars in range [min, max]
func (r *RG) Str(minLen, maxLen int, min, max byte) string {
	l := r._int(minLen, maxLen)
	sb := &strings.Builder{}
	sb.Grow(l)
	for i := 0; i < l; i++ {
		sb.WriteByte(byte(r._int(int(min), int(max))))
	}
	s := sb.String()
	r.sb.WriteString(s)
	r.Space()
	return s
}

// StrInSet generates a random string with length in range [minLen, maxLen] and its chars in chars
func (r *RG) StrInSet(minLen, maxLen int, chars string) string {
	l := r._int(minLen, maxLen)
	sb := &strings.Builder{}
	sb.Grow(l)
	for i := 0; i < l; i++ {
		sb.WriteByte(chars[rand.Intn(len(chars))])
	}
	s := sb.String()
	r.sb.WriteString(s)
	r.Space()
	return s
}

// intSlice generates a int slices
func (r *RG) intSlice(size int, min, max int) []int {
	a := make([]int, size)
	for i := range a {
		a[i] = r._int(min, max)
	}
	return a
}

func (r *RG) intSliceInSet(size int, set []int) []int {
	a := make([]int, size)
	for i := range a {
		a[i] = set[rand.Intn(len(set))]
	}
	return a
}

func (r *RG) IntSliceInSet(size int, set []int) []int {
	a := r.intSliceInSet(size, set)
	for _, v := range a {
		r.sb.WriteString(strconv.Itoa(v))
		r.Space()
	}
	r.NewLine()
	return a
}

// IntSlice generates a random int slice with a fixed size and its values in range [min, max]
func (r *RG) IntSlice(size int, min, max int) []int {
	a := r.intSlice(size, min, max)
	for _, v := range a {
		r.sb.WriteString(strconv.Itoa(v))
		r.Space()
	}
	r.NewLine()
	return a
}

// IntSliceOrdered generates a random ordered int slice with a fixed size and its values in range [min, max]
func (r *RG) IntSliceOrdered(size int, min, max int, inc, unique bool) []int {
	var a []int
	if unique {
		a = r.uniqueSlice(size, min, max)
	} else {
		a = r.intSlice(size, min, max)
	}
	if inc {
		sort.Ints(a)
	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(a)))
	}
	for _, v := range a {
		r.sb.WriteString(strconv.Itoa(v))
		r.Space()
	}
	r.NewLine()
	return a
}

// IntMatrix generates a random int matrix with fixed row and col and its values in range [min, max]
func (r *RG) IntMatrix(row, col int, min, max int) [][]int {
	a := make([][]int, row)
	for i := range a {
		a[i] = r.intSlice(col, min, max)
	}
	for _, row := range a {
		for _, v := range row {
			r.sb.WriteString(strconv.Itoa(v))
			r.Space()
		}
		r.NewLine()
	}
	return a
}

// TODO: O(size) 做法 https://mivik.blog.luogu.org/the-art-of-randomness
func (r *RG) uniqueSlice(size int, min, max int) []int {
	if size > max-min+1 {
		panic("size is too large")
	}
	p := rand.Perm(max - min + 1)[:size]
	for i := range p {
		p[i] += min
	}
	return p
}
