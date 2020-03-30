package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"roller-go/dice"
	"time"
)

var Usage = func() {
	fmt.Fprintf(flag.CommandLine.Output(), "")
	flag.PrintDefaults()
}

func main() {

	advantage := flag.Bool("advantage", false, "Roll with advantage")
	disadvantage := flag.Bool("disadvantage", false, "Roll with disadvantage")

	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	inputStr := os.Args[len(os.Args)-1]
	dice, err := dice.Parse(inputStr)
	panicIf(err)

	fmt.Print("Rolled: ")
	nums := dice.Roll()
	fmt.Println(nums)

	result := 0
	if *advantage {
		adv := maxIn(dice.Roll())
		orig := maxIn(nums)
		result = max(orig, adv)
	} else if *disadvantage {
		disadv := minIn(dice.Roll())
		orig := minIn(nums)
		result = min(orig, disadv)
	} else {
		result = sum(nums)
	}

	if dice.ModSign != "" {
		fmt.Println("Mod:", dice.ModSign, dice.Mod)
		result = dice.ModFn(result)
	}

	fmt.Println("Result:", result)
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func sum(nums []int) int {
	result := 0
	for _, n := range nums {
		result += n
	}
	return result
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func minIn(nums []int) int {
	result := nums[0]
	for _, n := range nums {
		if n <= result {
			result = n
		}
	}
	return result
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func maxIn(nums []int) int {
	result := 0
	for _, n := range nums {
		if n >= result {
			result = n
		}
	}
	return result
}
