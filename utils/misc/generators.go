package misc

import (
	"math/rand"
	"sync"
	"time"
)

//Singleton instance

type idGenerator struct {
	next int
}

var instance *idGenerator
var once sync.Once

func (g *idGenerator) Generate() int {
	g.next += 1
	return g.next
}

func GetInstance(size int) *idGenerator {
	once.Do(func() {
		instance = &idGenerator{size}
	})
	return instance
}

func generateRandomNumbers(n int) []int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Perm(n)
}
