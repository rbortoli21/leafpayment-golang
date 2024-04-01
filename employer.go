package main

import (
	"time"
)

type Employer struct {
	Name         string     `json:"name"`
	EmployeeList []Employee `json:"employees"`
}

type Employee struct {
	ID                 int       `json:"id"`
	Name               string    `json:"name"`
	Position           string    `json:"position"`
	Salary             float64   `json:"salary"`
	Transportation     float64   `json:"transportation"`
	Children           int       `json:"children"`
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
	EmployeeID        string
	Name              string
	Position          string
	ExtraHours50      float64
	ExtraHours100     float64
	GrossSalary       float64
	TotalDiscounts    float64
	NetSalary         float64
	Fgts              float64
	ExtraHours        float64
	Inss              float64
	Irrf              float64
	Vt                float64
	Va                float64
	UnionContribution float64
	AuxCrecheBenefit  float64
	DSR               float64
}
