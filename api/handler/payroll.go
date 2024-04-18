package handler

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/ghenoo/folhadepagamento/models"
)

func GeneratePayrollHandler(w http.ResponseWriter, r *http.Request) {
	employeeIDStr := r.URL.Path[len("/generate-payroll/"):]

	employeeId, err := strconv.Atoi(employeeIDStr)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	payroll := calculateEmployeePayroll((uint(employeeId)))

	responseBody, err := json.Marshal(payroll)
	if err != nil {
		http.Error(w, "Failed to encode response body", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func calculateEmployeePayroll(employeeId uint) models.Payroll {
	var payroll models.Payroll

	employee := getEmployeeById(employeeId)
	employer := getEmployerById(*employee.EmployerID)

	if employee.Type != models.CLT {
		return models.Payroll{ReturnMessage: "Employee type not supported"}
	}

	payroll.EmployeePayrollDto.Name = employee.Name
	payroll.EmployeePayrollDto.Position = employee.Position

	payroll.EmployerPayrollDto.Name = employer.Name
	payroll.EmployerPayrollDto.Cnpj = employer.CNPJ

	payroll.DiscountsPayrollDto = calculateDiscounts(employee)
	payroll.AdditionsPayrollDto = calculateAddition(employee, employer)
	payroll.SalaryPayrollDto = calculateSalary(employee, payroll.AdditionsPayrollDto.Total, payroll.DiscountsPayrollDto.Total)
	payroll.WorkedHours = calculateWorkedHours(employee)

	return payroll
}

func calculateDiscounts(employee models.Employee) models.DiscountsPayrollDto {
	var discounts models.DiscountsPayrollDto

	discounts.Inss = calculateINSS(employee.BaseSalary)
	discounts.Irff = calculateIRRF(employee.BaseSalary)
	discounts.Fgts = calculateFGTS(employee.BaseSalary)
	if employee.HasAlimentation {
		discounts.FoodVoucher = calculateFoodVoucher(employee.BaseSalary)
	}
	if employee.HasTransportation {
		discounts.TransportVoucher = calculateTransportVoucher(employee.BaseSalary)
	}

	discounts.Total = calculateTotalDiscounts(discounts)

	return discounts
}

func calculateAddition(employee models.Employee, employer models.Employer) models.AdditionsPayrollDto {
	var additions models.AdditionsPayrollDto

	if employee.HasPericolous {
		additions.Dangerousness.Ref = models.Level1Dangerousness.Percent
		additions.Dangerousness.Value = (models.Level1Dangerousness.Percent / 100) * employee.BaseSalary
	}

	switch employee.UnhealthynessLevel {
	case models.Level1Unhealthyness.Id:
		additions.Unhealthyness.Ref = models.Level1Unhealthyness.Percent
		additions.Unhealthyness.Value = (models.Level1Unhealthyness.Percent / 100) * employee.BaseSalary
	case models.Level2Unhealthyness.Id:
		additions.Unhealthyness.Ref = models.Level2Unhealthyness.Percent
		additions.Unhealthyness.Value = (models.Level2Unhealthyness.Percent / 100) * employee.BaseSalary
	case models.Level3Unhealthyness.Id:
		additions.Unhealthyness.Ref = models.Level3Unhealthyness.Percent
		additions.Unhealthyness.Value = (models.Level3Unhealthyness.Percent / 100) * employee.BaseSalary
	default:
		additions.Unhealthyness.Ref = 0
		additions.Unhealthyness.Value = 0
	}

	var employeeWorkload models.EmployeeWorkload

	if employeeWorkload.NocturnalHours > 0 {
		additions.Nocturnal.Ref = models.Level1Nocturnal.Percent
		additions.Nocturnal.Value = (models.Level1Nocturnal.Percent / 100) * employee.BaseSalary
	}

	if employeeWorkload.ExtraHours > 0 {
		if employeeWorkload.DaysWorked > employer.Configurator.WorkDays {
			additions.Overtime.Weekend.Ref = models.Level1OvertimeWeekends.Percent
			additions.Overtime.Weekend.Value = (models.Level1OvertimeWeekends.Percent / 100) * employee.BaseSalary
			additions.Overtime.Weekend.Quantity = float64(employeeWorkload.ExtraHours)
		} else {
			additions.Overtime.Default.Ref = models.Level1OvertimeDefault.Percent
			additions.Overtime.Default.Value = (models.Level1OvertimeDefault.Percent / 100) * employee.BaseSalary
			additions.Overtime.Default.Quantity = float64(employeeWorkload.ExtraHours)
		}
	}

	if employee.Dependents != nil && len(*employee.Dependents) > 0 {
		for _, dependent := range *employee.Dependents {
			var dependentAge int = time.Now().Year() - dependent.Birthday.Year()
			if dependentAge <= models.Level1FamilySalary.MaxAge || dependent.Invalid {
				var value = (models.Level1FamilySalary.Percent / 100) * employee.BaseSalary
				additions.FamilySalary.Dependents =
					append(
						additions.FamilySalary.Dependents,
						models.AdditionDependentDto{Name: dependent.Name, Value: value},
					)
				additions.FamilySalary.Total += value
			}
		}
	}

	var femaleEmployeesWithMoreThan16YearsOld []models.Employee
	if employer.Employees != nil && len(*employer.Employees) > 0 {
		for _, employeeFor := range *employer.Employees {
			if employeeFor.Gender == models.Fem && employeeFor.Birthday.Year() > 16 {
				femaleEmployeesWithMoreThan16YearsOld = append(femaleEmployeesWithMoreThan16YearsOld, employeeFor)
			}
		}
	}

	if len(femaleEmployeesWithMoreThan16YearsOld) >= models.Level1ChildcareAssistance.FemaleEmployeesWithMoreThan16YearsOld {
		for _, femaleEmployee := range femaleEmployeesWithMoreThan16YearsOld {
			var dependentAge int = time.Now().Year() - femaleEmployee.Birthday.Year()
			if dependentAge <= models.Level1ChildcareAssistance.MaxMonthsAge {
				var value = (models.Level1ChildcareAssistance.Percent / 100) * femaleEmployee.BaseSalary
				additions.ChildcareAssistance.Dependents =
					append(
						additions.ChildcareAssistance.Dependents,
						models.AdditionDependentDto{Name: femaleEmployee.Name, Value: value},
					)
				additions.ChildcareAssistance.Total += value
			}
		}
	}

	additions.Total = calculateTotalAddition(additions)

	return additions
}

func calculateINSS(salary float64) models.DiscountDto {
	var discountDto models.DiscountDto

	switch {
	case salary >= models.Level1Incss.Min && salary <= models.Level1Incss.Max:
		discountDto.Ref = models.Level1Incss.Percent
		discountDto.Value = salary * (models.Level1Incss.Percent / 100)
	case salary >= models.Level2Incss.Min && salary <= models.Level2Incss.Max:
		discountDto.Ref = models.Level2Incss.Percent
		discountDto.Value = salary * (models.Level2Incss.Percent / 100)
	case salary >= models.Level3Incss.Min && salary <= models.Level3Incss.Max:
		discountDto.Ref = models.Level3Incss.Percent
		discountDto.Value = salary * (models.Level3Incss.Percent / 100)
	case salary >= models.Level4Incss.Min && salary <= models.Level4Incss.Max:
		discountDto.Ref = models.Level4Incss.Percent
		discountDto.Value = models.Level4Incss.Value
	default:
		discountDto.Ref = 0
		discountDto.Value = 0
	}

	return discountDto
}

func calculateIRRF(salary float64) models.DiscountDto {
	var discountDto models.DiscountDto

	switch {
	case salary >= models.Level1Irff.Min && salary <= models.Level1Irff.Max:
		discountDto.Ref = models.Level1Irff.Percent
		discountDto.Value = salary * (models.Level1Irff.Percent / 100)
	case salary >= models.Level2Irff.Min && salary <= models.Level2Irff.Max:
		discountDto.Ref = models.Level2Irff.Percent
		discountDto.Value = (salary * models.Level2Irff.Percent / 100)
	case salary >= models.Level3Irff.Min && salary <= models.Level3Irff.Max:
		discountDto.Ref = models.Level3Irff.Percent
		discountDto.Value = (salary * models.Level3Irff.Percent / 100)
	case salary >= models.Level4Irff.Min && salary <= models.Level4Irff.Max:
		discountDto.Ref = models.Level4Irff.Percent
		discountDto.Value = (salary * models.Level4Irff.Percent / 100)
	case salary >= models.Level5Irff.Min:
		discountDto.Ref = models.Level5Irff.Percent
		discountDto.Value = (salary * models.Level5Irff.Percent / 100)
	default:
		discountDto.Ref = 0
		discountDto.Value = 0
	}

	return discountDto
}

func calculateFGTS(salary float64) models.DiscountDto {
	return models.DiscountDto{Ref: models.FgtsDiscount.Value, Value: (models.FgtsDiscount.Value / 100) * salary}
}

func calculateTotalDiscounts(discounts models.DiscountsPayrollDto) float64 {
	return discounts.Inss.Value + discounts.Irff.Value + discounts.Fgts.Value + discounts.TransportVoucher.Value + discounts.FoodVoucher.Value
}

func calculateTransportVoucher(salary float64) models.DiscountDto {
	return models.DiscountDto{Ref: models.VtDiscount.Value, Value: (models.VtDiscount.Value / 100) * salary}
}

func calculateFoodVoucher(salary float64) models.DiscountDto {
	return models.DiscountDto{Ref: models.VaDiscount.Value, Value: (models.VaDiscount.Value / 100) * salary}
}

func calculateTotalAddition(addition models.AdditionsPayrollDto) float64 {
	return addition.Dangerousness.Value + addition.Unhealthyness.Value + addition.Nocturnal.Value + addition.Overtime.Weekend.Value + addition.Overtime.Default.Value + addition.FamilySalary.Total + addition.ChildcareAssistance.Total
}

func calculateSalary(employee models.Employee, totalAdditions float64, totalDiscounts float64) models.SalaryPayrollDto {
	var salary models.SalaryPayrollDto

	salary.Gross = employee.BaseSalary
	salary.Net = formatDecimal(employee.BaseSalary + totalAdditions - totalDiscounts)

	return salary
}

func calculateWorkedHours(employee models.Employee) float64 {
	var employeeWorkload models.EmployeeWorkload

	return float64(employeeWorkload.DaysWorked * employee.HourlyWorkloadPerDay)
}

func formatDecimal(num float64) float64 {
	return math.Round(num*100) / 100
}
