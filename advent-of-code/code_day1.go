package adventofcode

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)
func adventDay1() {
	dat, err := os.ReadFile("inputs/input_day1.txt")

	if err != nil {
		panic(err)
	}
	inputs := strings.Split(string(dat), "\n")
	count := 0
	Lefts := make([]int, 0)
	Rights := make([]int, 0)
	freq := make(map[int]int, 0)

	for ind, v := range inputs {
		// Split into left and right list
		row := strings.Fields(v)
		left, l_err := strconv.Atoi(row[0])
		right, r_err := strconv.Atoi(row[1])
		if l_err != nil || r_err != nil {
			fmt.Printf("Error Occured on Line : %d\n", ind)
			return
		}
		Lefts = append(Lefts, left)
		Rights = append(Rights, right)
		val, ok := freq[right]; if ok {
			freq[right] = val + 1
		} else {
			freq[right] = 1
		}
	}
	sort.Ints(Lefts)
	sort.Ints(Rights)
	similarity := 0
	
	for ind, _ := range Lefts {
		fmt.Printf("Index %d : (%d, %d)\n", ind, Lefts[ind], Rights[ind])
		if Lefts[ind] > Rights[ind] {
			count += Lefts[ind] - Rights[ind]
		} else {
			count += Rights[ind] - Lefts[ind]
		}
		val, ok := freq[Lefts[ind]]; if ok {
			similarity += val * Lefts[ind]
		}
	}

	fmt.Println("Total Count :", count)
	fmt.Println("Similarity  :", similarity)
}