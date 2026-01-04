package main
import "fmt"
func main() {
	arr := [6]int{12, 45, 7, 89, 23, 5}

	max := arr[0]
	min := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}

	fmt.Println("Maximum value:", max)
	fmt.Println("Minimum value:", min)

}
