package main

import (
	"sort"
)

type DealerAnalysis struct {
	// Inventory
	TotalCars          int
	AvailableCars      int
	SoldCars           int
	ReservedCars       int
	TotalInventoryUAH  float64 // sum of FinalPriceUAH() for all cars
	TotalInventoryUSD  float64
	AverageCarPriceUAH float64
	MostExpensiveCar   *Car
	CheapestCar        *Car

	// Employees
	TotalEmployees  int
	TotalPayrollUAH float64 // sum of TotalCost() for all employees
	EmployeesByRole map[EmployeeRole][]Employee

	// Deals
	TotalDeals      int
	TotalRevenueUAH float64
	TotalRevenueUSD float64
	AverageDealUAH  float64
	TopSeller       *Employee // employee with most deals
	DealsByPayment  map[PaymentType]int

	// Categories
	CategoryStats map[string]CategoryStat
}

type CategoryStat struct {
	Count      int
	RevenueUAH float64
	AvgUAH     float64
}

func Analyze() *DealerAnalysis {
	return nil
}

func (Dealer Dealer) FilterByStatus(status CarStatus) []Car {
	carsInStatus := []Car{}
	for _, item := range Dealer.Inventory {
		if item.Status == status {
			carsInStatus = append(carsInStatus, item)
		}
	}
	return carsInStatus
}

func (Dealer Dealer) FilterByPriceRange(minUAH, maxUAH float64) []Car {
	carsInStatus := []Car{}
	for _, item := range Dealer.Inventory {
		if minUAH < item.PriceUAH && item.PriceUAH < maxUAH {
			carsInStatus = append(carsInStatus, item)
		}
	}
	return carsInStatus
}

func (Dealer Dealer) FilterByCategory(category string) []Car {
	carsInStatus := []Car{}
	for _, item := range Dealer.Inventory {
		if item.Category == category {
			carsInStatus = append(carsInStatus, item)
		}
	}
	return carsInStatus
}

func (Dealer Dealer) SortByPrice(asc bool) []Car {
	items := Dealer.Inventory
	sort.Slice(items, func(i, j int) bool {
		return items[i].FinalPriceUAH() < items[j].FinalPriceUAH()
	})
	return items
}

func (Dealer Dealer) FindCarByVIN(vin string) (*Car, bool) {
	for _, item := range Dealer.Inventory {
		if item.VIN == vin {
			return &item, true
		}
	}
	return nil, false
}

func (Dealer Dealer) FindEmployeeByID(id string) (*Employee, bool) {
	for _, item := range Dealer.Employees {
		if item.ID == id {
			return &item, true
		}
	}
	return nil, false
}

func (Dealer Dealer) GetDealsByEmployee(empID string) []Deal {
	dealsClosedByEmp := []Deal{}
	for _, item := range Dealer.Deals {
		if item.EmployeeID == empID {
			dealsClosedByEmp = append(dealsClosedByEmp, item)
		}
	}
	return dealsClosedByEmp
}

func countMaxOccurences[T comparable](items []T) (int, T) {
	max := 0
	var value T
	for _, evaluatedItem := range items {
		counter := 0
		for _, item := range items {
			if item == evaluatedItem {
				counter++
			}
		}
		if counter > max {
			max = counter
			value = evaluatedItem
		}
	}
	return max, value
}

func (Dealer Dealer) CalcTopSeller() (int, *Employee) {
	EmpsWithDeals := []string{}
	for _, item := range Dealer.Deals {
		EmpsWithDeals = append(EmpsWithDeals, item.EmployeeID)
	}
	count, value := countMaxOccurences(EmpsWithDeals)
	var employee Employee
	for _, item := range Dealer.Employees {
		if item.ID == value {
			employee = item
		}
	}
	return count, &employee
}

// Q: ne znau norm li etot podhod ?
func (Dealer Dealer) CalcRevenueByCategory() map[string]CategoryStat {
	averagePriceByCarCategory := make(map[string]CategoryStat)
	statusToPrices := make(map[string][]int)

	for _, status := range allCarStatuses {
		carPrices := []int{}
		for _, car := range Dealer.Inventory {
			if car.Status == status {
				carPrices = append(carPrices, int(car.PriceUAH))
			}
		}
		statusToPrices[string(status)] = carPrices
	}
	for key, value := range statusToPrices {
		var sum int
		for _, item := range value {
			sum = sum + item
		}
		if len(value) == 0 {
			continue
		}
		average := CategoryStat{AvgUAH: float64(sum / len(value))}
		averagePriceByCarCategory[key] = average
	}

	return averagePriceByCarCategory
}

func (Dealer Dealer) CalcEmployeesByRole() map[EmployeeRole][]Employee {
	employeesByRole := make(map[EmployeeRole][]Employee)
	for _, role := range allEmployeeRoles {
		for _, employee := range Dealer.Employees {
			if employee.Role == role {
				employeesByRole[employee.Role] = append(employeesByRole[employee.Role], employee)
			}
		}
	}
	return employeesByRole
}

func (Dealer Dealer) CalcDealsByPayment() map[PaymentType]int {
	dealCountByPaymentType := make(map[PaymentType]int)
	for _, paymentType := range allPaymentTypes {
		for _, deal := range Dealer.Deals {
			if deal.PaymentType == paymentType {
				dealCountByPaymentType[paymentType]++
			}
		}
	}
	return dealCountByPaymentType
}

func (Dealer Dealer) CalcTotalPayroll() float64 {
	var sum float64
	for _, employee := range Dealer.Employees {
		sum = sum + employee.TotalCost()
	}
	return sum
}
