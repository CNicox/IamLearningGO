package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Validatable interface {
	Validate() []error
}

// its not the same implementation as in the document but i kinda like it
func ValidateAll[T Validatable](items []T) []error {
	var errs []error
	for _, item := range items {
		errs = append(errs, item.Validate()...)
	}
	return errs
}

func (Car Car) Validate() []error {
	var errs []error
	if len(Car.VIN) != 17 {
		errs = append(errs, errors.New("wrong vin format"))
	}
	if Car.Year <= 1990 && Car.Year >= CURRENT_YEAR+1 {
		errs = append(errs, errors.New("year is invalid"))
	}
	if Car.PriceUAH <= 0 || Car.PriceUSD <= 0 {
		errs = append(errs, errors.New("price is invalid"))
	}
	if IsInStatuses(allCarStatuses, Car.Status) == false {
		errs = append(errs, errors.New("invalid status"))
	}
	return errs
}

func (e Employee) Validate() []error {
	var errs []error
	isRoleAllowed := IsInRoles(allRoles, e.Role)
	if e.ID == "" {
		errs = append(errs, errors.New("id is required"))
	}
	if e.Salary <= 0 {
		errs = append(errs, errors.New("salary must be greater than 0"))
	}
	if !isRoleAllowed {
		errs = append(errs, errors.New("role is not allowed"))
	}
	return errs
}

func (Deal Deal) Validate() []error {
	var errs []error
	_, err := time.Parse("2006-01-02", Deal.Date)
	if err != nil {
		errs = append(errs, errors.New("wrong date format"))
	}
	if Deal.SalePriceUAH <= 0 || Deal.SalePriceUSD <= 0 {
		errs = append(errs, errors.New("price is invalid"))
	}
	return errs
}

func (Dealer Dealer) Validate() []error {
	var errs []error
	if (Dealer.Name) != "" {
		errs = append(errs, errors.New("Name is required"))
	}
	if len(Dealer.Inventory) <= 0 {
		errs = append(errs, errors.New("A car is required"))
	}
	if len(Dealer.Employees) <= 0 {
		errs = append(errs, errors.New("An employees is required"))
	}
	if len(Dealer.Deals) <= 0 {
		errs = append(errs, errors.New("A deal is required"))
	}
	errs = append(errs, ValidateAll(Dealer.Inventory)...)
	return errs
}

func (CarOptions CarOptions) String() string {
	ObjValue := reflect.ValueOf(CarOptions)
	ObjType := reflect.TypeOf(CarOptions)
	var activeOptions []string
	for i := 0; i < ObjValue.NumField(); i++ {
		value := ObjValue.Field(i)
		if value.Bool() {
			activeOptions = append(activeOptions, ObjType.Field(i).Name)
		}
	}
	if len(activeOptions) == 0 {
		return "(no options)."
	}
	return strings.Join(activeOptions, ",")
}

// I don't understand how to work with a pointer value of CarOptions here
func (CarOptions CarOptions) CountOptions() int {
	counter := 0
	ObjValue := reflect.ValueOf(CarOptions)
	for i := 0; i < ObjValue.NumField(); i++ {
		value := ObjValue.Field(i)
		if value.Bool() {
			counter++
		}
	}
	return counter
}

func IsInStatuses(statuses []CarStatus, s CarStatus) bool {
	for _, status := range statuses {
		if status == s {
			return true
		}
	}
	return false
}

func NewCar(vin, make, model string, year int, priceUAH, priceUSD float64, status CarStatus, DiscountPct float64) *Car {
	car := Car{
		VIN:         vin,
		Make:        make,
		Year:        year,
		PriceUAH:    priceUAH,
		PriceUSD:    priceUSD,
		Mileage:     0,
		Status:      status,
		Options:     nil,
		DiscountPct: DiscountPct,
	}
	return &car
}

func (Car Car) FinalPriceUAH() float64 {
	return Car.PriceUAH * (1 - Car.DiscountPct/100)
}

func (Car Car) FinalPriceUSD() float64 {
	return Car.PriceUSD * (1 - Car.DiscountPct/100)
}

func (Car Car) IsAvailable() bool {
	if Car.Status == StatusAvailable {
		return true
	}
	return false
}

func (Car Car) HasOptions() bool {
	if Car.Options != nil && Car.Options.CountOptions() > 0 {
		return true
	}
	return false
}

func (Car Car) String() string {
	ObjValue := reflect.ValueOf(Car)
	ObjType := reflect.TypeOf(Car)
	var activeOptions []string
	for i := 0; i < ObjValue.NumField(); i++ {
		activeOptions = append(activeOptions, ObjType.Field(i).Name)
	}
	return strings.Join(activeOptions, " ")
}

func IsInRoles(roles []EmployeeRole, s EmployeeRole) bool {
	for _, role := range roles {
		if role == s {
			return true
		}
	}
	return false
}

func NewEmployee(id, first, last string, role EmployeeRole, salary float64) *Employee {
	employee := Employee{
		ID:        id,
		FirstName: first,
		LastName:  last,
		Role:      role,
		Salary:    salary,
	}
	return &employee
}

func (Employee Employee) FullName() string {
	return Employee.FirstName + Employee.LastName
}

func (Employee Employee) TotalCost() float64 {
	return Employee.Salary * (1 + Employee.BonusPct/100)
}

func (Employee Employee) String() string {
	return Employee.FullName() + string(Employee.Role) + "|" + "Salary" + strconv.FormatFloat(Employee.Salary, 'f', 2, 64) + "₴" + "|" + "ID:" + Employee.ID
}

func NewDeal(id, vin, empID, clientName, date string,
	priceUAH, priceUSD float64, pt PaymentType) *Deal {
	deal := Deal{
		ID:           id,
		VIN:          vin,
		EmployeeID:   empID,
		ClientName:   clientName,
		Date:         date,
		SalePriceUAH: priceUAH,
		SalePriceUSD: priceUSD,
		PaymentType:  pt,
	}
	return &deal
}

func (Deal Deal) String() string {
	return "DEAL-" + Deal.ID + Deal.VIN + " | " + Deal.ClientName + " | " + strconv.FormatFloat(Deal.SalePriceUAH, 'f', 2, 64) + " ₴ " + " | " + string(Deal.PaymentType)
}

func (Deal Deal) HasNotes() bool {
	if len(Deal.Notes) > 0 {
		return true
	}
	return false
}

func (d Dealer) String() string {
	return fmt.Sprintf("%s | cars:%d | %d | %d", d.Name, len(d.Inventory), len(d.Employees), len(d.Deals))
}

func CarOptionsFunctionalityCheck() {
	CarOptionsObj := CarOptions{
		Sunroof:     true,
		HeatedSeats: true,
	}
	CarOptionsObj_2 := CarOptions{
		Sunroof:     false,
		HeatedSeats: false,
	}
	howManyTrue := CarOptionsObj.CountOptions()
	fmt.Println(howManyTrue)
	whichOnesTrue := CarOptionsObj.String()
	fmt.Println(whichOnesTrue)
	fmt.Println(CarOptionsObj_2.String())
}

func CarFunctionalityCheck() {
	NewCar_1 := NewCar("12345678912345678", "make", "Ford", 1999, 5000, 100, "available", 10.0)
	fmt.Println(NewCar_1.FinalPriceUAH(), NewCar_1.FinalPriceUSD(), NewCar_1.IsAvailable(), NewCar_1.String())
	NewCar_2 := NewCar("1234567891234567", "make", "Ford", 1800, 5000, 100, "aveilable", 10.0)
	fmt.Println(NewCar_2.PriceUAH)
	fmt.Println(len(NewCar_1.Validate()))
	fmt.Println(len(NewCar_2.Validate()))
}

func EmployeeFunctionalityCheck() {
	employee_1 := NewEmployee("123", "Mykola", "Cornienko", "director", 10000.0)
	employee_1.TotalCost()
	fmt.Println(employee_1.String())
	fmt.Println(len(employee_1.Validate()), employee_1.FullName(), employee_1.TotalCost())
}

func DealFunctionalityCheck() {
	deal := NewDeal(
		"001",
		"WVWZZZ1JZXW000001",
		"emp-123",
		"John Doe",
		"2025-02-01",
		2500000.00,
		65000.00,
		PaymentCash,
	)
	fmt.Println(deal.String())
	fmt.Println(deal.HasNotes())
	fmt.Println(len(deal.Validate()))
}

func DealerFunctionalityCheck() {
	dealer := NewTestDealer()
	fmt.Println(dealer.String())
	fmt.Println(len(dealer.Validate()))
}

func PrintAll[T fmt.Stringer](items []T, w io.Writer) {
	for _, item := range items {
		fmt.Fprintln(w, item)
	}
}

func PrintAllFunctionalityCheck() {
	car1 := Car{
		VIN:          "WBA3A5C57CF123456",
		Make:         "BMW",
		Model:        "320i",
		Year:         2018,
		Color:        "Black",
		PriceUAH:     850000,
		PriceUSD:     22500,
		Mileage:      72000,
		Status:       StatusAvailable,
		Category:     "Sedan",
		FuelType:     "Petrol",
		Transmission: "Automatic",
		Options:      nil,
		DiscountPct:  5,
	}

	car2 := Car{
		VIN:          "1HGCV1F14JA765432",
		Make:         "Honda",
		Model:        "Accord",
		Year:         2020,
		Color:        "White",
		PriceUAH:     920000,
		PriceUSD:     24500,
		Mileage:      45000,
		Status:       StatusReserved,
		Category:     "Sedan",
		FuelType:     "Petrol",
		Transmission: "CVT",
		Options:      nil,
	}

	car3 := Car{
		VIN:          "JTMEB3FV60D098765",
		Make:         "Toyota",
		Model:        "RAV4 Hybrid",
		Year:         2022,
		Color:        "Silver",
		PriceUAH:     1350000,
		PriceUSD:     36000,
		Mileage:      18000,
		Status:       StatusSold,
		Category:     "SUV",
		FuelType:     "Hybrid",
		Transmission: "Automatic",
		Options:      nil,
	}
	cars := []Car{car1, car2, car3}
	cars_available := GetAvailableCars(cars)
	PrintAll(cars_available, os.Stdout)
}

func GetAvailableCars(cars []Car) []Car {
	cars_available := []Car{}
	for _, car := range cars {
		if car.Status == StatusAvailable {
			cars_available = append(cars_available, car)
		}
	}
	return cars_available
}

type Number interface {
	~int | ~int64 | ~float32 | ~float64
}

func Sum[T Number](nums []T) int {
	var sum int
	for _, n := range nums {
		sum += int(n)
	}
	return sum
}
