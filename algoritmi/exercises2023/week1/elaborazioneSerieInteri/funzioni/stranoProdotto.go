package funzioni

func StranoProdotto(numeri []int) int {
	res := 0
	for _, val := range numeri {
		if val%4 == 0 && val > 7 {
			if res == 0 {
				res = 1
			}
			res *= val
		}
	}
	return res
}
