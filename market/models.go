package market

// Тип сортировки
const (
	ASC sortDirection = iota
	DESC
	MANY
)

//  Статус объекта
const (
	NORMAL = iota
	PREMIUM
)

// Types
type sortDirection int
type typeStatus int

// Товар пробник
type Probe string

// Structs
type Shop struct {
	Products map[string]Product
	Kits     map[string]Kit
	Probes   []Probe
	Users    map[string]User
	Cache    map[string]float32
}

type pair struct {
	key   string
	value float32
}

type Product struct {
	Name   string
	Price  float32
	Status typeStatus
}

type Kit struct {
	MainProduct       string
	AdditionalProduct string
	Probe             Probe
	Discount          float32
}

type Order struct {
	User     string
	Products []string
	Kits     []string
}

type CacheInfo struct {
	IsHas bool
	Total float32
	Key   string
}

type User struct {
	Bill   float32
	Status typeStatus
}
