package models

type Payroll struct {
	EmployeePayrollDto  EmployeePayrollDto  `json:"employee"`
	EmployerPayrollDto  EmployerPayrollDto  `json:"employer"`
	DiscountsPayrollDto DiscountsPayrollDto `json:"discounts"`
	AdditionsPayrollDto AdditionsPayrollDto `json:"additions"`
	SalaryPayrollDto    SalaryPayrollDto    `json:"salary"`
	WorkedHours         float64             `json:"worked_hours"`
	ReturnMessage       string              `json:"return_message"`
}

type EmployeePayrollDto struct {
	Name     string `json:"name"`
	Position string `json:"position"`
}

type EmployerPayrollDto struct {
	Name string `json:"name"`
	Cnpj string `json:"cnpj"`
}

type DiscountsPayrollDto struct {
	Inss             DiscountDto `json:"inss"`
	Irff             DiscountDto `json:"irrf"`
	Fgts             DiscountDto `json:"fgts"`
	TransportVoucher DiscountDto `json:"transport_voucher"`
	FoodVoucher      DiscountDto `json:"food_voucher"`
	Absence          DiscountDto `json:"absence"`
	Total            float64     `json:"total"`
}

type DiscountDto struct {
	Ref   float64 `json:"ref"`
	Value float64 `json:"value"`
}

type AdditionsPayrollDto struct {
	Dangerousness       AdditionDto             `json:"dangerousness"`
	Unhealthyness       AdditionDto             `json:"unhealthyness"`
	Nocturnal           AdditionDto             `json:"nocturnal"`
	Overtime            AdditionOvertimeDto     `json:"overtime"`
	FamilySalary        AdditionFamilySalaryDto `json:"family_salary"`
	ChildcareAssistance AdditionFamilySalaryDto `json:"childcare_assistance"`
	Total               float64                 `json:"total"`
}

type AdditionDto struct {
	Ref   float64 `json:"ref"`
	Value float64 `json:"value"`
}

type AdditionOvertimeDto struct {
	Weekend AdditionNocturnalWeekendDto `json:"weekend"`
	Default AdditionNocturnalDefaultDto `json:"default"`
}

type AdditionNocturnalDefaultDto struct {
	Ref      float64 `json:"ref"`
	Value    float64 `json:"value"`
	Quantity float64 `json:"quantity"`
}

type AdditionNocturnalWeekendDto struct {
	Ref      float64 `json:"ref"`
	Value    float64 `json:"value"`
	Quantity float64 `json:"quantity"`
}

type AdditionFamilySalaryDto struct {
	Total      float64                `json:"total"`
	Dependents []AdditionDependentDto `json:"dependents"`
}

type AdditionDependentDto struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type SalaryPayrollDto struct {
	Gross float64 `json:"grossSalary"`
	Net   float64 `json:"netSalary"`
}
