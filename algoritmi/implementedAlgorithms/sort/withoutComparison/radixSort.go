package withoutComparison

import (
	_func "algoritmi/algoritmi/implementedAlgorithms/func"
	util "algoritmi/algoritmi/implementedAlgorithms/sort"
	"math"
)

// !!!DA COMPLETARE!!!

func RadixSort(data []*util.Record) {
	index := 0.0
	max := _func.RicercaMax(*util.Keys(data))
	for math.Floor(float64(max)/math.Pow(float64(2), index)) != 0 {
		AdvancedBucketSort(data, max, int(index))
		index++
	}
}
