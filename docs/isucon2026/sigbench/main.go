package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

func sig(key []byte, id string, exp int64) []byte {
	value := []byte(strconv.FormatInt(exp, 10) + ":" + id)
	mac := hmac.New(sha256.New, key)
	sum := make([]byte, 0, sha256.Size)
	for range 1000 {
		mac.Reset()
		mac.Write(value)
		sum = mac.Sum(sum[:0])
		value = append(value[:0], sum...)
	}
	return value
}

func main() {
	key := make([]byte, 32)
	const n = 2000
	start := time.Now()
	for i := range n {
		sig(key, "0b8f9a0e-1c2d-4e5f-8a9b-"+strconv.Itoa(100000000000+i), 1793500000)
	}
	el := time.Since(start)
	fmt.Printf("署名 %d 件: %v（1 件 %.3f ms、1 コアで %.0f 件/秒）\n", n, el, float64(el.Microseconds())/n/1000, n/el.Seconds())
}
