package withoutComparison

import (
	"algoritmi/algoritmi/implementedAlgorithms/dataStructures/queue"
	util "algoritmi/algoritmi/implementedAlgorithms/sort"
	"fmt"
	"strconv"
)

// !!!DA COMPLETARE!!!

func AdvancedBucketSort(data []*util.Record, b int, pos int) {
	// b chiave maggiore
	tmp := make([]*queue.RecordQueue, b+1) // slice di puntatori a code di record
	for i := 0; i <= b; i++ {              // creazione di una coda per ogni chiave in data
		tmp[i] = queue.CreateRecordQueue()
	}
	fmt.Println(tmp)
	for ind, _ := range data { // riempimento bucket accodando ogni elemento nella coda associata alla chiave
		x := int(strconv.FormatInt(int64(data[ind].Key), b)[pos])
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
