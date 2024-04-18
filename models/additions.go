package models

import "time"

type Dangerousness struct {
	Id      int     `json:"id"`
	Percent float64 `json:"percent"`
}

var (
	Level1Dangerousness = Dangerousness{0, 30}
)

type Unhealthiness struct {
	Id      int     `json:"id"`
	Title   string  `json:"title"`
	Percent float64 `json:"percent"`
}

var (
	Level1Unhealthiness = Unhealthiness{0, "Insalubridade de grau mínimo", 10}
	Level2Unhealthiness = Unhealthiness{1, "Insalubridade de grau médio", 20}
	Level3Unhealthiness = Unhealthiness{2, "Insalubridade de grau máximo", 40}
)

type Nocturnal struct {
	Id      int       `json:"id"`
	Begin   time.Time `json:"begin_date"`
	End     time.Time `json:"end_date"`
	Percent float64   `json:"percent"`
}

var (
	Level1Nocturnal = Nocturnal{0, time.Now(), time.Now(), 20}
)

type Overtime struct {
	Id      int     `json:"id"`
	Percent float64 `json:"percent"`
}

var (
	Level1OvertimeDefault  = Overtime{0, 50}
	Level1OvertimeWeekends = Overtime{1, 100}
)

type FamilySalary struct {
	Id         int     `json:"id"`
	MaxAge     int     `json:"max_age"`
	Invalid    bool    `json:"invalid"`
	Percent    float64 `json:"percent"`
	SalaryBase float64 `json:"salary_base"`
}

var (
	Level1FamilySalary = FamilySalary{0, 14, true, 5, 1412}
)

type ChildcareAssistance struct {
	Id                                    int     `json:"id"`
	FemaleEmployeesWithMoreThan16YearsOld int     `json:"femaleEmployeesWithMoreThan16YearsOld"`
	MaxMonthsAge                          int     `json:"maxMonthsAge"`
	Percent                               float64 `json:"percent"`
}

var (
	Level1ChildcareAssistance = ChildcareAssistance{0, 30, 6, 5}
)
