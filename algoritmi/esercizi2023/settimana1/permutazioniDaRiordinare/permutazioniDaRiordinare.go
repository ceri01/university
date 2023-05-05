package main

import (
	"fmt"
	"os"
)

type myStruct struct {
	name string
	key  int
}

func bubbleSort(slMyStruct []myStruct) {
	ln := len(slMyStruct)
	for {
		t := 0
		for i := 1; i < ln-1; i++ {
			if slMyStruct[i].key > slMyStruct[i-1].key {
				tmp := slMyStruct[i-1]
				slMyStruct[i-1] = slMyStruct[i]
				slMyStruct[i] = tmp
				t = i
			}
		}
		ln = t
		if t <= 0 {
			break
		}
	}
}

func main() {
	var limit int
	var key int
	var name string
	var sl []myStruct

	_, err := fmt.Scanf("%d", &limit)
	if err != nil {
		fmt.Println("Errore di input")
		os.Exit(-1)
	}

	for i := 0; i < limit; i++ {
		_, err := fmt.Scanf("%d %s", &key, &name)
		if err != nil {
			fmt.Println("Errore di input")
			os.Exit(-1)
		}

		sl = append(sl, myStruct{name: name, key: key})
	}

	bubbleSort(sl)

	for _, el := range sl {
		fmt.Println(el.key, " ", el.name)
	}

}
