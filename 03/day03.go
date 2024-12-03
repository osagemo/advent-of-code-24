package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/osagemo/advent-of-code-24/internal/utils"
)

//go:embed input.txt
var input string

type State string

const (
	LookForStart          State = "LookForStart"
	EvaluatingInstruction State = "EvaluatingInstruction"
	EvaluatingArguments   State = "EvaluatingArguments"
)

type TransitionFunc func(rune, *Context) (State, error)

type Transition struct {
	Condition func(rune) bool
	Action    TransitionFunc
}

type StateMachine struct {
	State           State
	TransitionTable map[State][]Transition
}

type Operation string

const (
	Mul Operation = "mul"
)

type Instruction struct {
	X         int
	Y         int
	Operation Operation
}

type Context struct {
	Instructions         []Instruction
	CandidateInstruction Instruction
	Buffer               strings.Builder
}

func (fsm *StateMachine) HandleNext(c rune, ctx *Context) (State, error) {
	transitions := fsm.TransitionTable[fsm.State]
	for _, t := range transitions {
		if t.Condition(c) {
			return t.Action(c, ctx)
		}
	}
	return LookForStart, fmt.Errorf("invalid input '%c' in state %s", c, fsm.State)
}

func ParseInstructions(input string) []Instruction {
	ctx := &Context{Instructions: []Instruction{}, CandidateInstruction: Instruction{}}
	fsm := &StateMachine{
		State: LookForStart,
		TransitionTable: map[State][]Transition{
			LookForStart: {
				{
					Condition: func(c rune) bool { return unicode.IsLetter(c) },
					Action: func(c rune, ctx *Context) (State, error) {
						ctx.Buffer.WriteRune(c)
						return EvaluatingInstruction, nil
					},
				},
			},
			EvaluatingInstruction: {
				{
					Condition: unicode.IsLetter,
					Action: func(c rune, ctx *Context) (State, error) {
						ctx.Buffer.WriteRune(c)
						if ctx.Buffer.String() == string(Mul) {
							ctx.CandidateInstruction.Operation = Mul
							ctx.Buffer.Reset()
							return EvaluatingArguments, nil
						}
						return EvaluatingInstruction, nil
					},
				},
			},
			EvaluatingArguments: {
				{
					Condition: func(c rune) bool { return unicode.IsDigit(c) },
					Action: func(c rune, ctx *Context) (State, error) {
						ctx.Buffer.WriteRune(c)
						return EvaluatingArguments, nil
					},
				},
				{
					Condition: func(c rune) bool { return c == ',' },
					Action: func(c rune, ctx *Context) (State, error) {
						ctx.CandidateInstruction.X = utils.MustParseInt(ctx.Buffer.String())
						ctx.Buffer.Reset()
						return EvaluatingArguments, nil
					},
				},
				{
					Condition: func(c rune) bool { return c == ')' },
					Action: func(c rune, ctx *Context) (State, error) {
						ctx.CandidateInstruction.Y = utils.MustParseInt(ctx.Buffer.String())
						ctx.Instructions = append(ctx.Instructions, ctx.CandidateInstruction)
						ctx.CandidateInstruction = Instruction{}
						ctx.Buffer.Reset()
						return LookForStart, nil
					},
				},
			},
		},
	}

	for _, c := range input {
		nextState, err := fsm.HandleNext(c, ctx)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		fsm.State = nextState
	}

	return ctx.Instructions
}

func Part1(input string) int {
	instructions := ParseInstructions(input)
	sum := 0
	for _, i := range instructions {
		sum += ExecuteInstruction(i)
	}

	return sum
}

func Part2(input string) int {
	return 0
}

func ExecuteInstruction(i Instruction) int {
	return i.X * i.Y
}

func main() {
	input := strings.ReplaceAll(input, "\r\n", "\n")
	input = strings.Trim(input, "\n")
	fmt.Println("Day 1")
	start := time.Now()
	fmt.Println("Part 1: ", Part1(input))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", Part2(input))
	fmt.Println(time.Since(start))
}
