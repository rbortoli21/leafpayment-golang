package models

import (
	"time"
	
)

type Employee struct {
	Id                   uint                `json:"id" gorm:"primaryKey"`
	Name                 string              `json:"name"`
	Position             string              `json:"position"`
	Gender               Sex                 `json:"gender" gorm:"type:varchar(100)"`
	Type                 EmployeeType        `json:"type" gorm:"type:varchar(100)"`
	Birthday             time.Time           `json:"birthday"`
	BaseSalary           float64             `json:"base_salary"`
	HasTransportation    bool                `json:"transportation"`
	HasAlimentation      bool                `json:"alimentation"`
	HasPericolous        bool                `json:"pericolous"`
	UnhealthynessLevel   int                 `json:"unhealthyness_level"`
	HourlyWorkloadPerDay int                 `json:"hourly_workload"`
	EmployerID           *uint               `json:"-"`
	Employer             *Employer           `json:"employer"`
	Dependents           *[]Dependent        `json:"dependents" gorm:"foreignkey:EmployeeID;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Workloads            *[]EmployeeWorkload `json:"workloads" gorm:"foreignkey:EmployeeID;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type EmployeeType string

const (
	CLT EmployeeType = "CLT"
	PJ  EmployeeType = "PJ"
)

type Dependent struct {
	Id         uint      `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	Invalid    bool      `json:"invalidate"`
	Birthday   time.Time `json:"birthday"`
	EmployeeID uint      `json:"employee_id"`
}

type Sex string

const (
	Masc Sex = "Masc"
	Fem  Sex = "Fem"
)

type EmployeeWorkload struct {
	Id             uint `json:"id" gorm:"primaryKey"`
	EmployeeID     uint `json:"employee_id"`
	DaysWorked     int  `json:"daysNotWorked"`
	NocturnalHours int  `json:"nocturnalHours"`
	ExtraHours     int  `json:"extraHours"`
}
