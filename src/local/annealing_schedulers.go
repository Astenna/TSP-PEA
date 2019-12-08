package local

import "math"

type AnnealingScheduler interface {
	NextTemperature(lastTemperature float64, step int) float64
}

type LinearAnnealing struct {
	Ratio float64
}

func (l LinearAnnealing) NextTemperature(lastTemperature float64, step int) float64 {
	return lastTemperature - l.Ratio
}

type GeometricAnnealing struct {
	Ratio float64
}

func (l GeometricAnnealing) NextTemperature(lastTemperature float64, step int) float64 {
	return lastTemperature * l.Ratio
}

type BoltzmanAnnealing struct {
	InitialTemperature float64
}

func (l BoltzmanAnnealing) NextTemperature(lastTemperature float64, step int) float64 {
	return l.InitialTemperature / (1 + math.Log(float64(1+step)))
}
