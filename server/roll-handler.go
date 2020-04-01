package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"roller-go/dice"
)

func (s *server) handleRoll() http.HandlerFunc {

	type request struct {
		DiceString   string `json:"diceStr"`
		Advantage    bool   `json:"advantage"`
		Disadvantage bool   `json:"disadvantage"`
	}

	type response struct {
		Rolls  []int  `json:"rolls"`
		Mod    string `json:"mod"`
		Result int    `json:"result"`
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

		var roll dice.RollResponse

		if req.Advantage {
			roll, err = dice.RollAdvantage(req.DiceString)
		} else if req.Disadvantage {
			roll, err = dice.RollDisadvantage(req.DiceString)
		} else {
			roll, err = dice.Roll(req.DiceString)
		}

		if err != nil {
			w.WriteHeader(400)
			fmt.Println("Failed to roll dice with input:", req.DiceString)
			return
		}

		resp := &response{
			Rolls:  roll.Rolls,
			Mod:    roll.Mod,
			Result: roll.Total,
		}
		err = s.respond(w, resp)
		if err != nil {
			w.WriteHeader(500)
			fmt.Println("Failed to serialize response", err)
			return
		}
	}
}
