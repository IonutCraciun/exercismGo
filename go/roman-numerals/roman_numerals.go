package romannumerals

import "errors"

/*
Symbol	I	V	X	L	C	D	M
Value	1	5	10	50	100	500	1000
*/

var num = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
var sym = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}

// ToRomanNumeral function
func ToRomanNumeral(number int) (string, error) {
	var result string
	if number <= 0 || number > 3000 {
		return "", errors.New("error")
	}
	for i := len(num) - 1; number > 0; i-- {
		for number/num[i] > 0 {
			result += sym[i]
			number -= num[i]
		}
	}
	return result, nil
}
