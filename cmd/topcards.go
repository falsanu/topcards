package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"

	"github.com/falsanu/topcards/model"
	"github.com/wcharczuk/go-chart"
)

var numberOfCards int

func main() {
	flag.IntVar(&numberOfCards, "n", 32, "Number of Cards to play witth")
	flag.Parse()
	fmt.Printf("Playing with %d Cards\n", numberOfCards)

	var staple []model.Card

	YValues := make([]int, numberOfCards, numberOfCards)

	for i := 0; i < numberOfCards; i++ {

		sample := TruncatedNormal(float64(numberOfCards/2), float64(numberOfCards/10), 0, float64(numberOfCards))
		fmt.Println(sample)

		YValues[int(sample)] = YValues[int(sample)] + 1
		fmt.Println(sample, YValues[int(sample)])
		staple = append(staple, model.Card{Value: sample, Id: i})

	}
	var x, y []float64
	for key, value := range YValues {
		x = append(x, float64(key))
		y = append(y, float64(value))
		fmt.Printf("%d,%d\n", key, value)
	}

	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: x,
				YValues: y,
			},
		},
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)

	// shuflle
	rand.Shuffle(len(staple), func(i, j int) { staple[i], staple[j] = staple[j], staple[i] })
	var playerA, playerB []model.Card

	// give to each player
	for i := 0; i < len(staple); i++ {
		if i%2 == 0 {
			playerA = append(playerA, staple[i])
		} else {
			playerB = append(playerB, staple[i])
		}
	}

	fmt.Println("Number of Cards A:", len(playerA))
	fmt.Println("Number of Cards B:", len(playerB))

	// start playing
	// who starts
	nextPlayer := rand.Intn(1)
	otherPlayer := 0

	var players [][]model.Card
	players = append(players, playerA)
	players = append(players, playerB)

	var tieStaple []model.Card
	var counter = 0

	for len(players[0]) != 0 && len(players[1]) != 0 {
		counter++
		fmt.Println("Run: ", counter)
		fmt.Printf("Staple A: %d, Staple B:%d\n", len(players[0]), len(players[1]))
		if nextPlayer == 0 {
			otherPlayer = 1
		} else if nextPlayer == 1 {
			otherPlayer = 0
		}

		cardA := players[nextPlayer][0]
		cardB := players[otherPlayer][0]

		stapleA := players[nextPlayer][1:len(players[nextPlayer])]
		stapleB := players[otherPlayer][1:len(players[otherPlayer])]
		fmt.Printf("Value A: %g, Value B:%g\n", cardA.Value, cardB.Value)
		if cardA.Value > cardB.Value {
			fmt.Println("Player Nextplayer")
			players[nextPlayer] = append(stapleA, cardA, cardB)
			players[otherPlayer] = stapleB
			if len(tieStaple) > 0 {
				// add all tie cards to otherPlayer
				players[nextPlayer] = append(stapleA, tieStaple...)
				tieStaple = nil

			}

		} else if cardB.Value > cardA.Value {
			fmt.Println("Player otherplayer")
			players[otherPlayer] = append(stapleB, cardA, cardB)
			players[nextPlayer] = stapleA
			if len(tieStaple) > 0 {
				// add all tie cards to otherPlayer
				players[otherPlayer] = append(stapleB, tieStaple...)
				tieStaple = nil
			}

			oldPlayer := nextPlayer
			nextPlayer = otherPlayer
			otherPlayer = oldPlayer

		} else {
			tieStaple = append(tieStaple, cardA, cardB)
			players[nextPlayer] = stapleA
			players[otherPlayer] = stapleB
			fmt.Println("Tie Staple lenght", len(tieStaple))
		}
		//time.Sleep(1 * time.Second)
	}
	fmt.Println("Number of total runs:", counter)
	// use while player 1 oder 2  loop

}

func TruncatedNormal(mean, stdDev, low, high float64) float64 {
	if low >= high {
		panic("high must be greater than low")
	}

	for {
		x := rand.NormFloat64()*stdDev + mean
		if low <= x && x < high {
			return x
		}
	}
}
