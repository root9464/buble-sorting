/*bubble sorting*/
package main

import (
	"fmt"
	"math/rand"
)

func buble(array []int) []int {
	cheker := true
	length := len(array)

	// Массив в алгоритме считается отсортированным. При первой замене доказывается обратное и запускается еще одна итерация.
	// Цикл останавливается, когда все пары элементов в массиве пропускаются без замен

	for cheker {
		cheker = false

		for i := 1; i < length; i++ {

			// если порядок не верный
			if array[i-1] > array[i] {

				// меняем местами элементы
				array[i], array[i-1] = array[i-1], array[i]

				// продолжаем сортировку
				cheker = true
			}

		}

	}
	return array
}
func randomArr(n int) []int {
	array := rand.Perm(n)
	return array
}

func main() {
	a := randomArr(50)
	fmt.Println(a)
	fmt.Println(buble(a))
}
