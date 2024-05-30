package main

import (
	"errors"
	"fmt"
	"math"
)

type parrotType int

const (
	TypeEuropean      parrotType = 1
	TypeAfrican       parrotType = 2
	TypeNorwegianBlue parrotType = 3
)

// Parrot has a Speed.
type Parrot interface {
	Speed() (float64, error)
}

type mixedParrot struct {
	numberOfCoconuts int
	voltage          float64
	nailed           bool
}

func (parrot mixedParrot) computeBaseSpeedForVoltage(voltage float64) float64 {
	return math.Min(24.0, voltage*parrot.baseSpeed())
}

func (parrot mixedParrot) loadFactor() float64 {
	return 9.0
}

func (parrot mixedParrot) baseSpeed() float64 {
	return 12.0
}

var _ Parrot = (*EuropeanParrot)(nil)

type EuropeanParrot struct {
	mixedParrot
}

func (e EuropeanParrot) Speed() (float64, error) {
	if e.nailed {
		return 0, nil
	}

	return e.baseSpeed(), nil
}

var _ Parrot = (*AfricanParrot)(nil)

type AfricanParrot struct {
	mixedParrot
}

func (a AfricanParrot) Speed() (float64, error) {
	if a.nailed {
		return 0, nil
	}

	return math.Max(0, a.baseSpeed()-a.loadFactor()*float64(a.numberOfCoconuts)), nil
}

var _ Parrot = (*NorwegianParrot)(nil)

type NorwegianParrot struct {
	mixedParrot
}

func (n NorwegianParrot) Speed() (float64, error) {
	if n.nailed {
		return 0, nil
	}

	return n.computeBaseSpeedForVoltage(n.voltage), nil
}

func ParrotFactory(pType parrotType) (Parrot, error) {
	switch pType {
	case TypeEuropean:
		return EuropeanParrot{
			mixedParrot{
				voltage: 220.0,
			},
		}, nil
	case TypeAfrican:
		return AfricanParrot{
			mixedParrot{
				numberOfCoconuts: 2,
				voltage:          120.0,
				nailed:           true,
			},
		}, nil
	case TypeNorwegianBlue:
		return NorwegianParrot{
			mixedParrot{
				numberOfCoconuts: 4,
				voltage:          120.0,
			},
		}, nil
	default:
		return nil, errors.New("should be unreachable")
	}
}

func main() {
	parrot, err := ParrotFactory(1)
	if err != nil {
		panic(err)
	}

	speed, err := parrot.Speed()
	if err != nil {
		return
	}

	fmt.Println(speed)
}
