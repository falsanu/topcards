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
		staple = append(staple, model.Card{Value: sample})

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
