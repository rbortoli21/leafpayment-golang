package models

type Payroll struct {
	EmployeePayrollDto  EmployeePayrollDto  `json:"employee,omitempty"`
	EmployerPayrollDto  EmployerPayrollDto  `json:"employer,omitempty"`
	DiscountsPayrollDto DiscountsPayrollDto `json:"discounts,omitempty"`
	AdditionsPayrollDto AdditionsPayrollDto `json:"additions,omitempty"`
	SalaryPayrollDto    SalaryPayrollDto    `json:"salary,omitempty"`
	WorkedHours         float64             `json:"worked_hours,omitempty"`
	ReturnMessage       string              `json:"return_message,omitempty"`
}

type EmployeePayrollDto struct {
	Name     string `json:"name,omitempty"`
	Position string `json:"position,omitempty"`
}

type EmployerPayrollDto struct {
	Name string `json:"name,omitempty"`
	Cnpj string `json:"cnpj,omitempty"`
}

type DiscountsPayrollDto struct {
	Inss             DiscountDto    `json:"inss,omitempty"`
	Irff             DiscountDto    `json:"irrf,omitempty"`
	Fgts             DiscountDto    `json:"fgts,omitempty"`
	TransportVoucher DiscountDto    `json:"transport_voucher,omitempty"`
	FoodVoucher      DiscountDto    `json:"food_voucher,omitempty"`
	Absence          DiscountDto    `json:"absence,omitempty"`
	DSR              DiscountDsrDto `json:"dsr,omitempty"`
	Total            float64        `json:"total,omitempty"`
}

type DiscountDto struct {
	Ref     float64 `json:"ref,omitempty"`
	Percent float64 `json:"percent,omitempty"`
	Value   float64 `json:"value,omitempty"`
}

type DiscountDsrDto struct {
	Days  float64 `json:"days,omitempty"`
	Ref   float64 `json:"ref,omitempty"`
	Total float64 `json:"total,omitempty"`
}

type AdditionsPayrollDto struct {
	Dangerousness       AdditionDto                    `json:"dangerousness,omitempty"`
	Unhealthiness       AdditionDto                    `json:"unhealthiness,omitempty"`
	Nocturnal           AdditionDto                    `json:"nocturnal,omitempty"`
	Overtime            AdditionOvertimeDto            `json:"overtime,omitempty"`
	FamilySalary        AdditionFamilySalaryDto        `json:"family_salary,omitempty"`
	ChildcareAssistance AdditionChildcareAssistanceDto `json:"childcare_assistance,omitempty"`
	Total               float64                        `json:"total,omitempty"`
}

type AdditionDto struct {
	Ref     float64 `json:"ref,omitempty"`
	Percent float64 `json:"percent,omitempty"`
	Value   float64 `json:"value,omitempty"`
}

type AdditionOvertimeDto struct {
	Weekend AdditionNocturnalWeekendDto `json:"weekend,omitempty"`
	Default AdditionNocturnalDefaultDto `json:"default,omitempty"`
}

type AdditionNocturnalDefaultDto struct {
	Ref      float64 `json:"ref,omitempty"`
	Percent  float64 `json:"percent,omitempty"`
	Value    float64 `json:"value,omitempty"`
	Quantity float64 `json:"quantity,omitempty"`
}

type AdditionNocturnalWeekendDto struct {
	Ref      float64 `json:"ref,omitempty"`
	Percent  float64 `json:"percent,omitempty"`
	Value    float64 `json:"value,omitempty"`
	Quantity float64 `json:"quantity,omitempty"`
}

type AdditionFamilySalaryDto struct {
	Ref        float64                `json:"ref,omitempty"`
	Total      float64                `json:"total,omitempty"`
	Percent    float64                `json:"percent,omitempty"`
	Dependents []AdditionDependentDto `json:"dependents,omitempty"`
}

type AdditionChildcareAssistanceDto struct {
	Ref     float64 `json:"ref,omitempty"`
	Percent float64 `json:"percent,omitempty"`
	Total   float64 `json:"total,omitempty"`
}

type AdditionDependentDto struct {
	Name  string  `json:"name,omitempty"`
	Value float64 `json:"value,omitempty"`
}

type SalaryPayrollDto struct {
	Gross float64 `json:"gros_salary,omitempty"`
	Net   float64 `json:"net_salary,omitempty"`
}
