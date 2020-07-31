//package bookstore
package main

import (
	"fmt"
	"math"
)

var price = 800

var discount = map[int]float64{1:1, 2:0.05, 3:0.1, 4:0.2, 5:0.25}

// Cost fun
func Cost(books []int) (result int) {
	var total = map[int]int{}
	var values []int
	for _, val := range books {
		total[val]++
	}
	result = math.MaxInt32
	//for ii :=0; ii < len(books); ii++ {
		for i := 1; i <= 5; i++ {
			for !isEmpty(total) {
				values = append(values, getGroup(total, i))
			}
			if x:= calculateTotal(values); x < result {
				result = x
			}
		}
	//}
	return
}

func calculateTotal(input []int) (total int) {
	fmt.Println(input)
	for _, val := range input {
		total += int(float64(val)  * ( discount[val] * float64(price) ) * 100)
	}
	return 
}

func getGroup(input map[int]int, limit int) (total int) {
	for key, val := range input {
		if val != 0  && total <= limit{
			input[key]--
			total++
		}
	}
	return
}
func isEmpty(input map[int]int) bool {
	for _, val := range input {
		if val != 0 {
			return false
		}
	}
	return true
}


// it's not calculating the optimal price yet 
func main() {

	//description: "Two groups of four is cheaper than groups of five and three",
	//basket:      []int{1, 1, 2, 3, 4, 4, 5, 5},
	//var x = map[int]int{1:1,2:0,3:0,4:0,5:2}   
	// result 5120
	var y = []int{1, 1, 2, 3, 4, 4, 5, 5}
	fmt.Println(Cost(y))
}