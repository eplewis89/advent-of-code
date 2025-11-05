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
	day_three_part_one()
	day_three_part_two()
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

func day_three_part_one() {
	input, err := bootstrap.ReadInput(year, 3)

	if err != nil {
		log.Printf("failed to get data: %v", err)

		return
	}

	curloc := []int{0, 0}
	storage := make(map[int]map[int]bool)
	storage[curloc[0]] = map[int]bool{
		curloc[1]: true,
	}

	for _, x := range input {
		switch x {
		case '^':
			curloc[0] += 1
		case 'v':
			curloc[0] -= 1
		case '<':
			curloc[1] += 1
		case '>':
			curloc[1] -= 1
		}

		if storage[curloc[0]] == nil {
			storage[curloc[0]] = make(map[int]bool)
		}

		storage[curloc[0]][curloc[1]] = true
	}

	fmt.Printf("%v\n", curloc)

	total := 0

	for _, v := range storage {
		total += len(v)
	}

	fmt.Printf("%d day %d part 1 result\n", year, 3)
	fmt.Println(total)
}

func day_three_part_two() {
	input, err := bootstrap.ReadInput(year, 3)

	if err != nil {
		log.Printf("failed to get data: %v", err)

		return
	}

	cursantaloc := []int{0, 0}
	currobotloc := []int{0, 0}

	storage := make(map[int]map[int]bool)
	storage[cursantaloc[0]] = map[int]bool{
		cursantaloc[1]: true,
	}

	robosanta := false

	for _, x := range input {
		updown := 0
		leftright := 0

		switch x {
		case '^':
			updown = 1
		case 'v':
			updown = -1
		case '<':
			leftright = 1
		case '>':
			leftright = -1
		}

		if robosanta {
			currobotloc[0] += updown
			currobotloc[1] += leftright

			updown = currobotloc[0]
			leftright = currobotloc[1]
		} else {
			cursantaloc[0] += updown
			cursantaloc[1] += leftright

			updown = cursantaloc[0]
			leftright = cursantaloc[1]
		}

		robosanta = !robosanta

		if storage[updown] == nil {
			storage[updown] = make(map[int]bool)
		}

		storage[updown][leftright] = true
	}

	total := 0

	for _, v := range storage {
		total += len(v)
	}

	fmt.Printf("%d day %d part 2 result\n", year, 3)
	fmt.Println(total)
}
