package tests

import (
	"lessons/basics/market"
	"math/rand"
	"testing"
)

func TestCalcOrder(t *testing.T) {
	shop := market.ShopInit(InitProducts(), InitUsers())
	//products := InitProducts()
	//users := InitUsers()
	orders := InitOrder(shop)

	total := shop.CalcOrder(orders[1], false)

	//jekamas: не нужно создавать такое же константы в тесте. если они будут меняться, то придется менять их в двух местах. так ты увеличишь объем работы на поддержке кода.
	shop.PrintUsers(market.DESC)

	//t.Log(price)
	//t.Log(users["Gates"])

	// jekamas: нужно добавить проверку, что собственно мы ожидаем и получили ли мы ожидаемое
	var checkTotal float32

	for _, productName := range orders[1].Products {
		product := shop.Products[productName]
		t.Log(checkTotal, " + ", product.Price)
		checkTotal += product.Price
	}
	if checkTotal != total {
		t.Error("Итоговая сумма неверная")
	}
}

func TestStatusProducts(t *testing.T) {
	products := map[string]market.Product{
		"Pineapple": {"Pineapple", 500.0, market.PREMIUM},
		"Discount":  {"Discount", 500.0, market.NORMAL},
	}

	users := map[string]market.User{
		"Kevin": {1_000, market.PREMIUM},
		"Bred":  {1_000, market.NORMAL},
	}

	shop := market.ShopInit(products, users)

	kevinOrder := shop.CalcOrder(market.Order{
		User:     "Kevin",
		Products: []string{"Pineapple", "Discount"}}, true)
	bredOrder := shop.CalcOrder(market.Order{
		User:     "Bred",
		Products: []string{"Pineapple", "Discount"}}, true)

	shop.PrintUsers(market.ASC)
	if kevinOrder >= bredOrder {
		t.Fatal("Премиум пользователь не может платить больше обычного")
	}

	market.ExportShop(*shop, "TestStatus")
}

func TestKit(t *testing.T) {
	products := map[string]market.Product{
		"Pineapple": {"Pineapple", 500.0, market.NORMAL},
		"Cola":      {"Cola", 5.0, market.NORMAL},
	}

	users := map[string]market.User{
		"Kevin": {1_000, market.NORMAL},
	}
	probes := []market.Probe{
		"Cheesecake",
	}
	kits := map[string]market.Kit{
		"Pinakalada": market.CreateKit("Pineapple", "Cola", probes[0], 50),
	}
	shop := market.Shop{
		Products: products,
		Kits:     kits,
		Probes:   probes,
		Users:    users,
		Cache:    map[string]float32{},
	}

	total := shop.CalcOrder(market.Order{
		User: "Kevin",
		//Products: 	[]string {"Pineapple", "Discount"},
		Kits: []string{"Pinakalada"},
	}, true)

	market.ExportShop(shop, "TestKit")
	market.ImportShop("TestKit")

	t.Log(total)
}
func TestMemory(t *testing.T) {
	products := map[string]market.Product{
		"Pineapple": {"Pineapple", 500.0, market.NORMAL},
		"Cola":      {"Cola", 5.0, market.NORMAL},
	}

	users := map[string]market.User{
		"Kevin": {1_000, market.NORMAL},
	}
	probes := []market.Probe{
		"Cheesecake",
	}
	kits := map[string]market.Kit{
		"Pinakalada": market.CreateKit("Pineapple", "Cola", probes[0], 50),
	}
	shop := market.Shop{
		Products: products,
		Kits:     kits,
		Probes:   probes,
		Users:    users,
		Cache:    map[string]float32{},
	}

	shop.CalcOrder(market.Order{
		User: "Kevin",
		//Products: 	[]string {"Pineapple", "Discount"},
		Kits: []string{"Pinakalada"},
	}, true)

	market.ExportShop(shop, "TestMemory")
	shop2 := market.ImportShop("TestMemory")

	for name, user := range shop.Users {
		if user2, okay := shop2.Users[name]; !okay || user2 != user {
			t.Fatal("Пользователи не соответствуют ожидаемым")
		}
	}

	for name, product := range shop.Products {
		if product2, okay := shop2.Products[name]; !okay || product2 != product {
			t.Fatal("Товары не соответствуют ожидаемым")
		}
	}

	for name, kit := range shop.Kits {
		if kit2, okay := shop2.Kits[name]; !okay || kit2 != kit {
			t.Fatal("Наборы не соответствуют ожидаемым")
		}
	}

	for i, probe := range shop.Probes {
		if probe2 := shop2.Probes[i]; probe2 != probe {
			t.Fatal("Пробник не соответствуют ожидаемым")
		}
	}

	for name, cache := range shop.Cache {
		if cache2, okay := shop2.Cache[name]; !okay || cache2 != cache {
			t.Fatal("Кэш не соответствуют ожидаемому")
		}
	}
}
func BenchmarkCalcOrder(b *testing.B) {
	//products := InitProducts()
	//users := InitUsers()
	shop := market.ShopInit(InitProducts(), InitUsers())

	orders := InitOrder(shop)
	randIndex := make([]int, b.N)
	for i := range randIndex {
		randIndex[i] = rand.Intn(100)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		//jekamas:  тут есть проблема. ты заодно будешь измерять производительность rand.Intn(100). лучше вне этого цикла и до вызова .ResetTimer() сделать массив случайных значений нужной длины и после его использовать, чтобы мы мерили только нужное.
		shop.CalcOrder(orders[randIndex[i]], false)
	}

}

func BenchmarkCalcOrderSpeedCache(b *testing.B) {
	//products := InitProducts()
	//users := InitUsers()

	shop := market.ShopInit(InitProducts(), InitUsers())
	orders := InitOrder(shop)
	randIndex := make([]int, b.N)
	for i := range randIndex {
		randIndex[i] = rand.Intn(100)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		shop.CalcOrder(orders[randIndex[i]], true)
	}
	//market.ExportShop(*shop, "Benchmark")
}

func InitUsers() map[string]market.User {
	return map[string]market.User{
		"Kevin": {1_500_000_000, market.PREMIUM},
		"Gates": {1_000_000_000, market.NORMAL},
		"Ford":  {500_000_000, market.NORMAL},
	}
}

func InitProducts() map[string]market.Product {
	return map[string]market.Product{
		"Pineapple": {"Pineapple", 15.0, market.NORMAL},
		"Discount":  {"Discount", 50.0, market.NORMAL},
		"Meat":      {"Meat", 125.0, market.NORMAL},
		"Whisky":    {"Whisky", 5_000.0, market.PREMIUM},
		"Chocolate": {"Chocolate", 35.5, market.PREMIUM},
		"Banana":    {"Banana", 25.25, market.PREMIUM},
		"Mango":     {"Mango", 10.0, market.NORMAL},
		"Marakua":   {"Marakua", 123.0, market.PREMIUM},
		"Potato":    {"Potato", 2.0, market.NORMAL},
		"Tea":       {"Tea", 10.0, market.NORMAL},
		"Coffee":    {"Coffee", 100.0, market.PREMIUM},
	}
}
func InitOrder(shop *market.Shop) []market.Order {
	keyUsers := make([]string, len(shop.Users))
	i := 0
	for user := range shop.Users {
		keyUsers[i] = user
		i++
	}

	keyProducts := make([]string, len(shop.Products))
	i = 0
	for product := range shop.Products {
		keyProducts[i] = product
		i++
	}
	orders := make([]market.Order, 100)
	for i := 0; i < 100; i++ {
		orders[i] = getOrder(keyUsers, keyProducts)
	}

	return orders
}

func getOrder(userList []string, productList []string) market.Order {
	length := rand.Intn(100)
	products := make([]string, length)

	for i := 0; i < length; i++ {
		products[i] = productList[rand.Intn(len(productList))]
	}

	return market.Order{
		User:     userList[rand.Intn(len(userList))],
		Products: products,
	}
}
