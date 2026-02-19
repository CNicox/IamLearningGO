
// this is how you iterate over a structs keys and values
func (CarOptions CarOptions) CountOptions() {
	ObjValue := reflect.ValueOf(CarOptions)
	ObjType := reflect.TypeOf(CarOptions)
	for i := 0; i < ObjValue.NumField(); i++ {
		// field key
		field := ObjType.Field(i)
		// field value
		value := ObjValue.Field(i).Interface()
		fmt.Printf("%s: %v\n", field, value)
	}
}


	car := Car{
		VIN:      vin,
		Make:     make,
		Year:     year,
		PriceUAH: priceUAH,
		PriceUSD: priceUSD,
		Mileage:  0,
		Status:   status,
		Options:  nil,
	}