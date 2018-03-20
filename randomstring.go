package randomstring

import (
	"errors"
	"math"
	"math/rand"
	"time"
)

const (
	CharsetAlphaNum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	CharsetNum      = "0123456789"
)

var (
	rsg *RandomStringGenerator
)

func init() {
	rsg, _ = NewGenerator(CharsetAlphaNum)
}

func Generate(n int) string {
	return rsg.Generate(n)
}

type RandomStringGenerator struct {
	charset       string
	charsetLength int
	letterIdxBits uint
}

func NewGenerator(charset string) (*RandomStringGenerator, error) {
	ret := &RandomStringGenerator{}
	return ret.WithCharset(charset)
}

func (rsg *RandomStringGenerator) WithCharset(c string) (*RandomStringGenerator, error) {
	rsg.charset = c
	rsg.charsetLength = len(c)
	letterIdxBits := math.Ceil(math.Log2(float64(rsg.charsetLength)))
	if letterIdxBits == 0 {
		return nil, errors.New("charset too long")
	}
	rsg.letterIdxBits = uint(letterIdxBits)
	return rsg, nil
}

func (rsg *RandomStringGenerator) Generate(n int) string {
	var letterIdxMask int64 = 1<<rsg.letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	var letterIdxMax = 63 / rsg.letterIdxBits          // # of letter indices fitting in 63 bits
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < rsg.charsetLength {
			b[i] = rsg.charset[idx]
			i--
		}
		cache >>= rsg.letterIdxBits
		remain--
	}

	return string(b)
}
