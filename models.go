package main

import (
	"time"
)

var CURRENT_YEAR int = time.Now().Year()

type CarOptions struct {
	Sunroof        bool `json:"sunroof,omitempty"`
	HeatedSeats    bool `json:"heated_seats,omitempty"`
	Navigation     bool `json:"navigation,omitempty"`
	ParkingSensors bool `json:"parking_sensors,omitempty"`
	LeatherSeats   bool `json:"leather_seats,omitempty"`
	CruiseControl  bool `json:"cruise_control,omitempty"`
}

type CarStatus string

const (
	StatusAvailable CarStatus = "available"
	StatusSold      CarStatus = "sold"
	StatusReserved  CarStatus = "reserved"
)

var allCarStatuses = []CarStatus{
	StatusAvailable,
	StatusSold,
	StatusReserved,
}

type Car struct {
	VIN          string      `json:"vin"`
	Make         string      `json:"make"`
	Model        string      `json:"model"`
	Year         int         `json:"year"`
	Color        string      `json:"color"`
	PriceUAH     float64     `json:"price_uah"`
	PriceUSD     float64     `json:"price_usd"`
	Mileage      int         `json:"mileage"`
	Status       CarStatus   `json:"status"`
	Category     string      `json:"category"`
	FuelType     string      `json:"fuel_type"`
	Transmission string      `json:"transmission"`
	Options      *CarOptions `json:"options,omitempty"`
	DiscountPct  float64     `json:"discount_pct,omitempty"`
}

type EmployeeRole string

const (
	RoleSalesManager EmployeeRole = "sales_manager"
	RoleTechnician   EmployeeRole = "technician"
	RoleReceptionist EmployeeRole = "receptionist"
	RoleDirector     EmployeeRole = "director"
)

var allRoles = []EmployeeRole{
	RoleSalesManager,
	RoleTechnician,
	RoleReceptionist,
	RoleDirector,
}

type ContactInfo struct {
	Phone string `json:"phone"`
	Email string `json:"email,omitempty"`
}

type Employee struct {
	ID          string       `json:"id"`
	FirstName   string       `json:"first_name"`
	LastName    string       `json:"last_name"`
	Role        EmployeeRole `json:"role"`
	HireDate    string       `json:"hire_date"` // format "YYYY-MM-DD"
	Salary      float64      `json:"salary"`    // monthly, in UAH
	BonusPct    float64      `json:"bonus_pct,omitempty"`
	Department  string       `json:"department"`
	ContactInfo              // embedded — Phone and Email promoted to Employee
}

type PaymentType string

const (
	PaymentCash    PaymentType = "cash"
	PaymentCredit  PaymentType = "credit"
	PaymentLeasing PaymentType = "leasing"
)

type Deal struct {
	ID           string      `json:"id"`
	VIN          string      `json:"vin"`
	EmployeeID   string      `json:"employee_id"`
	ClientName   string      `json:"client_name"`
	Date         string      `json:"date"`
	SalePriceUAH float64     `json:"sale_price_uah"`
	SalePriceUSD float64     `json:"sale_price_usd"`
	PaymentType  PaymentType `json:"payment_type"`
	Notes        string      `json:"notes,omitempty"` // optional field
}

type DealerInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Address     string `json:"address,omitempty"`
	Phone       string `json:"phone,omitempty"`
	FoundedYear int    `json:"founded_year,omitempty"`
}

type Dealer struct {
	DealerInfo            // embedding: Dealer.Name, Dealer.City accessible directly
	Inventory  []Car      `json:"inventory"`
	Employees  []Employee `json:"employees"`
	Deals      []Deal     `json:"deals"`
}

type Root struct {
	Dealer Dealer `json:"dealer"`
}

// DealerReporter implements Reporter for a Dealer.
type DealerReporter struct {
	Dealer   *Dealer
	Analysis *DealerAnalysis // result from analyzer.go
}

func main() {
	// CarOptionsFunctionalityCheck()
	// CarFunctionalityCheck()
	// EmployeeFunctionalityCheck()
	// DealFunctionalityCheck()
	// DealerFunctionalityCheck()

}
