package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	advantage := flag.Bool("advantage", false, "Roll with advantage")
	disadvantage := flag.Bool("disadvantage", false, "Roll with disadvantage")

	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	diceStr := os.Args[len(os.Args)-1]
	diceR := regexp.MustCompile("(\\d)d(\\d)([+|-])?(\\d)?")
	if !diceR.MatchString(diceStr) {
		fmt.Printf("%s is not a valid dice string\n", diceStr)
		os.Exit(1)
	}
	parsed := diceR.FindAllStringSubmatch(diceStr, -1)

	numDice, err := strconv.Atoi(parsed[0][1])
	panicIf(err)

	sides, err := strconv.Atoi(parsed[0][2])
	panicIf(err)

	nums := make([]int, numDice)
	fmt.Print("Rolled: ")
	for n := 0; n < numDice; n++ {
		i := randInt(sides)
		fmt.Printf("%d ", i)
		nums = append(nums, i)
	}
	fmt.Println()

	result := 0
	if *advantage {
		result = max(nums)
	} else if *disadvantage {
		result = min(nums)
	} else {
		result = sum(nums)
	}

	mod := parsed[0][3]
	modVal := parsed[0][4]
	if mod != "" && modVal != "" {
		modValue, err := strconv.Atoi(modVal)
		panicIf(err)

		if mod == "-" {
			result -= modValue
		} else {
			result += modValue
		}

		fmt.Printf("Mod: %s%d\n", mod, modValue)
	}

	fmt.Printf("Result: %d\n", result)
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func randInt(max int) int {
	return (int)(1 + rand.Intn(max))
}

func sum(nums []int) int {
	result := 0
	for n := range nums {
		result += nums[n]
	}
	return result
}

func min(nums []int) int {
	result := 0
	for n := range nums {
		if nums[n] <= result {
			result = nums[n]
		}
	}
	return result
}

func max(nums []int) int {
	result := 0
	for n := range nums {
		if nums[n] >= result {
			result = nums[n]
		}
	}
	return result
}
