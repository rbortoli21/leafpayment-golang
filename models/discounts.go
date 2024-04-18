package models

import (
	"math"
)

type Inss struct {
	Id      int
	Min     float64
	Max     float64
	Value   float64
	Percent float64
}

var (
	Level1Incss = Inss{0, 1412, 1752.81, 0, 8}
	Level2Incss = Inss{1, 1752.82, 2919.72, 0, 9}
	Level3Incss = Inss{2, 2919.73, 5839.45, 0, 11}
	Level4Incss = Inss{3, 5839.46, math.MaxInt32, 642.34, 0}
)

type Irff struct {
	Id      int
	Min     float64
	Max     float64
	Percent float64
}

var (
	Level1Irff = Irff{0, 1412, 1903.98, 0}
	Level2Irff = Irff{1, 1903.99, 2826.65, 7.5}
	Level3Irff = Irff{2, 2826.66, 3751.05, 15}
	Level4Irff = Irff{3, 3751.06, 4664.68, 22.5}
	Level5Irff = Irff{4, 4664.69, math.MaxInt32, 27.5}
)

type BenefitsDiscounts struct {
	Id    int
	Title string
	Value float64
}

var (
	VtDiscount   = BenefitsDiscounts{0, "Vale Transporte", 6}
	VaDiscount   = BenefitsDiscounts{1, "Vale Alimentação", 1}
	FgtsDiscount = BenefitsDiscounts{2, "FGTS", 8}
)
