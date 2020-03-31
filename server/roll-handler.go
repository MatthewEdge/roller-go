package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"roller-go/dice"
	diceMath "roller-go/math"
)

func (s *server) handleRoll() http.HandlerFunc {

	type request struct {
		DiceString   string `json:"diceStr"`
		Advantage    bool   `json:"advantage"`
		Disadvantage bool   `json:"disadvantage"`
	}

	type response struct {
		Rolls  []int `json:"rolls"`
		Mod    int   `json:"mod"`
		Result int   `json:"result"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req request

		err := json.NewDecoder(r.Body).Decode(&req)
		defer r.Body.Close()
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

		resp := &response{}
		resp.Rolls = dice.Roll()
		resp.Result = 0

		if req.Advantage {
			adv := diceMath.MaxIn(dice.Roll())
			orig := diceMath.MaxIn(resp.Rolls)
			resp.Result = diceMath.Max(orig, adv)
		} else if req.Disadvantage {
			disadv := diceMath.MinIn(dice.Roll())
			orig := diceMath.MinIn(resp.Rolls)
			resp.Result = diceMath.Min(orig, disadv)
		} else {
			resp.Result = diceMath.Sum(resp.Rolls)
		}

		if dice.ModSign != "" {
			fmt.Println("Mod:", dice.ModSign, dice.Mod)
			resp.Result = dice.ModFn(resp.Result)
		}

		err = s.respond(w, resp)
		if err != nil {
			w.WriteHeader(500)
			fmt.Println("Failed to serialize response", err)
			return
		}
	}
}
