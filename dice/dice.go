package dice

import (
	"fmt"
	"math/rand"
	"regexp"
	diceMath "roller-go/math"
	"strconv"
)

var diceR *regexp.Regexp = regexp.MustCompile("(\\d+)?d(\\d+)([+|-])?(\\d+)?")

type RollResponse struct {
	Rolls []int
	Total int
	Mod   string
}

type InvalidDiceStr struct {
	diceStr string
}

func (e *InvalidDiceStr) Error() string {
	return fmt.Sprintf("%s is an invalid dice string", e.diceStr)
}

// Roll the given diceStr
func Roll(diceStr string) (RollResponse, error) {
	resp := RollResponse{}

	dice, err := parse(diceStr)
	if err != nil {
		return resp, &InvalidDiceStr{diceStr: diceStr}
	}

	resp.Rolls = dice.roll()
	resp.Total = diceMath.MaxIn(resp.Rolls)

	if dice.ModSign != "" {
		resp.Mod = fmt.Sprint(dice.ModSign, dice.Mod)
		resp.Total = dice.ModFn(resp.Total)
	}

	return resp, nil

}

// RollAdvantage rolls the given diceStr twice, taking the higher of the two rolls
func RollAdvantage(diceStr string) (RollResponse, error) {
	resp, err := Roll(diceStr)
	if err != nil {
		return resp, err
	}

	second, err := Roll(diceStr)
	resp.Rolls = append(resp.Rolls, second.Rolls...)
	resp.Total = diceMath.Max(resp.Total, second.Total)

	return resp, nil
}

// RollDisadvantage rolls the given diceStr twice, taking the lower of the two rolls
func RollDisadvantage(diceStr string) (RollResponse, error) {
	resp, err := Roll(diceStr)
	if err != nil {
		return resp, err
	}

	second, err := Roll(diceStr)
	resp.Rolls = append(resp.Rolls, second.Rolls...)
	resp.Total = diceMath.Min(resp.Total, second.Total)

	return resp, nil
}

type rollReq struct {
	orig    string
	NumDice int
	Sides   int
	Mod     int
	ModSign string
	ModFn   func(int) int
}

// parse takes a diceStr of the form captured by diceR. For example:
// 3d6-6 or d6+4
// The leading number can be omitted. Thus: 1d6 and d6 are equivalent
func parse(diceStr string) (*rollReq, error) {
	result := &rollReq{
		orig:    diceStr,
		NumDice: 0,
		Sides:   0,
		ModFn:   func(total int) int { return total }, // no-op
	}

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
func (d *rollReq) roll() []int {
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

func (d *rollReq) addMod(total int) int {
	return total + d.Mod
}

func (d *rollReq) subMod(total int) int {
	return total - d.Mod
}
