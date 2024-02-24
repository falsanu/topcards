package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"

	"github.com/wcharczuk/go-chart"
)

var numberOfCards int

func main() {
	flag.IntVar(&numberOfCards, "n", 32, "Number of Cards to play witth")
	flag.Parse()
	fmt.Printf("Playing with %d Cards\n", numberOfCards)

	// var staple []model.Card
	// var XValues []float64
	//var YValues []float64

	// type byValue []model.Card
	YValues := make([]int, numberOfCards, numberOfCards)
	// func (a byValue) Len() float64           { return len(a) }
	// func (a byValue) Less(i, j float64) bool { return a[i].Value < a[j].Value }
	// func (a byValue) Swap(i, j float64)      { a[i], a[j] = a[j], a[i] }

	for i := 0; i < numberOfCards; i++ {

		sample := TruncatedNormal(float64(numberOfCards/2), float64(numberOfCards/10), 0, float64(numberOfCards))
		fmt.Println(sample)

		YValues[int(sample)] = YValues[int(sample)] + 1
		fmt.Println(sample, YValues[int(sample)])
		// staple = append(staple, model.Card{Value: sample})
		// XValues = append(XValues, +sample)
		//XValues = append(XValues, float64(i))

	}
	var x, y []float64
	// //slices.Sort(YValues)
	// // gett the keys out of YValues
	for key, value := range YValues {
		x = append(x, float64(key))
		y = append(y, float64(value))
		fmt.Printf("%d,%d\n", key, value)
	}

	// sort.Sort(byValue(staple))

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
func FuncVar(f32 []int) []float64 {
	f64 := make([]float64, len(f32))
	var f int
	var i int
	for i, f = range f32 {
		f64[i] = float64(f)
	}
	return f64
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
		// fmt.Println("missed!", x)
	}
}
