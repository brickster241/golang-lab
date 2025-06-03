package adventofcode

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isReportSafeWithoutDampener(report []int) bool {
	if len(report) <= 1 {
		return true
	}

	is_incr := (report[1]- report[0]) > 0
	
	for i:= 1; i < len(report); i++ {
		
		curr_diff := report[i] - report[i - 1]
		if is_incr == false {
			curr_diff = report[i - 1] - report[i]
		}
		if (curr_diff >= 1 && curr_diff <= 3) == false {
			return false
		}
	}
	return true

}

func isReportSafeWithDampener(rpt []int) bool {
	if len(rpt) <= 2 {
		return true
	}
	// Assume we are going increasing sequence
	pref_check := make([]bool, len(rpt))
	suff_check := make([]bool, len(rpt))

	pref_check[0] = true
	for i:=1; i<len(rpt); i++ {
		diff := rpt[i] - rpt[i - 1]
		if (diff >= 1 && diff <= 3) {
			pref_check[i] = pref_check[i - 1]
		} else {
			pref_check[i] = false
		}
	}
	suff_check[len(rpt) - 1] = true
	for i:=len(rpt) - 2; i >= 0; i-- {
		diff := rpt[i + 1] - rpt[i]
		if (diff >= 1 && diff <= 3) {
			suff_check[i] = suff_check[i + 1]
		} else {
			suff_check[i] = false
		}
	}
	
	// CHECK FOR EACH INDEX
	for i:=0; i< len(rpt); i++ {
		// CORNER CHECK
		if i == 0 {
			if (suff_check[0] || suff_check[1]) {
				return true
			}
		} else if i == len(rpt) - 1 {
			if pref_check[len(rpt) - 1] || pref_check[len(rpt) - 2] {
				return true
			}
		} else {
			// i is in the middle and dampener not required
			if pref_check[i] && suff_check[i] {
				return true
			}
			// remove i and check 
			diff := rpt[i + 1] - rpt[i - 1]
			if pref_check[i - 1] && suff_check[i + 1] && diff >= 1 && diff <= 3 {
				return true
			}
		}
	}
	return false
}

func adventDay2() {
	dat, err := os.ReadFile("advent-of-code/inputs/input_day2.txt")

	// Check whether there is data present
	if err != nil {
		fmt.Printf("Error Occured While Parsing File : %s", err)
		return
	}

	inputs := strings.Split(string(dat), "\n")
	wd_safe_count := 0	// Without Dampener safe count
	d_safe_count := 0 // With Dampener safe count

	for _, v := range inputs {
		report_st := strings.Fields(v)
		report := make([]int, 0)
		for _, v := range report_st {
			val, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("Error Occured : %s", err)
				return
			}
			report = append(report, val)
		}
		rev_report := make([]int, len(report))
		for idx, _ := range report {
			rev_report[idx] = report[len(report) - 1 - idx]
		}
		if isReportSafeWithoutDampener(report) {
			wd_safe_count += 1
		}
		if isReportSafeWithDampener(report) || isReportSafeWithDampener(rev_report) {
			d_safe_count += 1
		}
	}
	fmt.Println("Safe Reports Without Dampener:", wd_safe_count)
	fmt.Println("Safe Reports With Dampener:", d_safe_count)

}