package main

import (
	"testing"
)

func TestCountMaxOccurences(t *testing.T) {
	items := []string{"Alice", "Bob", "Alice", "Charlie", "Bob", "Alice"}

	count, value := countMaxOccurences(items)

	if count != 3 {
		t.Errorf("expected count 3, got %d", count)
	}
	if value != "Alice" {
		t.Errorf("expected value Alice, got %s", value)
	}
}

func TestCalcTopSeller(t *testing.T) {
	emp1 := Employee{ID: "e1", FirstName: "Alice"}
	emp2 := Employee{ID: "e2", FirstName: "Bob"}
	emp3 := Employee{ID: "e3", FirstName: "Charlie"}

	deals := []Deal{
		{EmployeeID: "e1", ClientName: "Client1"},
		{EmployeeID: "e2", ClientName: "Client2"},
		{EmployeeID: "e1", ClientName: "Client3"},
		{EmployeeID: "e3", ClientName: "Client4"},
		{EmployeeID: "e1", ClientName: "Client5"},
	}

	dealer := Dealer{
		Employees: []Employee{emp1, emp2, emp3},
		Deals:     deals,
	}

	count, top := dealer.CalcTopSeller()

	if count != 3 {
		t.Errorf("expected top seller count 3, got %d", count)
	}
	if top.ID != "e1" {
		t.Errorf("expected top seller ID e1, got %s", top.ID)
	}
	if top.FirstName != "Alice" {
		t.Errorf("expected top seller Alice, got %s", top.FirstName)
	}
}

func TestCalcRevenueByCategory(t *testing.T) {
	dealer := Dealer{
		Inventory: []Car{
			{PriceUAH: 1000, Status: StatusAvailable},
			{PriceUAH: 1500, Status: StatusAvailable},
			{PriceUAH: 2000, Status: StatusAvailable},
			{PriceUAH: 500, Status: StatusReserved},
			{PriceUAH: 600, Status: StatusReserved},
			{PriceUAH: 700, Status: StatusReserved},
			{PriceUAH: 800, Status: StatusReserved},
			{PriceUAH: 1200, Status: StatusAvailable},
			{PriceUAH: 1300, Status: StatusAvailable},
			{PriceUAH: 900, Status: StatusReserved},
		},
	}

	result := dealer.CalcRevenueByCategory()

	if result[string(StatusAvailable)].AvgUAH != 1400 {
		t.Errorf("expected available average 1400, got %v", result["new"].AvgUAH)
	}
	if result[string(StatusReserved)].AvgUAH != 700 {
		t.Errorf("expected reserved average 650, got %v", result["used"].AvgUAH)
	}
}

func TestCalcEmployeesByRole(t *testing.T) {
	dealer := Dealer{
		Employees: []Employee{
			{ID: "1", FirstName: "Alice", Role: "sales_manager"},
			{ID: "2", FirstName: "Bob", Role: "technician"},
			{ID: "3", FirstName: "Charlie", Role: "sales_manager"},
			{ID: "4", FirstName: "Diana", Role: "receptionist"},
			{ID: "5", FirstName: "Eve", Role: "technician"},
			{ID: "6", FirstName: "Frank", Role: "director"},
		},
	}

	result := dealer.CalcEmployeesByRole()

	// Check sales_manager
	if len(result["sales_manager"]) != 2 {
		t.Errorf("expected 2 sales managers, got %d", len(result["sales_manager"]))
	}

	// Check technician
	if len(result["technician"]) != 2 {
		t.Errorf("expected 2 technicians, got %d", len(result["technician"]))
	}

	// Check receptionist
	if len(result["receptionist"]) != 1 {
		t.Errorf("expected 1 receptionist, got %d", len(result["receptionist"]))
	}

	// Check director
	if len(result["director"]) != 1 {
		t.Errorf("expected 1 director, got %d", len(result["director"]))
	}

	// Optional: check specific employee IDs
	if result["sales_manager"][0].ID != "1" || result["sales_manager"][1].ID != "3" {
		t.Errorf("sales managers IDs mismatch: %+v", result["sales_manager"])
	}
}

func TestCalcDealsByPayment(t *testing.T) {
	dealer := Dealer{
		Deals: []Deal{
			{ID: "1", PaymentType: PaymentCash},
			{ID: "2", PaymentType: PaymentCash},
			{ID: "3", PaymentType: PaymentCredit},
			{ID: "4", PaymentType: PaymentLeasing},
			{ID: "5", PaymentType: PaymentLeasing},
			{ID: "6", PaymentType: PaymentLeasing},
		},
	}

	result := dealer.CalcDealsByPayment()

	if result[PaymentCash] != 2 {
		t.Errorf("expected 2 cash deals, got %d", result[PaymentCash])
	}

	if result[PaymentCredit] != 1 {
		t.Errorf("expected 1 credit deal, got %d", result[PaymentCredit])
	}

	if result[PaymentLeasing] != 3 {
		t.Errorf("expected 3 leasing deals, got %d", result[PaymentLeasing])
	}
}

func TestCalcTotalPayroll(t *testing.T) {
	dealer := Dealer{
		Employees: []Employee{
			{ID: "1", Salary: 10000, BonusPct: 10}, // 11000
			{ID: "2", Salary: 20000, BonusPct: 0},  // 20000
			{ID: "3", Salary: 15000, BonusPct: 20}, // 18000
			{ID: "4", Salary: 8000, BonusPct: 5},   // 8400
		},
	}

	result := dealer.CalcTotalPayroll()

	expected := 11000.0 + 20000.0 + 18000.0 + 8400.0 // 57400

	if result != expected {
		t.Errorf("CalcTotalPayroll: got %.2f, want %.2f", result, expected)
	}
}
