package main

import "fmt"

func main() {

	fmt.Println(fatorial(5))

}

// 1) fatorial recursivo
func fatorial(n int) int {

	if n == 0 || n == 1 {
		return 1
	}
	return n * fatorial(n-1)
}

// 2) conta numeros - recursivo
func contaNumeros(value, limit int) {
	if value > limit {
		return
	}
	fmt.Println(value)
	contaNumeros(value+1, limit)
}

// 3) algoritmo de ordenacao de lista/slice, o conhecido bubble sort.
func bubbleSortRecursivo(arr []int, n int) {
	if n == 1 {
		return
	}

	swapped := false

	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			swap(&arr, i, i+1)
			swapped = true
		}

	}

	if !swapped {
		return
	}

	bubbleSortRecursivo(arr, n-1)

}

// 4) busca binaria recursiva
// retorna a localizacao de um dado x de um array/slice, se n estiver presente retorna -1
func binarySearch(arr []int, x, left, right int) int {
	if right >= left {
		mid := left + (right-left)/2

		if arr[mid] == x {
			return mid
		}

		if arr[mid] > x {
			return binarySearch(arr, x, left, mid-1)
		}

		return binarySearch(arr, x, mid+1, right)
	}

	return -1
}

//funcao auxiliar swap que troca posicoes de um array
func swap(arr *[]int, i, j int) {
	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
}
