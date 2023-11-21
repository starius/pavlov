package pavlov

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPureStrategies(t *testing.T) {
	cases := []struct {
		name        string
		game        PairGame
		wantScoreA  int
		wantScoreB  int
		wantHistory string
	}{
		{
			name: "always-cooperate vs always-defeat",
			game: PairGame{
				A:       StrategyAlwaysCooperate,
				B:       StrategyAlwaysDefeat,
				Payoffs: DefaultPayoffs,
				Rounds:  10,
			},
			wantScoreA:  -10,
			wantScoreB:  30,
			wantHistory: "C/D -> C/D -> C/D -> C/D -> C/D -> C/D -> C/D -> C/D -> C/D -> C/D",
		},
		{
			name: "always-defeat vs always-cooperate",
			game: PairGame{
				A:       StrategyAlwaysDefeat,
				B:       StrategyAlwaysCooperate,
				Payoffs: DefaultPayoffs,
				Rounds:  10,
			},
			wantScoreA:  30,
			wantScoreB:  -10,
			wantHistory: "D/C -> D/C -> D/C -> D/C -> D/C -> D/C -> D/C -> D/C -> D/C -> D/C",
		},
		{
			name: "always-defeat vs always-defeat",
			game: PairGame{
				A:       StrategyAlwaysDefeat,
				B:       StrategyAlwaysDefeat,
				Payoffs: DefaultPayoffs,
				Rounds:  10,
			},
			wantScoreA:  0,
			wantScoreB:  0,
			wantHistory: "D -> D -> D -> D -> D -> D -> D -> D -> D -> D",
		},
		{
			name: "always-cooperate vs always-cooperate",
			game: PairGame{
				A:       StrategyAlwaysCooperate,
				B:       StrategyAlwaysCooperate,
				Payoffs: DefaultPayoffs,
				Rounds:  10,
			},
			wantScoreA:  20,
			wantScoreB:  20,
			wantHistory: "C -> C -> C -> C -> C -> C -> C -> C -> C -> C",
		},

		{
			name: "TFT vs always-cooperate",
			game: PairGame{
				A:       StrategyTFT,
				B:       StrategyAlwaysCooperate,
				Payoffs: DefaultPayoffs,
				Rounds:  10,
			},
			wantScoreA:  20,
			wantScoreB:  20,
			wantHistory: "C -> C -> C -> C -> C -> C -> C -> C -> C -> C",
		},
		{
			name: "TFT vs always-defeat",
			game: PairGame{
				A:       StrategyTFT,
				B:       StrategyAlwaysDefeat,
				Payoffs: DefaultPayoffs,
				Rounds:  10,
			},
			wantScoreA:  -1,
			wantScoreB:  3,
			wantHistory: "C/D -> D -> D -> D -> D -> D -> D -> D -> D -> D",
		},
		{
			name: "TFT vs TFT",
			game: PairGame{
				A:       StrategyTFT,
				B:       StrategyTFT,
				Payoffs: DefaultPayoffs,
				Rounds:  10,
			},
			wantScoreA:  20,
			wantScoreB:  20,
			wantHistory: "C -> C -> C -> C -> C -> C -> C -> C -> C -> C",
		},

		{
			name: "Pavlov vs always-cooperate",
			game: PairGame{
				A:       StrategyPavlov,
				B:       StrategyAlwaysCooperate,
				Payoffs: DefaultPayoffs,
				Rounds:  10,
			},
			wantScoreA:  20,
			wantScoreB:  20,
			wantHistory: "C -> C -> C -> C -> C -> C -> C -> C -> C -> C",
		},
		{
			name: "Pavlov vs always-defeat",
			game: PairGame{
				A:       StrategyPavlov,
				B:       StrategyAlwaysDefeat,
				Payoffs: DefaultPayoffs,
				Rounds:  10,
			},
			wantScoreA:  -5,
			wantScoreB:  15,
			wantHistory: "C/D -> D -> C/D -> D -> C/D -> D -> C/D -> D -> C/D -> D",
		},
		{
			name: "Pavlov vs TFT",
			game: PairGame{
				A:       StrategyPavlov,
				B:       StrategyTFT,
				Payoffs: DefaultPayoffs,
				Rounds:  10,
			},
			wantScoreA:  20,
			wantScoreB:  20,
			wantHistory: "C -> C -> C -> C -> C -> C -> C -> C -> C -> C",
		},
		{
			name: "Pavlov vs Pavlov",
			game: PairGame{
				A:       StrategyPavlov,
				B:       StrategyPavlov,
				Payoffs: DefaultPayoffs,
				Rounds:  10,
			},
			wantScoreA:  20,
			wantScoreB:  20,
			wantHistory: "C -> C -> C -> C -> C -> C -> C -> C -> C -> C",
		},

		{
			name: "always-cooperate (mistake) vs always-defeat",
			game: PairGame{
				A:         StrategyAlwaysCooperate,
				B:         StrategyAlwaysDefeat,
				MistakesA: 1 << 5,
				Payoffs:   DefaultPayoffs,
				Rounds:    10,
			},
			wantScoreA:  -9,
			wantScoreB:  27,
			wantHistory: "C/D -> C/D -> C/D -> C/D -> C/D -> D -> C/D -> C/D -> C/D -> C/D",
		},
		{
			name: "always-defeat (mistake) vs always-cooperate",
			game: PairGame{
				A:         StrategyAlwaysDefeat,
				B:         StrategyAlwaysCooperate,
				MistakesA: 1 << 5,
				Payoffs:   DefaultPayoffs,
				Rounds:    10,
			},
			wantScoreA:  30,
			wantScoreB:  -10,
			wantHistory: "D/C -> D/C -> D/C -> D/C -> D/C -> D/C -> D/C -> D/C -> D/C -> D/C",
		},
		{
			name: "always-defeat (mistake) vs always-defeat",
			game: PairGame{
				A:         StrategyAlwaysDefeat,
				B:         StrategyAlwaysDefeat,
				MistakesA: 1 << 5,
				Payoffs:   DefaultPayoffs,
				Rounds:    10,
			},
			wantScoreA:  0,
			wantScoreB:  0,
			wantHistory: "D -> D -> D -> D -> D -> D -> D -> D -> D -> D",
		},
		{
			name: "always-cooperate (mistake) vs always-cooperate",
			game: PairGame{
				A:         StrategyAlwaysCooperate,
				B:         StrategyAlwaysCooperate,
				MistakesA: 1 << 5,
				Payoffs:   DefaultPayoffs,
				Rounds:    10,
			},
			wantScoreA:  21,
			wantScoreB:  17,
			wantHistory: "C -> C -> C -> C -> C -> D/C -> C -> C -> C -> C",
		},

		{
			name: "TFT (mistake) vs always-cooperate",
			game: PairGame{
				A:         StrategyTFT,
				B:         StrategyAlwaysCooperate,
				MistakesA: 1 << 5,
				Payoffs:   DefaultPayoffs,
				Rounds:    10,
			},
			wantScoreA:  21,
			wantScoreB:  17,
			wantHistory: "C -> C -> C -> C -> C -> D/C -> C -> C -> C -> C",
		},
		{
			name: "TFT vs always-cooperate (mistake)",
			game: PairGame{
				A:         StrategyTFT,
				B:         StrategyAlwaysCooperate,
				MistakesB: 1 << 5,
				Payoffs:   DefaultPayoffs,
				Rounds:    10,
			},
			wantScoreA:  18,
			wantScoreB:  18,
			wantHistory: "C -> C -> C -> C -> C -> C/D -> D/C -> C -> C -> C",
		},
		{
			name: "TFT (mistake) vs always-defeat",
			game: PairGame{
				A:         StrategyTFT,
				B:         StrategyAlwaysDefeat,
				MistakesA: 1 << 5,
				Payoffs:   DefaultPayoffs,
				Rounds:    10,
			},
			wantScoreA:  -1,
			wantScoreB:  3,
			wantHistory: "C/D -> D -> D -> D -> D -> D -> D -> D -> D -> D",
		},
		{
			name: "TFT vs always-defeat (mistake)",
			game: PairGame{
				A:         StrategyTFT,
				B:         StrategyAlwaysDefeat,
				MistakesB: 1 << 5,
				Payoffs:   DefaultPayoffs,
				Rounds:    10,
			},
			wantScoreA:  -1,
			wantScoreB:  3,
			wantHistory: "C/D -> D -> D -> D -> D -> D -> D -> D -> D -> D",
		},
		{
			name: "TFT (mistake) vs TFT",
			game: PairGame{
				A:         StrategyTFT,
				B:         StrategyTFT,
				MistakesA: 1 << 5,
				Payoffs:   DefaultPayoffs,
				Rounds:    10,
			},
			wantScoreA:  17,
			wantScoreB:  13,
			wantHistory: "C -> C -> C -> C -> C -> D/C -> C/D -> D/C -> C/D -> D/C",
		},

		{
			name: "Pavlov (mistake) vs always-cooperate",
			game: PairGame{
				A:         StrategyPavlov,
				B:         StrategyAlwaysCooperate,
				MistakesA: 1 << 5,
				Payoffs:   DefaultPayoffs,
				Rounds:    10,
			},
			wantScoreA:  25,
			wantScoreB:  5,
			wantHistory: "C -> C -> C -> C -> C -> D/C -> D/C -> D/C -> D/C -> D/C",
		},
		{
			name: "Pavlov vs always-cooperate (mistake)",
			game: PairGame{
				A:         StrategyPavlov,
				B:         StrategyAlwaysCooperate,
				MistakesB: 1 << 5,
				Payoffs:   DefaultPayoffs,
				Rounds:    10,
			},
			wantScoreA:  21,
			wantScoreB:  9,
			wantHistory: "C -> C -> C -> C -> C -> C/D -> D/C -> D/C -> D/C -> D/C",
		},
		{
			name: "Pavlov (mistake) vs always-defeat",
			game: PairGame{
				A:         StrategyPavlov,
				B:         StrategyAlwaysDefeat,
				MistakesA: 1 << 6,
				Payoffs:   DefaultPayoffs,
				Rounds:    10,
			},
			wantScoreA:  -5,
			wantScoreB:  15,
			wantHistory: "C/D -> D -> C/D -> D -> C/D -> D -> D -> C/D -> D -> C/D",
		},
		{
			name: "Pavlov vs always-defeat (mistake)",
			game: PairGame{
				A:         StrategyPavlov,
				B:         StrategyAlwaysDefeat,
				MistakesB: 1 << 5,
				Payoffs:   DefaultPayoffs,
				Rounds:    10,
			},
			wantScoreA:  -5,
			wantScoreB:  15,
			wantHistory: "C/D -> D -> C/D -> D -> C/D -> D -> C/D -> D -> C/D -> D",
		},
		{
			name: "Pavlov (mistake) vs TFT",
			game: PairGame{
				A:         StrategyPavlov,
				B:         StrategyTFT,
				MistakesA: 1 << 5,
				Payoffs:   DefaultPayoffs,
				Rounds:    10,
			},
			wantScoreA:  15,
			wantScoreB:  11,
			wantHistory: "C -> C -> C -> C -> C -> D/C -> D -> C/D -> D/C -> D",
		},
		{
			name: "Pavlov vs TFT (mistake)",
			game: PairGame{
				A:         StrategyPavlov,
				B:         StrategyTFT,
				MistakesB: 1 << 5,
				Payoffs:   DefaultPayoffs,
				Rounds:    10,
			},
			wantScoreA:  14,
			wantScoreB:  14,
			wantHistory: "C -> C -> C -> C -> C -> C/D -> D/C -> D -> C/D -> D/C",
		},
		{
			name: "Pavlov (mistake) vs Pavlov",
			game: PairGame{
				A:         StrategyPavlov,
				B:         StrategyPavlov,
				MistakesA: 1 << 5,
				Payoffs:   DefaultPayoffs,
				Rounds:    10,
			},
			wantScoreA:  19,
			wantScoreB:  15,
			wantHistory: "C -> C -> C -> C -> C -> D/C -> D -> C -> C -> C",
		},

		{
			name: "TFTWF vs TFTWF",
			game: PairGame{
				A:       StrategyTFTWF,
				B:       StrategyTFTWF,
				Payoffs: DefaultPayoffs,
				Rounds:  10,
			},
			wantScoreA:  20,
			wantScoreB:  20,
			wantHistory: "C -> C -> C -> C -> C -> C -> C -> C -> C -> C",
		},
		{
			name: "TFTWF vs always-cooperate",
			game: PairGame{
				A:       StrategyTFTWF,
				B:       StrategyAlwaysCooperate,
				Payoffs: DefaultPayoffs,
				Rounds:  10,
			},
			wantScoreA:  20,
			wantScoreB:  20,
			wantHistory: "C -> C -> C -> C -> C -> C -> C -> C -> C -> C",
		},
		{
			name: "TFTWF vs always-defeat",
			game: PairGame{
				A:       StrategyTFTWF,
				B:       StrategyAlwaysDefeat,
				Payoffs: DefaultPayoffs,
				Rounds:  10,
			},
			wantScoreA:  -2,
			wantScoreB:  6,
			wantHistory: "C/D -> C/D -> D -> D -> D -> D -> D -> D -> D -> D",
		},
		{
			name: "TFTWF vs TFT",
			game: PairGame{
				A:       StrategyTFTWF,
				B:       StrategyTFT,
				Payoffs: DefaultPayoffs,
				Rounds:  10,
			},
			wantScoreA:  20,
			wantScoreB:  20,
			wantHistory: "C -> C -> C -> C -> C -> C -> C -> C -> C -> C",
		},
		{
			name: "TFTWF vs Pavlov",
			game: PairGame{
				A:       StrategyTFTWF,
				B:       StrategyPavlov,
				Payoffs: DefaultPayoffs,
				Rounds:  10,
			},
			wantScoreA:  20,
			wantScoreB:  20,
			wantHistory: "C -> C -> C -> C -> C -> C -> C -> C -> C -> C",
		},
		{
			name: "TFTWF (mistake) vs TFTWF",
			game: PairGame{
				A:         StrategyTFTWF,
				B:         StrategyTFTWF,
				MistakesA: 1 << 5,
				Payoffs:   DefaultPayoffs,
				Rounds:    10,
			},
			wantScoreA:  21,
			wantScoreB:  17,
			wantHistory: "C -> C -> C -> C -> C -> D/C -> C -> C -> C -> C",
		},
		{
			name: "TFTWF (mistake) vs always-cooperate",
			game: PairGame{
				A:         StrategyTFTWF,
				B:         StrategyAlwaysCooperate,
				MistakesA: 1 << 5,
				Payoffs:   DefaultPayoffs,
				Rounds:    10,
			},
			wantScoreA:  21,
			wantScoreB:  17,
			wantHistory: "C -> C -> C -> C -> C -> D/C -> C -> C -> C -> C",
		},
		{
			name: "TFTWF vs always-cooperate (mistake)",
			game: PairGame{
				A:         StrategyTFTWF,
				B:         StrategyAlwaysCooperate,
				MistakesB: 1 << 5,
				Payoffs:   DefaultPayoffs,
				Rounds:    10,
			},
			wantScoreA:  17,
			wantScoreB:  21,
			wantHistory: "C -> C -> C -> C -> C -> C/D -> C -> C -> C -> C",
		},
		{
			name: "TFTWF (mistake) vs always-defeat",
			game: PairGame{
				A:         StrategyTFTWF,
				B:         StrategyAlwaysDefeat,
				MistakesA: 1 << 5,
				Payoffs:   DefaultPayoffs,
				Rounds:    10,
			},
			wantScoreA:  -2,
			wantScoreB:  6,
			wantHistory: "C/D -> C/D -> D -> D -> D -> D -> D -> D -> D -> D",
		},
		{
			name: "TFTWF vs always-defeat (mistake)",
			game: PairGame{
				A:         StrategyTFTWF,
				B:         StrategyAlwaysDefeat,
				MistakesB: 1 << 5,
				Payoffs:   DefaultPayoffs,
				Rounds:    10,
			},
			wantScoreA:  -2,
			wantScoreB:  6,
			wantHistory: "C/D -> C/D -> D -> D -> D -> D -> D -> D -> D -> D",
		},
		{
			name: "TFTWF (mistake) vs TFT",
			game: PairGame{
				A:         StrategyTFTWF,
				B:         StrategyTFT,
				MistakesA: 1 << 5,
				Payoffs:   DefaultPayoffs,
				Rounds:    10,
			},
			wantScoreA:  18,
			wantScoreB:  18,
			wantHistory: "C -> C -> C -> C -> C -> D/C -> C/D -> C -> C -> C",
		},
		{
			name: "TFTWF vs TFT (mistake)",
			game: PairGame{
				A:         StrategyTFTWF,
				B:         StrategyTFT,
				MistakesB: 1 << 5,
				Payoffs:   DefaultPayoffs,
				Rounds:    10,
			},
			wantScoreA:  17,
			wantScoreB:  21,
			wantHistory: "C -> C -> C -> C -> C -> C/D -> C -> C -> C -> C",
		},
		{
			name: "TFTWF (mistake) vs Pavlov",
			game: PairGame{
				A:         StrategyTFTWF,
				B:         StrategyPavlov,
				MistakesA: 1 << 5,
				Payoffs:   DefaultPayoffs,
				Rounds:    10,
			},
			wantScoreA:  14,
			wantScoreB:  14,
			wantHistory: "C -> C -> C -> C -> C -> D/C -> C/D -> C/D -> D -> D/C",
		},
		{
			name: "TFTWF vs Pavlov (mistake)",
			game: PairGame{
				A:         StrategyTFTWF,
				B:         StrategyPavlov,
				MistakesB: 1 << 5,
				Payoffs:   DefaultPayoffs,
				Rounds:    10,
			},
			wantScoreA:  10,
			wantScoreB:  18,
			wantHistory: "C -> C -> C -> C -> C -> C/D -> C/D -> D -> D/C -> C/D",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			gotScoreA, gotScoreB, history := SimulateMatch(tc.game)
			historyStr := HistoryToString(history, int(tc.game.Rounds))
			assert.Equal(t, tc.wantScoreA, gotScoreA, tc.game)
			assert.Equal(t, tc.wantScoreB, gotScoreB, tc.game)
			assert.Equal(t, tc.wantHistory, historyStr, tc.game)
		})
	}
}
