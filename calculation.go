package main

// calculateEmployeePayroll calcula a folha de pagamento de um funcionário.
func calculateEmployeePayroll(employee Employee) EmployeePayroll {
	var employeePayroll EmployeePayroll

	// Definir os campos do funcionário na folha de pagamento
	employeePayroll.Employe.Name = employee.Name
	employeePayroll.Employe.Position = employee.Position

	// Calcular os descontos
	discounts := calculateDiscounts(employee)
	employeePayroll.Discounts = discounts
	employeePayroll.Discounts.TotalDiscounts = calculateTotalDiscounts(discounts)

	// Calcular as adições
	addition := calculateAddition(employee)
	employeePayroll.Addition = addition
	employeePayroll.Addition.TotalAddition = calculateTotalAddition(addition)

	// Calcular o salário bruto
	grossSalary := calculateGrossSalary(employee, addition)
	employeePayroll.Salary.SalaryGross = grossSalary

	// Calcular o salário líquido
	netSalary := grossSalary - employeePayroll.Discounts.TotalDiscounts
	employeePayroll.Salary.SalaryNet = netSalary

	// Calcular os demais campos
	employeePayroll.UnionContribution = calculateUnionContribution(employee)
	employeePayroll.ExtraHours = calculateExtraHours(employee)
	employeePayroll.Holiday = calculateHoliday(employee)
	employeePayroll.AuxCreche = calculateAuxCreche(employee)
	employeePayroll.Dependents = calculateDependents(employee)
	employeePayroll.WorkedHours = calculateWorkedHours(employee)
	employeePayroll.FamilySalary = calculateFamilySalary(employee)
	employeePayroll.Absence = calculateAbsence(employee)
	employeePayroll.Overtime = calculateOvertime(employee)
	employeePayroll.NightShift = calculateNightShift(employee)
	employeePayroll.Insalubrity = calculateInsalubrity(employee)
	employeePayroll.Pericolous = calculatePericolous(employee)
	employeePayroll.TravelAllowance = calculateTravelAllowance(employee)
	employeePayroll.DSR = calculateDSR(employee)

	return employeePayroll
}

// calculateNightHours calcula as horas noturnas trabalhadas por um funcionário.
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

// calculateGrossSalary calcula o salário bruto de um funcionário.
func calculateGrossSalary(employee Employee, addition Addition) float64 {
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

	grossSalary += addition.Insalubrity
	grossSalary += addition.Pericolous
	grossSalary += addition.NightShift

	return grossSalary
}

// calculateDiscounts calcula os descontos na folha de pagamento de um funcionário.
func calculateDiscounts(employee Employee) Discounts {
	var discounts Discounts

	discounts.Inss = calculateINSS(employee.Salary)
	discounts.Irff = calculateIRRF(employee.Salary)
	discounts.Fgts = calculateFGTS(employee.Salary)
	discounts.Vt = calculateTransportationDiscount(employee.Transportation)
	discounts.Va = calculateVADiscount(employee.Salary)
	discounts.Dsr = employee.DSR

	return discounts
}

// calculateTotalDiscounts calcula o total de descontos na folha de pagamento de um funcionário.
func calculateTotalDiscounts(discounts Discounts) float64 {
	return discounts.Inss + discounts.Irff + discounts.Fgts + discounts.Vt + discounts.Va + discounts.Dsr
}

// calculateAddition calcula as adições na folha de pagamento de um funcionário.
func calculateAddition(employee Employee) Addition {
	var addition Addition

	if employee.Pericolous {
		addition.Pericolous = 0.30 * employee.Salary
	}

	if employee.InsalubrityLevel == "maximum" {
		addition.Insalubrity = 0.40 * employee.Salary
	}

	if employee.NightShift {
		nightHours := calculateNightHours(employee)
		addition.NightShift = 0.20 * employee.Salary * nightHours
	}

	return addition
}

// calculateUnionContribution calcula a contribuição sindical de um funcionário.
func calculateUnionContribution(employee Employee) float64 {
	if employee.Dependents > 0 || employee.Invalid {
		return 0.05 * employee.Salary
	}
	return 0
}

// calculateExtraHours calcula as horas extras trabalhadas por um funcionário.
func calculateExtraHours(employee Employee) ExtraHours {
	var extraHours ExtraHours

	if employee.HourlyWorkload > 40 {
		extraHours.Total = float64(employee.HourlyWorkload-40) * 1.5 * employee.Salary
		extraHours.Holiday = float64(employee.HourlyWorkload-40) * 0.5 * employee.Salary
	}

	return extraHours
}

// calculateHoliday calcula o valor referente a feriados trabalhados por um funcionário.
func calculateHoliday(employee Employee) Holiday {
	var holiday Holiday

	holiday.Quantity = float64(0) // Ainda precisa ser implementado

	return holiday
}

// calculateAuxCreche calcula o benefício de auxílio-creche para um funcionário.
func calculateAuxCreche(employee Employee) AuxCreche {
	var auxCreche AuxCreche

	if employee.AuxCreche && employee.Dependents > 0 {
		auxCreche.Dependents = float64(employee.Dependents)
		auxCreche.Value = 300.0 * auxCreche.Dependents
	}

	return auxCreche
}

// calculateDependents calcula o valor referente aos dependentes de um funcionário.
func calculateDependents(employee Employee) Dependents {
	var dependents Dependents

	// Ainda precisa ser implementado

	return dependents
}

// calculateWorkedHours calcula o total de horas trabalhadas por um funcionário em diferentes períodos.
func calculateWorkedHours(employee Employee) WorkedHours {
	var workedHours WorkedHours

	// Ainda precisa ser implementado

	return workedHours
}

// calculateFamilySalary calcula o salário família de um funcionário.
func calculateFamilySalary(employee Employee) FamilySalary {
	var familySalary FamilySalary

	if employee.Dependents > 0 {
		familySalary.Total = calculateFamilySalaryPerDependent() * float64(employee.Dependents)
	}

	return familySalary
}

// calculateFamilySalaryPerDependent calcula o valor do salário família por dependente.
func calculateFamilySalaryPerDependent() float64 {
	// O valor do salário família por dependente é fixado pelo governo.
	return 48.62
}

// calculateINSS calcula o desconto do INSS com base no salário de um funcionário.
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

// calculateIRRF calcula o desconto do IRRF com base no salário de um funcionário.
func calculateIRRF(salary float64) float64 {
	if salary <= 1903.98 {
		return 0
	} else if salary <= 2826.65 {
		return (salary * 0.075) - 142.80
	} else if salary <= 3751.05 {
		return (salary * 0.15) - 354.80
	} else if salary <= 4664.68 {
		return (salary * 0.225) - 636.13
	} else {
		return (salary * 0.275) - 869.36
	}
}

// calculateFGTS calcula o FGTS com base no salário de um funcionário.
func calculateFGTS(salary float64) float64 {
	return 0.08 * salary
}

// calculateTransportationDiscount calcula o desconto de vale-transporte com base no valor do transporte de um funcionário.
func calculateTransportationDiscount(transportation float64) float64 {
	return 0.06 * transportation
}

// calculateVADiscount calcula o desconto de vale-alimentação com base no salário de um funcionário.
func calculateVADiscount(salary float64) float64 {
	return 0.1 * salary
}

func calculateTotalAddition(addition Addition) float64 {
	totalAddition := addition.Pericolous + addition.Insalubrity + addition.NightShift
	return totalAddition
}
