package main

import (
	"chekers/bot"
	"chekers/core"
	gamer2 "chekers/gamer"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type mind struct {
	kingCost    float64
	checkerCost float64
}

func test(level int) (float64, float64) {
	fmt.Println("init population")
	var population [8]mind
	population[0] = mind{1, 1}
	for i := 1; i < 8; i++ {
		population[i] = mind{rand.Float64() * 10, 1}
	}

	for generation := 0; generation < 5; generation++ {
		fmt.Println("Step")
		winners := grandCompare(population, level)
		population = createChilds(winners)
	}
	winners := grandCompare(population, level)
	return winners[0].kingCost, winners[0].checkerCost
}

func createChilds(winners [3]mind) [8]mind {
	var population [8]mind
	for i := range winners {
		population[i] = winners[i]
	}
	population[3] = mind{1, 1}
	population[4] = mind{rand.Float64() * 10, 1}
	population[5] = cross(winners[0], winners[1])
	population[6] = cross(winners[1], winners[2])
	population[7] = cross(winners[0], winners[2])
	return population
}

func cross(l, r mind) mind {
	return mind{(l.kingCost + r.kingCost) / 2, 1}
}

func grandCompare(population [8]mind, level int) [3]mind {
	var wg sync.WaitGroup
	var score [8]atomic.Int32
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if i == j {
				continue
			}
			wg.Add(1)
			i := i
			j := j
			mindi := population[i]
			mindj := population[j]
			go func() {
				for c := 0; c < 1; c++ {
					cmp := compare(mindi, mindj, level)
					if cmp == 1 {
						score[i].Add(2)
					} else if cmp == -1 {
						score[j].Add(2)
					} else {
						score[i].Add(1)
						score[j].Add(1)
					}
					fmt.Println(i, j, cmp)
				}
				wg.Done()
			}()
		}
	}
	wg.Wait()
	var top1, top2, top3 int
	for i := 1; i < 8; i++ {
		if score[i].Load() > score[top1].Load() {
			top3 = top2
			top2 = top1
			top1 = i
		} else if score[i].Load() > score[top2].Load() {
			top3 = top2
			top2 = i
		} else if score[i].Load() > score[top3].Load() {
			top3 = i
		}
	}
	return [3]mind{population[top1], population[top2], population[top3]}
}

func compare(l, r mind, level int) int {
	var bots [2]bot.Bot
	bots[0] = bot.Bot{bot.NewMinMaxV2(level, l.kingCost, l.checkerCost)}
	bots[1] = bot.Bot{bot.NewMinMaxV2(level, r.kingCost, r.checkerCost)}

	var c core.GameCore
	c.InitField(core.NewStandart8x8Field())
	c.InitTurnGamerId(0)

	var gamers [2]gamer2.Gamer
	gamers[0] = gamer2.Gamer{0, &c}
	gamers[1] = gamer2.Gamer{1, &c}

	for i := 0; i < 300; i++ {
		isFinished, winner := gamers[0].GetWinner()
		if isFinished {
			nullGamer := gamer2.Gamer{0, nil}
			if winner == gamers[0] {
				fmt.Println(i)
				return 1
			} else if winner == nullGamer {
				fmt.Println(i)
				return 0
			} else if winner == gamers[1] {
				fmt.Println(i)
				return -1
			}
		}
		if gamers[0].IsTurn() {
			bots[0].Move(gamers[0])
		} else {
			bots[1].Move(gamers[1])
		}
	}

	return 0
}

func main() {
	start := time.Now()
	compare(mind{1, 1}, mind{1, 1}, 3)
	fmt.Println(time.Now().Sub(start))
	fmt.Println("213")
	fmt.Println(test(3))

}

// for level 3 1 cmp ij ji - 6.3461017059905505 1
