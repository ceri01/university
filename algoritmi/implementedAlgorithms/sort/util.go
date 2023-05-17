package sort

type Record struct {
	Key  int
	Data string
}

func Keys(data []*Record) *[]int {
	var keys []int
	for _, el := range data {
		keys = append(keys, el.Key)
	}
	return &keys
}
