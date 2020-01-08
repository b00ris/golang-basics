package market

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const separator = string(filepath.Separator)
const pathToMemory = ".." + separator + "memory" + separator

// Экспорт магазина в JSON файл в указанный каталог
func ExportShop(shop Shop, dirname string) {
	os.Mkdir(pathToMemory+dirname, 0777)
	//Сохраняем магазин
	file, err := os.Create(pathToMemory + dirname + separator + "Shop.json")
	if err != nil {
		panic(err)
	}
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(shop); err != nil {
		panic(err)
	}
	file.Close()
}

//Импорт магазина из JSON файла
func ImportShop(dirname string) (shop Shop) {
	//Получаем магазин
	data, err := ioutil.ReadFile(pathToMemory + dirname + separator + "Shop.json")
	if err != nil {
		fmt.Print(err)
	}
	err = json.Unmarshal(data, &shop)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return shop
}
