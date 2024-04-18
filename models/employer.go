package models

type Employer struct {
	Id           uint                  `json:"id" gorm:"primaryKey"`
	Name         string                `json:"name"`
	CNPJ         string                `json:"cnpj"`
	Configurator *EmployerConfigurator `json:"configurator" gorm:"foreignKey:EmployerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Employees    *[]Employee           `json:"employees" gorm:"foreignKey:EmployerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type EmployerConfigurator struct {
	Id               uint    `json:"id" gorm:"primaryKey"`
	TransportVoucher float64 `json:"transport_voucher"`
	FoodVoucher      float64 `json:"food_voucher"`
	WorkDays         int     `json:"work_days"`
	EmployerID       uint    `json:"employer"`
}
