package hello

import (
	"rsc.io/quote/v3"
)

// Hello is func
func Hello() string {
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}
