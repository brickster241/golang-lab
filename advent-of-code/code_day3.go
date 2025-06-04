package adventofcode

import (
	"fmt"
	"os"
	"strconv"
)

// mul()
func mulCount(input_st string) int {
	n := len(input_st)
	idx := 0
	count := 0
	isEnabled := true
	for idx < n {
		if idx >= n - 4 {
			break
		}
		if input_st[idx: idx + 4] == "do()" {
			isEnabled = true
			idx += 4
		} else if idx + 7 < n && input_st[idx: idx + 7] == "don't()" {
			isEnabled = false
			idx += 7
		} else if input_st[idx: idx + 4] == "mul(" {
			// check if there is some number present after it until there is a ,
			for j := idx + 4; j < n; j++ {
				if input_st[j] <= 57 && input_st[j] >= 48 {
					fmt.Printf("Num1 Digit Found : %v\n", input_st[j])
				} else if input_st[j] == ',' {
					// Store the number from idx + 4 to j - 1
					num1, err1 := strconv.Atoi(input_st[idx + 4: j])
					if err1 != nil {
						fmt.Printf("Cannot convert to number : %s\n", input_st[idx + 4: j])
						idx = j + 1
						break
					}
					// Valid Number1, now search for comma and valid Number 2
					
					for k := j + 1; k < n; k++ {
						if input_st[k] <= 57 && input_st[k] >= 48 {
							fmt.Printf("Num2 Digit Found : %v\n", input_st[k])
						} else if input_st[k] == ')' {
							// Store the number from j + 1 to k
							num2, err2 := strconv.Atoi(input_st[j + 1: k])
							if err2 != nil {
								fmt.Printf("Cannot convert to number : %s\n", input_st[j + 1: k])
								idx = k + 1
								j = n
								break
							}
							
							if isEnabled {
								count += (num1 * num2)
							}
							idx = k + 1
							j = n
							break
						} else {
							idx = k + 1
							j = n
							break
						}
					}
				} else {
					idx = j + 1
					break
				}
			}
			
		} else {
			idx += 1
		}
	}
	return count
}

func adventDay3() {
	dat, err := os.ReadFile("inputs/input_day3.txt")

	if err != nil {
		panic(err)
	}
	input := string(dat)

	fmt.Println(mulCount(input))
}