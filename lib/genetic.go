package lib

import (
	"math"
	"math/rand"
	"time"
)

type Chromosome struct {
	Dna     []float64
	Fitness float64
	A, B, C float64
}

func generate(a, b, c float64) (ch Chromosome) {
	dna := make([]float64, 2)
	for i := 0; i < 2; i++ {
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		dna[i] = r.NormFloat64()
	}
	ch = Chromosome{
		Dna:     dna,
		Fitness: 0.0,
		A:       a,
		B:       b,
		C:       c,
	}
	//ch.setVars(a,b,c)
	return ch
}
func (ch *Chromosome) fitness() float64 {
	var fitRate float64
	fitRate = (ch.A * (math.Pow(ch.Dna[0], 2))) + (ch.B * (ch.Dna[1])) + ch.C
	fitRate = math.Abs(fitRate)

	return 1 / (fitRate + 1)
}
func InitPopulation(a, b, c float64) []Chromosome {
	population := make([]Chromosome, 20)
	for i := 0; i < 20; i++ {
		population[i] = generate(a, b, c)
		population[i].Fitness = population[i].fitness()
	}
	return population
}
func CreatePool(population []Chromosome, bestFit float64) (pool []Chromosome) {
	pool = make([]Chromosome, 0)
	for i := 0; i < 20; i++ {
		population[i].fitness()
		iterations := int(population[i].Fitness/bestFit) * 100
		for j := 0; j < iterations; j++ {
			pool = append(pool, population[i])
		}
	}
	return pool
}
func crossover(ch1, ch2 Chromosome) (ch *Chromosome) {
	ch = &Chromosome{
		Dna:     make([]float64, 2),
		Fitness: 0.0,
		A:       ch1.A,
		B:       ch1.B,
		C:       ch1.C,
	}
	randomNumber := rand.Intn(2)
	for i := 0; i < 2; i++ {
		if i > randomNumber {
			ch.Dna[i] = ch1.Dna[i]
		} else {
			ch.Dna[i] = ch2.Dna[i]
		}
	}
	return ch
}
func (ch *Chromosome) mutate(mutationRate float64) *Chromosome {

	for i := 0; i < 2; i++ {
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		if rand.Float64() < mutationRate {
			ch.Dna[i] = r.NormFloat64()
		}
	}
	return ch
}
func BestFit(population []Chromosome) Chromosome {
	bestOne := 0.0
	index := 0
	for i := 0; i < len(population); i++ {
		if population[i].Fitness > bestOne {
			bestOne = population[i].Fitness
			index = i
		}
	}
	return population[index]
}
func Selection(pool, population []Chromosome, mutationRate float64) []Chromosome {
	l := len(population)
	nextGeneration := make([]Chromosome, l)
	for i := 0; i < l; i++ {
		r1, r2 := rand.Intn(len(pool)), rand.Intn(len(pool))
		a := pool[r1]
		b := pool[r2]
		child := crossover(a, b).mutate(mutationRate)
		ch := &Chromosome{
			Dna:     child.Dna,
			Fitness: child.fitness(),
			A:       child.A,
			B:       child.B,
			C:       child.C,
		}
		nextGeneration[i] = *ch
	}
	return nextGeneration
}
