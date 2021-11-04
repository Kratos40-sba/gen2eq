package algorithm

import (
	"flag"
	"fmt"
	"github.com/Kratos40-sba/gen2eq/lib"
	"math/rand"
	"time"
)

const MUTATION = 0.75

var (
	a = flag.Float64("a", 5, "first variable")
	b = flag.Float64("b", 6, "second variable")
	c = flag.Float64("c", 1, "third variable")
)

func Start() {
	flag.Parse()
	start := time.Now()
	rand.Seed(time.Now().UTC().UnixNano())
	population := lib.InitPopulation(*a, *b, *c)
	solutionFound := false
	generations := 0
	for !solutionFound {
		generations++
		bestFit := lib.BestFit(population)
		fmt.Printf("\r Generation : %d | %f| Fitness : %.2f |", generations, bestFit.Dna, bestFit.Fitness)
		if bestFit.Fitness >= 0.9999 {
			solutionFound = true
		} else {
			bestFitness := bestFit.Fitness
			pool := lib.CreatePool(population, bestFitness)
			population = lib.Selection(pool, population, MUTATION)
		}
	}
	end := time.Since(start)
	fmt.Printf("\n Time : %s \n", end.String())
}
