package pwd_generator

import (
	"math/rand"
	"time"
)

func ComplexGenerate(size int) *RandomResult {
	allSets = append(allSets, numbericSets...)
	allSets = append(allSets, lowerAlphabetSets...)
	allSets = append(allSets, upperAlphabetSets...)
	allSets = append(allSets, superSets...)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(allSets), func(i, j int) { allSets[i], allSets[j] = allSets[j], allSets[i] })
	destSet := getDestSet(size, allSets)
	return &RandomResult{len(*destSet), string(*destSet), time.Now()}
}
