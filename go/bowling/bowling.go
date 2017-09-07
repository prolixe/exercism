package bowling

import "errors"

const testVersion = 1

type Frame struct {
	balls []int
}

func (f *Frame) IsSpare() bool {
	if f.balls[0] != 10 && f.balls[0]+f.balls[1] == 10 {
		return true
	}
	return false
}

func (f *Frame) IsStrike() bool {
	if f.balls[0] == 10 {
		return true
	}
	return false
}

type Game struct {
	f            [10]Frame
	currentFrame int
	currentBall  int
}

func NewGame() *Game {
	g := Game{currentBall: 0, currentFrame: 0}
	for i := range g.f {
		g.f[i] = Frame{balls: make([]int, 2)}
		if i == 9 {
			g.f[i] = Frame{balls: make([]int, 3)}
		}
	}
	return &g
}

func (g *Game) validate() error {

	if g.currentFrame != 9 {
		return errors.New("Score cannot be taken until the end of the game")
	}

	if g.currentBall < 2 {
		return errors.New("Score cannot be taken until the end of the game")
	}

	if g.f[9].balls[0] == 10 && g.f[9].balls[0] == 10 && g.currentBall != 3 {
		return errors.New("Score cannot be taken until the end of the game")
	}

	if g.f[9].IsSpare() && g.currentBall != 3 {
		return errors.New("Score cannot be taken until the end of the game")
	}
	return nil
}

func (g *Game) Score() (int, error) {

	score := 0

	if err := g.validate(); err != nil {
		return 0, err
	}

	for i, frame := range g.f {
		switch {
		case i == 9:
			score += frame.balls[0] + frame.balls[1] + frame.balls[2]
		case frame.IsStrike():
			score += 10 + g.f[i+1].balls[0]

			if g.f[i+1].balls[0] == 10 {
				if i == 8 {
					score += g.f[i+1].balls[1]
				} else {
					score += g.f[i+2].balls[0]
				}

			} else {
				score += g.f[i+1].balls[1]
			}
		case frame.IsSpare():
			score += 10 + g.f[i+1].balls[0]

		default:
			score += frame.balls[0] + frame.balls[1]

		}
	}

	return score, nil
}

func (g *Game) Frame() *Frame {
	return &g.f[g.currentFrame]
}

func (g *Game) Roll(pins int) error {

	if pins < 0 {
		return errors.New("cannot have a negative pin number")
	}
	if pins > 10 {
		return errors.New("cannot have a pins greater than 10")
	}

	if g.currentFrame > len(g.f) {
		return errors.New("reach the end of the game, cannot roll")
	}
	if g.currentFrame == 9 && g.currentBall > 2 {
		return errors.New("exceeded balls")
	}

	g.f[g.currentFrame].balls[g.currentBall] = pins
	g.currentBall++

	if g.currentFrame == 9 {

		if g.Frame().balls[0] == 10 && g.Frame().balls[1] != 10 && g.Frame().balls[1]+g.Frame().balls[2] > 10 {
			return errors.New("bonus roll exceed limit")
		}
		if (g.Frame().IsSpare() || g.Frame().IsStrike()) && g.currentBall > 3 {
			return errors.New("rolling after game over")
		}

		if !(g.Frame().IsSpare() || g.Frame().IsStrike()) && g.currentBall > 2 {
			return errors.New("rolling after game over")
		}
		return nil
	}

	if g.Frame().balls[0]+g.Frame().balls[1] > 10 {
		return errors.New("Pin count exceed pins on the lane")
	}

	if g.Frame().IsStrike() || g.Frame().IsSpare() || g.currentBall >= len(g.Frame().balls) {
		g.currentFrame++
		g.currentBall = 0
	}

	return nil
}
