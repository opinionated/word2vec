package word2vec

import (
	"fmt"
	"github.com/biogo/cluster/meanshift"
)

// Feature a
type Feature struct {
	id   string
	data []float64
}

// Features of the data
type Features []Feature

// Len as
func (f Features) Len() int {
	return len(f)
}

// Values a
func (f Features) Values(i int) []float64 {
	return []float64(f[i].data)
}

func buildFeatureArray(m *Model, e Expr) (Features, error) {
	f := make(Features, len(e))
	realCount := 0

	for word := range e {
		vec, ok := m.getWord(word)
		if !ok {
			continue
		}

		castVec := make([]float64, len(vec))
		for i := range vec {
			castVec[i] = float64(vec[i])
		}

		f[realCount] = Feature{word, castVec}
		realCount++
	}

	ret := make(Features, realCount)
	for i := 0; i < realCount; i++ {
		ret[i] = f[i]
	}

	return ret, nil
}

// Cluster similar words
func Cluster(m *Model, main, related Expr) {

	fmt.Println("\narticle:")
	//fmt.Println(buildFeatureArray(m, e))
	features, _ := buildFeatureArray(m, e)

	ms := meanshift.New(features, meanshift.NewTruncGauss(0.65, 3), 0.50, 8)
	err := ms.Cluster()
	if err != nil {
		panic(err)
	}

	for _, c := range ms.Centers() {
		fmt.Println("")
		for _, i := range c.Members() {
			f := features[i]
			fmt.Println(f.id)
		}
	}
}
