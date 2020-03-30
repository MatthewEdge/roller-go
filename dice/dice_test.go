package dice

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		diceStr   string
		result    *DiceString
		shouldErr bool
	}{
		{
			"2d6",
			&DiceString{
				Sides:   6,
				NumDice: 2,
				Mod:     0,
				ModFn:   func(a int) int { return a },
			},
			false,
		},
		{
			"d6",
			&DiceString{
				Sides:   6,
				NumDice: 1,
				Mod:     0,
				ModFn:   func(a int) int { return a },
			},
			false,
		},
		{
			"2d20+49",
			&DiceString{
				Sides:   20,
				NumDice: 2,
				Mod:     49,
				ModFn:   func(a int) int { return a },
			},
			false,
		},
		{
			"woops",
			&DiceString{},
			true,
		},
	}

	for _, scenario := range tests {
		t.Run(scenario.diceStr, func(t *testing.T) {
			res, err := Parse(scenario.diceStr)
			if err != nil {
				fmt.Println(err.Error())
			}

			if !scenario.shouldErr && err != nil {
				t.Error(scenario.diceStr, "should have succeeded. Got error: ", err.Error())
				return
			}

			if !scenario.shouldErr && res == nil {
				t.Error(scenario.diceStr, "should have succeeded. Got nil result!")
				return
			}

			if scenario.shouldErr && err == nil {
				t.Error(scenario.diceStr, "should have returned an error")
				return
			}

			if !scenario.shouldErr && !diceEqual(res, scenario.result) {
				t.Error(scenario.diceStr, "should have returned", scenario.result, "but got", res)
				return
			}
		})
	}
}

func diceEqual(a, b *DiceString) bool {
	if a == nil {
		fmt.Println("diceEqual: a was nil!")
		return false
	}

	if b == nil {
		fmt.Println("diceEqual: b was nil!")
		return false
	}

	return a.NumDice == b.NumDice &&
		a.Sides == b.Sides &&
		a.Mod == b.Mod
}
