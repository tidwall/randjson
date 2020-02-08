package randjson

import (
	"math/rand"
	"strconv"

	"github.com/tidwall/pretty"
	"github.com/tidwall/words"
)

// Options for Make()
type Options struct {
	// Pretty formats and indents the random json. Default true
	Pretty bool
	// Spread is the number of unique words to use. Default 1,000
	Words int
	// Rand is the random number generator to use. Default global rng
	Rand *rand.Rand
}

// DefaultOptions for Make()
var DefaultOptions = &Options{
	Pretty: true,
	Words:  1000,
	Rand:   nil,
}

// Append a random json document to dst. The depth param is the maximum nested
// depth of json document
func Append(dst []byte, depth int, opts *Options) []byte {
	if opts == nil {
		opts = DefaultOptions
	}
	var p float64
	if opts.Words > len(words.Words) {
		p = 1.0
	} else if opts.Words < 1 {
		p = 1 / float64(len(words.Words))
	} else {
		p = float64(opts.Words) / float64(len(words.Words))
	}
	s := int(float64(len(words.Words)) * p)
	t := (len(words.Words) / s)
	mark := len(dst)
	dst = appendRandObject(dst, opts.Rand, s, t, depth)
	if opts.Pretty {
		dst = append(dst[:mark], pretty.Pretty(dst[mark:])...)
	}
	return dst
}

// Make returns a random json document. The depth param is the maximum nested
// depth of json document
func Make(depth int, opts *Options) []byte {
	return Append(nil, depth, opts)
}

func randInt(rng *rand.Rand) int {
	if rng == nil {
		return rand.Int()
	}
	return rng.Int()
}
func appendRandString(dst []byte, rng *rand.Rand, s, t int) []byte {
	dst = append(dst, '"')
	dst = append(dst, words.Words[(randInt(rng)%s)*t]...)
	return append(dst, '"')
}
func appendRandAny(dst []byte, rng *rand.Rand, nested bool, s, t, d int) []byte {
	switch randInt(rng) % 7 {
	case 0:
		dst = appendRandString(dst, rng, s, t)
	case 1:
		if !nested {
			dst = appendRandAny(dst, rng, nested, s, t, d)
		} else {
			dst = append(dst, '[')
			if d > 1 {
				n := randInt(rng) % (d - 1)
				for i := 0; i < n; i++ {
					if i > 0 {
						dst = append(dst, ',')
					}
					dst = appendRandAny(dst, rng, false, s, t, d-1)
				}
			}
			dst = append(dst, ']')
		}
	case 2:
		if !nested {
			dst = appendRandAny(dst, rng, nested, s, t, d)
		} else {
			if d > 1 {
				d = randInt(rng) % (d - 1)
			}
			dst = appendRandObject(dst, rng, s, t, d)
		}
	case 3:
		dst = strconv.AppendFloat(dst,
			float64(randInt(rng)%10000)/100, 'f', 2, 64)
	case 4:
		dst = append(dst, "true"...)
	case 5:
		dst = append(dst, "false"...)
	case 6:
		dst = append(dst, "null"...)
	}
	return dst
}
func appendRandObject(dst []byte, rng *rand.Rand, s, t, d int) []byte {
	dst = append(dst, '{')
	for i := 0; i < d; i++ {
		if i > 0 {
			dst = append(dst, ',')
		}
		dst = appendRandString(dst, rng, s, t)
		dst = append(dst, ':')
		dst = appendRandAny(dst, rng, true, s, t, d-1)
	}
	return append(dst, '}')
}
