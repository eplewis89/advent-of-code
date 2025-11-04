package fifteen

import (
	"github.com/eplewis89/advent-of-code/bootstrap"
)

func Solve() {
	print("===== AOC 2015 =====")
	day_one()
}

func day_one() {
	input, err := bootstrap.ReadInput(2015, 1)

	if err != nil {
		return
	}

	print(input)
}
