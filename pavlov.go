package pavlov

import (
	"fmt"
	"strings"
)

// PureStrategy defines a pure strategy (not using random).
//
// A strategy has memory for up to 3 rounds. The result of the first round
// is determined by Move1 (1 = cooperate, 0 = defeat).
// The result of the second round is the i-th right bit of Move2,
// where i = prev_round_my (concatenated with) prev_round_they:
//
// value of i | they defeat | they coop
// -----------+-------------+----------
// I defeat   | 00 (0)      | 01 (1)
// I coop     | 10 (2)      | 11 (3)
//
// Let's call i the outcome of a round.
// For the third round, Move3 is used. The outcome of the first round
// is concatenated with the outcome of the second round (i_1 << 2 | i_2).
//
// Subsequent rounds use MoveN. Outcomes of 3 previous rounds are concatenated
// into a 6-bit value, which serves of bit index (counting from the right side):
// (outcome n-3) << 4 | (outcome n-2) << 2 | (outcome n).
type PureStrategy struct {
	Move1 uint8
	Move2 uint8
	Move3 uint16
	MoveN uint64
}

func (s PureStrategy) String() string {
	return fmt.Sprintf("PureStrategy(move1=%01b move2=%04b move3=%016b move2=%064b)", s.Move1, s.Move2, s.Move3, s.MoveN)
}

func NewPureStrategyWithoutMemory(move uint8) PureStrategy {
	return NewPureStrategyWithMemory1(move, (move<<3)|(move<<2)|(move<<1)|move)
}

func NewPureStrategyWithMemory1(move1 uint8, move2 uint8) PureStrategy {
	m2 := uint16(move2)
	return NewPureStrategyWithMemory2(move1, move2, (m2<<12)|(m2<<8)|(m2<<4)|m2)
}

func NewPureStrategyWithMemory2(move1 uint8, move2 uint8, move3 uint16) PureStrategy {
	m3 := uint64(move3)
	return NewPureStrategyWithMemory3(move1, move2, move3, (m3<<48)|(m3<<32)|(m3<<16)|m3)
}

func NewPureStrategyWithMemory3(move1 uint8, move2 uint8, move3 uint16, moven uint64) PureStrategy {
	return PureStrategy{
		Move1: move1,
		Move2: move2,
		Move3: move3,
		MoveN: moven,
	}
}

type PairGame struct {
	A, B PureStrategy

	// If i-th right bit is 1, i-th outcome is defeat from the corresponding side.
	MistakesA, MistakesB uint64

	// Payoffs specifies payoffs of the game.
	// Index in the array is the outcome (see PureStrategy for outcome definition).
	Payoffs [4]int8

	Rounds uint8 // Min 1, max 32.
}

var DefaultPayoffs = [4]int8{
	0,  // DD.
	3,  // DC.
	-1, // CD.
	2,  // CC.
}

const (
	MovesTFT    = 0b1010
	MovesPavlov = 0b1001
)

var (
	StrategyAlwaysCooperate = NewPureStrategyWithoutMemory(1)
	StrategyAlwaysDefeat    = NewPureStrategyWithoutMemory(0)
	StrategyTFT             = NewPureStrategyWithMemory1(1, MovesTFT)
	StrategyTFTWF           = NewPureStrategyWithMemory2(1, 0b1111, 0b1111101011111010)
	StrategyPavlov          = NewPureStrategyWithMemory1(1, MovesPavlov)
)

func SimulateMatch(g PairGame) (scoreA, scoreB int, history uint64) {
	nonMistakesA := ^g.MistakesA
	nonMistakesB := ^g.MistakesB

	moveA1 := g.A.Move1 & uint8(nonMistakesA&1)
	moveB1 := g.B.Move1 & uint8(nonMistakesB&1)
	outcomeA1 := (moveA1<<1 | moveB1)
	outcomeB1 := (moveB1<<1 | moveA1)
	scoreA += int(g.Payoffs[outcomeA1])
	scoreB += int(g.Payoffs[outcomeB1])

	history = uint64(outcomeA1)

	if g.Rounds == 1 {
		return scoreA, scoreB, history
	}

	nonMistakesA >>= 1
	nonMistakesB >>= 1

	moveA2 := (g.A.Move2 >> outcomeA1) & uint8(nonMistakesA&1)
	moveB2 := (g.B.Move2 >> outcomeB1) & uint8(nonMistakesB&1)
	outcomeA2 := (moveA2<<1 | moveB2)
	outcomeB2 := (moveB2<<1 | moveA2)
	scoreA += int(g.Payoffs[outcomeA2])
	scoreB += int(g.Payoffs[outcomeB2])

	history = (history << 2) | uint64(outcomeA2)

	if g.Rounds == 2 {
		return scoreA, scoreB, history
	}

	nonMistakesA >>= 1
	nonMistakesB >>= 1

	outcomeAhist := (outcomeA1 << 2) | outcomeA2
	outcomeBhist := (outcomeB1 << 2) | outcomeB2
	moveA3 := uint8((g.A.Move3>>outcomeAhist)&1) & uint8(nonMistakesA&1)
	moveB3 := uint8((g.B.Move3>>outcomeBhist)&1) & uint8(nonMistakesB&1)
	outcomeAlast := (moveA3<<1 | moveB3)
	outcomeBlast := (moveB3<<1 | moveA3)
	scoreA += int(g.Payoffs[outcomeAlast])
	scoreB += int(g.Payoffs[outcomeBlast])

	history = (history << 2) | uint64(outcomeAlast)

	for r := uint8(4); r <= g.Rounds; r++ {
		nonMistakesA >>= 1
		nonMistakesB >>= 1

		outcomeAhist = ((outcomeAhist << 2) | outcomeAlast) & (1<<6 - 1)
		outcomeBhist = ((outcomeBhist << 2) | outcomeBlast) & (1<<6 - 1)
		moveA := uint8((g.A.MoveN>>outcomeAhist)&1) & uint8(nonMistakesA&1)
		moveB := uint8((g.B.MoveN>>outcomeBhist)&1) & uint8(nonMistakesB&1)
		outcomeAlast = (moveA<<1 | moveB)
		outcomeBlast = (moveB<<1 | moveA)
		scoreA += int(g.Payoffs[outcomeAlast])
		scoreB += int(g.Payoffs[outcomeBlast])

		history = (history << 2) | uint64(outcomeAlast)
	}

	return scoreA, scoreB, history
}

var outcomesStr = [4]string{
	"D",
	"D/C",
	"C/D",
	"C",
}

func HistoryToString(history uint64, rounds int) string {
	results := make([]string, 0, rounds)
	for r := 1; r <= rounds; r++ {
		outcome := (history >> (2 * (rounds - r))) & (1<<2 - 1)
		results = append(results, outcomesStr[outcome])
	}
	return strings.Join(results, " -> ")
}
