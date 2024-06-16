package scripts

import (
	"math/rand"
)

func randomElement(r *rand.Rand, arr []string) string {
	return arr[r.Intn(len(arr))]
}
