package main

import (
	"strconv"
)

func GeneratePayroll(employer Employer) Payroll {
	var payroll Payroll
	payroll.Employer = employer
	payroll.EmployeeData = make([]EmployeePayroll, len(employer.EmployeeList))

	for i, employee := range employer.EmployeeList {
		employeePayroll := calculateEmployeePayroll(employee)
		payroll.EmployeeData[i] = employeePayroll
	}

	return payroll
}

func calculateEmployeePayroll(employee Employee) EmployeePayroll {
	var employeePayroll EmployeePayroll
	employeePayroll.EmployeeID = strconv.Itoa(employee.ID)
	employeePayroll.Name = employee.Name
	employeePayroll.Position = employee.Position
	nightHours := calculateNightHours(employee)
	employeePayroll.ExtraHours += calculateNightShiftBonus(employee.Salary, nightHours)

	if employee.Pericolous {
		employeePayroll.GrossSalary *= 1.30
	}

	if employee.AuxCreche && employee.Children > 0 {
		employeePayroll.GrossSalary += calculateAuxCrecheBenefit(employee.Children)
	}

	employeePayroll.GrossSalary += employee.DailyTravel
	grossSalary := calculateGrossSalary(employee)
	employeePayroll.GrossSalary = grossSalary

	discounts := calculateDiscounts(employee, grossSalary)
	employeePayroll.TotalDiscounts = discounts
	netSalary := grossSalary - discounts
	employeePayroll.NetSalary = netSalary

	employeePayroll.Fgts = calculateFGTS(grossSalary)
	employeePayroll.ExtraHours = calculateExtraHours(employee)
	employeePayroll.Inss = calculateINSS(grossSalary)
	employeePayroll.Irrf = calculateIRRF(netSalary)
	employeePayroll.Vt = calculateTransportationDiscount(employee.Transportation)
	employeePayroll.Va = calculateVADiscount(grossSalary)
	employeePayroll.UnionContribution = calculateUnionContribution(employee)

	return employeePayroll
}

func calculateNightHours(employee Employee) float64 {
	if employee.NightShift {
		entranceTime := float64(employee.EntranceTime.Hour())
		if entranceTime < 8 {
			entranceTime += 24
		}
		exitTime := float64(employee.ExitTime.Hour())
		if exitTime < 8 {
			exitTime += 24
		}
		return exitTime - entranceTime
	}
	return 0
}

func calculateGrossSalary(employee Employee) float64 {
	grossSalary := employee.Salary

	if employee.InsalubrityLevel == "minimum" {
		grossSalary *= 1.10
	} else if employee.InsalubrityLevel == "medium" {
		grossSalary *= 1.20
	} else if employee.InsalubrityLevel == "maximum" {
		grossSalary *= 1.40
	}

	if employee.ContactWithHazards {
		grossSalary *= 1.30
	}

	return grossSalary
}

func calculateDiscounts(employee Employee, grossSalary float64) float64 {
	return calculateINSS(grossSalary) + calculateIRRF(grossSalary) + calculateTransportationDiscount(employee.Transportation) + calculateVADiscount(grossSalary) + calculateUnionContribution(employee)
}

func calculateFGTS(grossSalary float64) float64 {
	return 0.08 * grossSalary
}

func calculateExtraHours(employee Employee) float64 {
	if employee.HourlyWorkload > 40 {
		extraHours := float64(employee.HourlyWorkload - 40)
		return extraHours * employee.Salary * 1.5
	}
	return 0
}

func calculateINSS(salary float64) float64 {
	if salary <= 1751.81 {
		return salary * 0.08
	} else if salary <= 2919.72 {
		return salary * 0.09
	} else if salary <= 5839.45 {
		return salary * 0.11
	} else {
		return 642.34
	}
}

func calculateIRRF(netSalary float64) float64 {
	if netSalary <= 1903.98 {
		return 0
	} else if netSalary <= 2826.65 {
		return (netSalary * 0.075) - 142.80
	} else if netSalary <= 3751.05 {
		return (netSalary * 0.15) - 354.80
	} else if netSalary <= 4664.68 {
		return (netSalary * 0.225) - 636.13
	} else {
		return (netSalary * 0.275) - 869.36
	}
}

func calculateTransportationDiscount(transportation float64) float64 {
	return 0.06 * transportation
}

func calculateVADiscount(grossSalary float64) float64 {
	return 0.1 * grossSalary
}

func calculateUnionContribution(employee Employee) float64 {
	if employee.Children > 0 || employee.Invalid {
		return 0.05 * employee.Salary
	}
	return 0
}

func calculateNightShiftBonus(salary, nightHours float64) float64 {
	return 0.20 * salary * nightHours
}

func calculateAuxCrecheBenefit(_ int) float64 {
	return 300.0
}
