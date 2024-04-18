package handler

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/rbortoli21/leafpayment-golang/models"
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
	employee.Employer = &employer

	if employee.Type != models.CLT {
		return models.Payroll{ReturnMessage: "Employee type not supported"}
	}

	payroll.EmployeePayrollDto.Name = employee.Name
	payroll.EmployeePayrollDto.Position = employee.Position

	payroll.EmployerPayrollDto.Name = employer.Name
	payroll.EmployerPayrollDto.Cnpj = employer.CNPJ

	payroll.DiscountsPayrollDto = calculateDiscounts(employee)
	payroll.AdditionsPayrollDto = calculateAddition(employee)
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
		discounts.FoodVoucher = calculateFoodVoucher(*employee.Employer)
	}
	if employee.HasTransportation {
		discounts.TransportVoucher = calculateTransportVoucher(*employee.Employer)
	}

	workload := GetLastWorkLoadByEmployeeID(employee.Id)
	
	if workload.DaysWorked < employee.Employer.Configurator.WorkDays {
		daysOff := employee.Employer.Configurator.WorkDays - workload.DaysWorked

		discounts.Absence.Ref = formatDecimal(employee.BaseSalary)
		discounts.Absence.Days = float64(daysOff)
		discounts.Absence.Value = formatDecimal((employee.BaseSalary / 22) * float64(daysOff))
	}

	discounts.DSR = models.DiscountDsrDto{
		Days:  float64(workload.MissedDSR),
		Ref:   formatDecimal(employee.BaseSalary / float64(employee.Employer.Configurator.WorkDays)),
		Total: formatDecimal((float64(workload.MissedDSR) * (employee.BaseSalary / 22))),
	}

	discounts.Total = calculateTotalDiscounts(discounts)

	return discounts
}

func calculateAddition(employee models.Employee) models.AdditionsPayrollDto {
	var additions models.AdditionsPayrollDto

	if employee.HasPericolous {
		additions.Dangerousness.Ref = employee.BaseSalary
		additions.Dangerousness.Percent = models.Level1Dangerousness.Percent
		additions.Dangerousness.Value = (models.Level1Dangerousness.Percent / 100) * employee.BaseSalary
	}

	switch employee.UnhealthinessLevel {
	case models.Level1Unhealthiness.Id:
		additions.Unhealthiness.Ref = employee.BaseSalary
		additions.Unhealthiness.Percent = models.Level1Unhealthiness.Percent
		additions.Unhealthiness.Value = (models.Level1Unhealthiness.Percent / 100) * employee.BaseSalary
	case models.Level2Unhealthiness.Id:
		additions.Unhealthiness.Ref = employee.BaseSalary
		additions.Unhealthiness.Percent = models.Level2Unhealthiness.Percent
		additions.Unhealthiness.Value = (models.Level2Unhealthiness.Percent / 100) * employee.BaseSalary
	case models.Level3Unhealthiness.Id:
		additions.Unhealthiness.Ref = employee.BaseSalary
		additions.Unhealthiness.Percent = models.Level3Unhealthiness.Percent
		additions.Unhealthiness.Value = (models.Level3Unhealthiness.Percent / 100) * employee.BaseSalary
	default:
		additions.Unhealthiness.Ref = 0
		additions.Unhealthiness.Percent = 0
		additions.Unhealthiness.Value = 0
	}

	employeeWorkload := GetLastWorkLoadByEmployeeID(employee.Id)
	hourValue := formatDecimal(employee.BaseSalary / float64(employeeWorkload.DaysWorked)) / float64(employee.HourlyWorkloadPerDay)
	
	if employeeWorkload.NocturnalHours > 0 {
		additions.Nocturnal.Ref = hourValue
		additions.Nocturnal.Percent = models.Level1Nocturnal.Percent
		additions.Nocturnal.Value = formatDecimal(hourValue + ((models.Level1Nocturnal.Percent / 100) * hourValue) * float64(employeeWorkload.NocturnalHours))
	}

	if employeeWorkload.WeekendHours > 0 {
		additions.Overtime.Weekend.Ref = hourValue
		additions.Overtime.Weekend.Percent = models.Level1OvertimeWeekends.Percent
		additions.Overtime.Weekend.Value = formatDecimal(hourValue + ((models.Level1OvertimeWeekends.Percent / 100) * hourValue) * float64(employeeWorkload.WeekendHours))
		additions.Overtime.Weekend.Quantity = float64(employeeWorkload.WeekendHours)

	}

	if employeeWorkload.ExtraHours > 0 {
		additions.Overtime.Default.Ref = hourValue
		additions.Overtime.Default.Percent = models.Level1OvertimeDefault.Percent
		additions.Overtime.Default.Value = formatDecimal(hourValue + ((models.Level1OvertimeDefault.Percent / 100) * hourValue) * float64(employeeWorkload.ExtraHours))
		additions.Overtime.Default.Quantity = float64(employeeWorkload.ExtraHours)
	}

	if employee.Dependents != nil && len(*employee.Dependents) > 0 {
		for _, dependent := range *employee.Dependents {
			var dependentAge int = time.Now().Year() - dependent.Birthday.Year()
			if dependentAge <= models.Level1FamilySalary.MaxAge || dependent.Invalid {
				var value = (models.Level1FamilySalary.Percent / 100) * models.Level1FamilySalary.SalaryBase
				additions.FamilySalary.Ref = models.Level1FamilySalary.SalaryBase
				additions.FamilySalary.Percent = models.Level1FamilySalary.Percent
				additions.FamilySalary.Total = formatDecimal(value)
			}
		}
	}

	if employee.Gender == models.Fem {
		var femaleEmployeesWithMoreThan16YearsOld = GetCountFemaleEmployeesWithMoreThan16YearsByEmployerID(*employee.EmployerID)
		if femaleEmployeesWithMoreThan16YearsOld > models.Level1ChildcareAssistance.FemaleEmployeesWithMoreThan16YearsOld {
			if employee.Dependents != nil && len(*employee.Dependents) > 0 {
				for _, dependent := range *employee.Dependents {
					var now = time.Now()
					var dependentAge int = int(now.Sub(dependent.Birthday).Hours() / 24 / 30)
					if dependentAge <= models.Level1ChildcareAssistance.MaxMonthsAge {
						additions.ChildcareAssistance.Ref = employee.BaseSalary
						additions.ChildcareAssistance.Percent = models.Level1ChildcareAssistance.Percent
						additions.ChildcareAssistance.Total = formatDecimal(employee.BaseSalary * (models.Level1ChildcareAssistance.Percent / 100))
					}
				}

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
		discountDto.Percent = models.Level1Incss.Percent
		discountDto.Ref = salary
		discountDto.Value = salary * (models.Level1Incss.Percent / 100)
	case salary >= models.Level2Incss.Min && salary <= models.Level2Incss.Max:
		discountDto.Percent = models.Level2Incss.Percent
		discountDto.Ref = salary
		discountDto.Value = salary * (models.Level2Incss.Percent / 100)
	case salary >= models.Level3Incss.Min && salary <= models.Level3Incss.Max:
		discountDto.Percent = models.Level3Incss.Percent
		discountDto.Ref = salary
		discountDto.Value = salary * (models.Level3Incss.Percent / 100)
	case salary >= models.Level4Incss.Min && salary <= models.Level4Incss.Max:
		discountDto.Percent = models.Level4Incss.Percent
		discountDto.Ref = salary
		discountDto.Value = models.Level4Incss.Value
	default:
		discountDto.Percent = 0
		discountDto.Ref = 0
		discountDto.Value = 0
	}

	return discountDto
}

func calculateIRRF(salary float64) models.DiscountDto {
	var discountDto models.DiscountDto

	switch {
	case salary >= models.Level1Irff.Min && salary <= models.Level1Irff.Max:
		discountDto.Percent = models.Level1Irff.Percent
		discountDto.Ref = salary
		discountDto.Value = salary * (models.Level1Irff.Percent / 100)
	case salary >= models.Level2Irff.Min && salary <= models.Level2Irff.Max:
		discountDto.Percent = models.Level2Irff.Percent
		discountDto.Ref = salary
		discountDto.Value = (salary * models.Level2Irff.Percent / 100)
	case salary >= models.Level3Irff.Min && salary <= models.Level3Irff.Max:
		discountDto.Percent = models.Level3Irff.Percent
		discountDto.Ref = salary
		discountDto.Value = (salary * models.Level3Irff.Percent / 100)
	case salary >= models.Level4Irff.Min && salary <= models.Level4Irff.Max:
		discountDto.Percent = models.Level4Irff.Percent
		discountDto.Ref = salary
		discountDto.Value = (salary * models.Level4Irff.Percent / 100)
	case salary >= models.Level5Irff.Min:
		discountDto.Percent = models.Level5Irff.Percent
		discountDto.Ref = salary
		discountDto.Value = (salary * models.Level5Irff.Percent / 100)
	default:
		discountDto.Percent = 0
		discountDto.Ref = 0
		discountDto.Value = 0
	}

	return discountDto
}

func calculateFGTS(salary float64) models.DiscountDto {
	return models.DiscountDto{Ref: salary, Percent: models.FgtsDiscount.Value, Value: (models.FgtsDiscount.Value / 100) * salary}
}

func calculateTransportVoucher(employer models.Employer) models.DiscountDto {
	return models.DiscountDto{Ref: employer.Configurator.TransportVoucher, Percent: models.VtDiscount.Value, Value: (models.VtDiscount.Value / 100) * employer.Configurator.TransportVoucher}
}

func calculateFoodVoucher(employer models.Employer) models.DiscountDto {
	return models.DiscountDto{Ref: employer.Configurator.FoodVoucher, Percent: models.VaDiscount.Value, Value: (models.VaDiscount.Value / 100) * employer.Configurator.FoodVoucher}
}

func calculateTotalDiscounts(discounts models.DiscountsPayrollDto) float64 {
	return formatDecimal(
		discounts.Inss.Value + discounts.Irff.Value + discounts.Fgts.Value + discounts.TransportVoucher.Value + discounts.FoodVoucher.Value)
}

func calculateTotalAddition(addition models.AdditionsPayrollDto) float64 {
	return formatDecimal(
		addition.Dangerousness.Value + addition.Unhealthiness.Value + addition.Nocturnal.Value + addition.Overtime.Weekend.Value + addition.Overtime.Default.Value + addition.FamilySalary.Total + addition.ChildcareAssistance.Total)
}

func calculateSalary(employee models.Employee, totalAdditions float64, totalDiscounts float64) models.SalaryPayrollDto {
	var salary models.SalaryPayrollDto

	salary.Gross = formatDecimal(employee.BaseSalary)
	salary.Net = formatDecimal(employee.BaseSalary + totalAdditions - totalDiscounts)

	return salary
}

func calculateWorkedHours(employee models.Employee) float64 {
	var employeeWorkload = GetLastWorkLoadByEmployeeID(employee.Id)

	return float64(employeeWorkload.DaysWorked * employee.HourlyWorkloadPerDay)
}

func formatDecimal(num float64) float64 {
	return math.Round(num*100) / 100
}
