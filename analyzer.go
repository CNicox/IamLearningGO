package main

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
