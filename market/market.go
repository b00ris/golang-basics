package market

var Users = map[string]float32{}
var Products = map[string]float32{}

type Order struct {
	User     string
	Products []string
}

func CalcOrder(order Order) float32 {
	var total float32
	for _, product := range order.Products {
		total += Products[product]
	}
	value, okay := Users[order.User]
	if okay == true && value >= total {
		Users[order.User] -= total
	}
	return total
}

func CalcOrderWithSaveSpeed(order Order) float32 {
	var total float32
	for _, product := range order.Products {
		total += Products[product]
	}
	value, okay := Users[order.User]
	if okay == true && value >= total {
		Users[order.User] -= total
	}
	return total
}

func addProduct(name string, price float32) {
	_, okay := Products[name]
	if okay == false {
		Products[name] = price
	}
}

func updateProduct(name string, price float32) {
	_, okay := Products[name]
	if okay == true {
		Products[name] = price
	}
}
