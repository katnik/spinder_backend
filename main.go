package main

import (
	"fmt"
	"github.com/katnik/spinder_backend/usecases"
	"math/rand"
	"time"
)

type Globals struct {
	Sitting  []int
	Standing []int
}

func generateRandomPermutation(n int) []int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Perm(n)
}

func main() {

	Glob := Globals{
		Sitting:  generateRandomPermutation(10),
		Standing: generateRandomPermutation(10),
	}

	t := new(usecases.TableInteractor)

	t.Sitting = Glob.Sitting
	t.Standing = Glob.Standing

	rounds := 10

	for i := 0; i < rounds; i++ {
		t.Questions = generateRandomPermutation(100)[15:25]
		fmt.Println("Round:", i+1)
		ts := t.CreateTables(10, i)

		for _, r := range ts {
			fmt.Println("table:", r.ID, "Sitting player:", r.Sitting, "Standing Player:", r.Standing, "Current Question:", r.Current)
		}
		fmt.Println("===================")

	}

}
