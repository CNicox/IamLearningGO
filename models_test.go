package main

import "testing"

func TestAdd(t *testing.T) {
	employee := Employee{
		ID:         "emp_001",
		FirstName:  "Olena",
		LastName:   "Kovalenko",
		Role:       RoleSalesManager,
		HireDate:   "2022-05-15",
		Salary:     45000.50,
		BonusPct:   10.0,
		Department: "Sales",
		ContactInfo: ContactInfo{
			Phone: "+380501234567",
			Email: "olena.kovalenko@company.com",
		},
	}

	if "+380501234567" != employee.ContactInfo.Phone {
		t.Fatal("expected +380501234567, got %d", employee.ContactInfo.Phone)
	}
	if "olena.kovalenko@company.com" != employee.ContactInfo.Email {
		t.Fatal("expected olena.kovalenko@company.com, got %d", employee.ContactInfo.Email)
	}

}

func TestDealerDealerInfo(t *testing.T) {
	dealer := Dealer{
		DealerInfo: DealerInfo{
			ID:          "dealer-001",
			Name:        "AutoGalaxy",
			City:        "Kyiv",
			Address:     "123 Main St",
			Phone:       "+380501234567",
			FoundedYear: 2005,
		},
		Inventory: []Car{
			{
				VIN:          "WVWZZZ1JZXW000001",
				Make:         "Volkswagen",
				Model:        "Golf",
				Year:         2022,
				Color:        "Red",
				PriceUAH:     800000,
				PriceUSD:     20000,
				Mileage:      0,
				Status:       StatusAvailable,
				Category:     "Hatchback",
				FuelType:     "Petrol",
				Transmission: "Manual",
				Options: &CarOptions{
					Sunroof:     true,
					HeatedSeats: true,
				},
				DiscountPct: 5.0,
			},
			{
				VIN:          "1HGCM82633A004352",
				Make:         "Honda",
				Model:        "Accord",
				Year:         2021,
				Color:        "Blue",
				PriceUAH:     950000,
				PriceUSD:     24000,
				Mileage:      15000,
				Status:       StatusSold,
				Category:     "Sedan",
				FuelType:     "Hybrid",
				Transmission: "Automatic",
			},
		},
		Employees: []Employee{
			{
				ID:         "emp-001",
				FirstName:  "Olena",
				LastName:   "Kovalenko",
				Role:       RoleSalesManager,
				HireDate:   "2022-05-15",
				Salary:     45000,
				BonusPct:   10.0,
				Department: "Sales",
				ContactInfo: ContactInfo{
					Phone: "+380501112233",
					Email: "olena.kovalenko@autogalaxy.com",
				},
			},
			{
				ID:         "emp-002",
				FirstName:  "Ivan",
				LastName:   "Shevchenko",
				Role:       RoleTechnician,
				HireDate:   "2021-03-10",
				Salary:     35000,
				Department: "Service",
				ContactInfo: ContactInfo{
					Phone: "+380501223344",
					Email: "ivan.shevchenko@autogalaxy.com",
				},
			},
		},
		Deals: []Deal{
			{
				ID:           "deal-001",
				VIN:          "WVWZZZ1JZXW000001",
				EmployeeID:   "emp-001",
				ClientName:   "John Doe",
				Date:         "2025-02-01",
				SalePriceUAH: 760000,
				SalePriceUSD: 19000,
				PaymentType:  PaymentCash,
				Notes:        "First-time customer discount applied",
			},
			{
				ID:           "deal-002",
				VIN:          "1HGCM82633A004352",
				EmployeeID:   "emp-001",
				ClientName:   "Maria Ivanova",
				Date:         "2025-02-15",
				SalePriceUAH: 950000,
				SalePriceUSD: 24000,
				PaymentType:  PaymentCredit,
			},
		},
	}
	dealerName := dealer.Name
	if dealerName != dealer.Name {
		t.Fatal("AutoGalaxy, got %d", dealer.Name)
	}
}
