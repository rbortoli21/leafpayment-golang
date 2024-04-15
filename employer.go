// employer.go
package main

import (
	"math"
	"time"
)

type Sex int

const (
	Masculino Sex = iota
	Feminino
	Outro
)

type Employer struct {
	NameEmployer string     `json:"name"`
	CNPJ         string     `json:"cnpj"`
	EmployeeList []Employee `json:"employees"`
}

type Employee struct {
	ID                 int       `json:"id"`
	Name               string    `json:"name"`
	Position           string    `json:"position"`
	Gender             Sex       `json:"gender"`
	Salary             float64   `json:"salary"`
	Transportation     float64   `json:"transportation"`
	Dependents         int       `json:"dependents"`
	Invalid            bool      `json:"invalid"`
	InsalubrityLevel   string    `json:"insalubrity_level"`
	ContactWithHazards bool      `json:"contact_with_hazards"`
	HourlyWorkload     int       `json:"hourly_workload"`
	NightShift         bool      `json:"night_shift"`
	EntranceTime       time.Time `json:"entrance_time"`
	ExitTime           time.Time `json:"exit_time"`
	Pericolous         bool      `json:"pericolous"`
	DailyTravel        float64   `json:"daily_travel"`
	AuxCreche          bool      `json:"aux_creche"`
	DSR                float64   `json:"dsr"`
}

type EmployeePayroll struct {
	Employer     Employer
	EmployeeList []Employee
}

type Employe struct {
	Name     string
	Position string
}

type Discounts struct {
	Inss           float64 `json:"inss"`
	Irff           float64 `json:"irrf"`
	Fgts           float64 `json:"fgts"`
	Vt             float64 `json:"vt"`
	Va             float64 `json:"va"`
	Dsr            float64 `json:"dsr"`
	TotalDiscounts float64 `json:"total"`
}

type Addition struct {
	Pericolous    float64 `json:"pericolous"`
	Insalubrity   float64 `json:"insalubrity"`
	NightShift    float64 `json:"night_shift"`
	TotalAddition float64 `json:"total"`
}

type Salary struct {
	SalaryGross float64 `json:"gross_salary"`
	SalaryNet   float64 `json:"net_salary"`
}

type ExtraHours struct {
	Holiday float64 `json:"holiday"`
	Total   float64 `json:"total"`
}

type Holiday struct {
	Quantity float64 `json:"quantity"`
}

type AuxCreche struct {
	Dependents float64 `json:"dependents"`
	Value      float64 `json:"value"`
}

type Dependents struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type WorkedHours struct {
	Month float64 `json:"month"`
	Week  float64 `json:"week"`
	Day   float64 `json:"day"`
}

type FamilySalary struct {
	Total      float64      `json:"total"`
	Dependents []Dependents `json:"dependents"`
}

func formatDecimal(num float64) float64 {
	return math.Round(num*100) / 100
}
