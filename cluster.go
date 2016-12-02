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

func buildFeatureArray(m *Model, a, b Expr) (Features, error) {
	f := make(Features, len(a)+len(b))
	realCount := 0

	for word := range a {
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

	for word := range b {
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

func doCluster(m *Model, main, related Expr, shifter meanshift.Shifter) {

	fmt.Println("============================")
	//fmt.Println(buildFeatureArray(m, e))
	features, _ := buildFeatureArray(m, main, related)

	ms := meanshift.New(features, shifter, 0.10, 10)
	//ms := meanshift.New(features, meanshift.NewUniform(1.0), 0.10, 8)
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

// Cluster similar words
func Cluster(m *Model, main, related Expr) {
	fmt.Println("\n\n*******************************************\n******************************************\narticle:")
	//ushifter = []meanshift.Shifter{meanshift.NewUniform(1.0), meanshift.NewUniform(1.5), meanshift.NewUniform(2.0)}
	//gshifter = []dd
	//fmt.Println("\n*******median:")
	//doCluster(m, main, related, meanshift.NewUniform(1.0))
	//doCluster(m, main, related, meanshift.NewUniform(2.0))
	//doCluster(m, main, related, meanshift.NewUniform(3.0))
	//fmt.Println("\n*******gaus:")
	//doCluster(m, main, related, meanshift.NewTruncGauss(0.55, 4.0))
	//doCluster(m, main, related, meanshift.NewTruncGauss(0.6, 4.0))
	//doCluster(m, main, related, meanshift.NewTruncGauss(0.65, 4.0))
}
