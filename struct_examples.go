package main

func NewTestDealer() *Dealer {
	return &Dealer{
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
}
func NewTestDealerReporter() *DealerReporter {
	dealer := NewTestDealer()

	analysis := &DealerAnalysis{
		// Inventory
		TotalCars:          len(dealer.Inventory),
		AvailableCars:      1,
		SoldCars:           1,
		ReservedCars:       0,
		TotalInventoryUAH:  1750000,
		TotalInventoryUSD:  44000,
		AverageCarPriceUAH: 875000,
		MostExpensiveCar:   &dealer.Inventory[1],
		CheapestCar:        &dealer.Inventory[0],

		TotalEmployees:  2,
		TotalPayrollUAH: 80000,
		EmployeesByRole: map[EmployeeRole][]Employee{
			RoleSalesManager: {dealer.Employees[0]},
			RoleTechnician:   {dealer.Employees[1]},
		},

		TotalDeals:      2,
		TotalRevenueUAH: 1710000,
		TotalRevenueUSD: 43000,
		AverageDealUAH:  855000,
		TopSeller:       &dealer.Employees[0],
		DealsByPayment: map[PaymentType]int{
			PaymentCash:   1,
			PaymentCredit: 1,
		},

		CategoryStats: map[string]CategoryStat{
			"Sedan":     {Count: 1},
			"Hatchback": {Count: 1},
		},
	}

	return &DealerReporter{
		Dealer:   dealer,
		Analysis: analysis,
	}
}
