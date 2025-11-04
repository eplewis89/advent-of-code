package fifteen

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/eplewis89/advent-of-code/bootstrap"
)

const year = 2015

func Solve() {
	fmt.Printf("===== AOC %d =====\n", year)

	day_one()
	day_two()
}

func day_one() {
	input, err := bootstrap.ReadInput(year, 1)

	if err != nil {
		log.Printf("failed to get data: %v", err)

		return
	}

	position := 0
	result := 0

	for i, c := range input {
		// part two
		if result < 0 {
			position = i
			break
		}

		switch c {
		case '(':
			result += 1
		case ')':
			result -= 1
		}
	}

	fmt.Printf("%d day %d result\n", year, 1)
	fmt.Println(position, result)
}

func day_two() {
	input, err := bootstrap.ReadInput(year, 2)

	if err != nil {
		log.Printf("failed to get data: %v", err)

		return
	}

	dims := strings.Fields(input)
	final_len := 0

	for _, dim := range dims {
		cur_dim := strings.Split(dim, "x")

		if len(cur_dim) != 3 {
			continue
		}

		len, _ := strconv.Atoi(cur_dim[0])
		wid, _ := strconv.Atoi(cur_dim[1])
		hgt, _ := strconv.Atoi(cur_dim[2])

		sorted := []int{len, wid, hgt}

		slices.Sort(sorted)

		bow_len := (sorted[0] * 2) + (sorted[1] * 2)
		total_vol := sorted[0] * sorted[1] * sorted[2]

		final_len += bow_len + total_vol
	}

	fmt.Printf("%d day %d result\n", year, 2)
	fmt.Println(final_len)
}
