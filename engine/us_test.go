package engine

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestUserAgent(t *testing.T) {
	index := len(m)
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(index)
	fmt.Println(x, m[x])

}
