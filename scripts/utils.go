package scripts

import (
	"math/rand"
)

func randomElement(r *rand.Rand, arr []string) string {
	return arr[r.Intn(len(arr))]
}

func randomName(r *rand.Rand, firstNames []string, lastNames []string) string {
	return randomElement(r, firstNames) + " " + randomElement(r, lastNames)
}
