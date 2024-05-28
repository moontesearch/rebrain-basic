package lession21

import "fmt"

type AmericanVelocity struct {
	m float64
	t float64
}

type EuropeanVelocity struct {
	m float64
	t float64
}

type Velocity interface {
	Velocity()
}

func (a AmericanVelocity) Velocity() string {
	var v float64 = ((a.m / a.t) * (3600 / 1000)) / 1.609
	return fmt.Sprintf("%.2f", v)
}

func (e EuropeanVelocity) Velocity() string {
	var v float64 = (e.m / e.t) * (3600 / 1000)
	return fmt.Sprintf("%.2f", v)
}

func Lession21() {
	AV := AmericanVelocity{m: 120.4, t: 1}
	EV := EuropeanVelocity{m: 130.0, t: 1}

	fmt.Println(AV.Velocity())
	fmt.Println(EV.Velocity())
}
