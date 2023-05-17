package withoutComparison

import (
	"algoritmi/algoritmi/implementedAlgorithms/dataStructures/queue"
	util "algoritmi/algoritmi/implementedAlgorithms/sort"
)

func BucketSort(data []*util.Record, b int) {
	// b chiave maggiore
	tmp := make([]*queue.RecordQueue, b+1) // slice di puntatori a code di record
	for i := 0; i <= b; i++ {              // creazione di una coda per ogni chiave in data
		tmp[i] = queue.CreateRecordQueue()
	}
	for ind, _ := range data { // riempimento bucket accodando ogni elemento nella coda associata alla chiave
		x := data[ind].Key
		tmp[x].Enqueue(data[ind])
	}
	index := 0
	for i := 0; i <= b; i++ { // Riordinamento slice data, viene svuotata una coda per volta nell'array passato come argomento
		for !tmp[i].IsEmpty() {
			data[index] = tmp[i].Dequeue()
			index++
		}
	}
}
