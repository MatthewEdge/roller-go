package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"roller-go/dice"
	"time"
)

type rollInput struct {
	DiceString   string `json:"diceStr"`
	Advantage    bool   `json:"advantage"`
	Disadvantage bool   `json:"disadvantage"`
}

type rollOutput struct {
	Rolls  []int `json:"rolls"`
	Mod    int   `json:"mod"`
	Result int   `json:"result"`
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}

func run() error {

	rand.Seed(time.Now().UnixNano())

	router := http.NewServeMux()
	router.HandleFunc("/roll", func(w http.ResponseWriter, r *http.Request) {
		var req rollInput

		err := json.NewDecoder(r.Body).Decode(&req)
		defer r.Body.Close() // TODO is this necessary?
		if err != nil {
			w.WriteHeader(422)
			fmt.Println("Failed to decode body to RollInput")
			return
		}

		dice, err := dice.Parse(req.DiceString)
		if err != nil {
			w.WriteHeader(400)
			fmt.Println("Failed to roll dice with input:", req.DiceString)
			return
		}

		resp := &rollOutput{}
		resp.Rolls = dice.Roll()
		resp.Result = 0

		if req.Advantage {
			adv := maxIn(dice.Roll())
			orig := maxIn(resp.Rolls)
			resp.Result = max(orig, adv)
		} else if req.Disadvantage {
			disadv := minIn(dice.Roll())
			orig := minIn(resp.Rolls)
			resp.Result = min(orig, disadv)
		} else {
			resp.Result = sum(resp.Rolls)
		}

		if dice.ModSign != "" {
			fmt.Println("Mod:", dice.ModSign, dice.Mod)
			resp.Result = dice.ModFn(resp.Result)
		}

		fmt.Println(resp)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			w.WriteHeader(500)
			fmt.Println("Failed to serialize response", err)
			return
		}
	})

	if err := http.ListenAndServe(":8080", router); err != nil {
		return err
	}

	return nil
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
