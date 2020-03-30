package dice

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
)

var diceR *regexp.Regexp = regexp.MustCompile("(\\d+)?d(\\d+)([+|-])?(\\d+)?")

type DiceString struct {
	orig    string
	NumDice int
	Sides   int
	Mod     int
	ModSign string
	ModFn   func(int) int
}

func newDiceString(diceStr string) *DiceString {
	return &DiceString{
		orig:    diceStr,
		NumDice: 0,
		Sides:   0,
		ModFn:   func(total int) int { return total }, // no-op
	}
}

// Parse takes a diceStr of the form captured by diceR. For example:
// 3d6-6 or d6+4
// The leading number can be omitted. Thus: 1d6 and d6 are equivalent
func Parse(diceStr string) (*DiceString, error) {
	result := newDiceString(diceStr)

	if !diceR.MatchString(diceStr) {
		err := fmt.Errorf("%s is not a valid dice string", diceStr)
		return nil, err
	}
	parsed := diceR.FindAllStringSubmatch(diceStr, -1)

	numDice, err := strconv.Atoi(parsed[0][1])
	if err != nil {
		numDice = 1
	}
	result.NumDice = numDice

	sides, err := strconv.Atoi(parsed[0][2])
	if err != nil {
		return nil, err
	}
	result.Sides = sides

	mod := parsed[0][3]
	result.ModSign = mod

	modValue, err := strconv.Atoi(parsed[0][4])
	if mod != "" {
		if err != nil {
			return nil, err
		}

		result.Mod = modValue

		if mod == "+" {
			result.ModFn = result.addMod
		} else {
			result.ModFn = result.subMod
		}
	}

	return result, nil
}

// Roll a number of dice with the given number of dice sides. Results of a roll are [1, sides]
func (d *DiceString) Roll() []int {
	nums := make([]int, d.NumDice)
	for n := 0; n < d.NumDice; n++ {
		i := randInt(d.Sides)
		nums[n] = i
	}
	return nums
}

func randInt(Max int) int {
	return (int)(1 + rand.Intn(Max))
}

func (d *DiceString) addMod(total int) int {
	return total + d.Mod
}

func (d *DiceString) subMod(total int) int {
	return total - d.Mod
}
