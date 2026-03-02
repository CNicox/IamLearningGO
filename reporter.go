package main

import (
	"fmt"
	"io"
)

/*
1. Declare interface  Reporter { Report(w io.Writer) error }.

2. Implement DealerReporter — receives *Dealer and *DealerAnalysis.
   Method Report(w io.Writer) error formats the report as shown above.

3. In main.go use io.MultiWriter(os.Stdout, file) for dual output.

4. Implement a second Reporter — JSONReporter:
   • Encodes DealerAnalysis as JSON with json.NewEncoder(w).Encode.

5. Implement RunReport(r Reporter, w io.Writer) error:
   • Calls r.Report(w)
   • defer logs the time taken to produce the report
   • Returns any error from Report
*/

type Reporter interface {
	Report(w io.Writer) error
}

func DealerReporterFormatter(d *Dealer, DealerAnalysis *DealerAnalysis) {
	header := "REPORT: AutoDealer Pro  (Kyiv, UA)"
	totalCars := len(d.Inventory)
	availableCars := d.FilterByStatus(StatusAvailable)
	soldCars := d.FilterByStatus(StatusSold)
	inventoryValueUAH, inventoryValueUSD := func(d *Dealer) (int, int) {
		var sumUAH int
		var sumUSD int
		for _, car := range d.Inventory {
			sumUAH = sumUAH + int(car.PriceUAH)
			sumUSD = sumUSD + int(car.PriceUSD)
		}
		return sumUAH, sumUSD
	}(d)
	avgCarPrice := func(d *Dealer) int {
		var sum int
		for _, car := range d.Inventory {
			sum = sum + int(car.PriceUAH)
		}
		return sum / len(d.Inventory)
	}(d)
	mostExpensiveCarName, mostExpensiveCarPrice := func(d *Dealer) (string, float64) {
		mostExpensivePrice := 0.0
		var mostExpensivePriceName string
		for _, car := range d.Inventory {
			if car.PriceUAH > float64(mostExpensivePrice) {
				mostExpensivePrice = car.PriceUAH
				mostExpensivePriceName = car.Model
			}
		}
		return mostExpensivePriceName, mostExpensivePrice
	}(d)
	numEmployees := len(d.Employees)
	monthlyPayroll := func(d *Dealer) int {
		var sum int
		for _, emp := range d.Employees {
			sum = sum + int(emp.Salary)
		}
		return sum
	}(d)
	empByRole := d.CalcEmployeesByRole()
	totalDeals := len(d.Deals)
	revenue := Sum(func(d *Dealer) []float64 {
		dealRevenues := []float64{}
		for _, deal := range d.Deals {
			dealRevenues = append(dealRevenues, deal.SalePriceUAH)
		}
		return dealRevenues
	}(d))
	avgDealPrice := revenue / len(d.Deals)
	topSellerRevenue, topSellerEmp := d.CalcTopSeller()
	dealsByType := d.CalcDealsByPayment()
	categoriesByRevenue := d.CalcRevenueByCategory()
	fmt.Printf(`
%s
----------------------------------------
Cars:
  Total cars: %d
  Available cars: %d
  Sold cars: %d
  Inventory value: %d UAH | %d USD
  Average car price (UAH): %d
  Most expensive car: %s (%.2f UAH)

Employees:
  Total employees: %d
  Monthly payroll: %d
  Employees by role: %v

Deals:
  Total deals: %d
  Total revenue (UAH): %.2f
  Average deal price (UAH): %.2f
  Top seller: %s (%.2f UAH)
  Deals by payment type: %v
  Revenue by category: %v
----------------------------------------
`,
		header,
		totalCars,
		len(availableCars),
		len(soldCars),
		inventoryValueUAH,
		inventoryValueUSD,
		avgCarPrice,
		mostExpensiveCarName,
		mostExpensiveCarPrice,
		numEmployees,
		monthlyPayroll,
		empByRole,
		totalDeals,
		revenue,
		avgDealPrice,
		topSellerEmp.FullName(),
		topSellerRevenue,
		dealsByType,
		categoriesByRevenue,
	)

}
