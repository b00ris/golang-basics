package shop

import (
	"bytes"
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/youricorocks/shop_competition"
	"golang.org/x/sync/errgroup"
	"io"
	"strconv"
	"sync"
	"time"
)

var (
	TimeoutError = errors.New("timeout error")
)

const DURATION time.Duration = time.Second

type Products struct {
	Products map[string]shop_competition.Product
	sync.RWMutex
}

func (products *Products) AddProductConc(product shop_competition.Product) error {
	ch := make(chan error, 1)
	go func() {
		ch <- products.AddProduct(product)
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(DURATION):
		//err = TimeoutError
		return TimeoutError
	}
}

func (products *Products) AddProduct(product shop_competition.Product) error {
	if product.Name == "" {
		return errors.New("product without name")
	}

	err := productCheck(product)
	if err != nil {
		return err
	}

	products.Lock()
	defer products.Unlock()

	if _, ok := products.Products[product.Name]; ok {
		return errors.New("product exist")
	}

	products.Products[product.Name] = product
	return nil
}

func (products *Products) ModifyProductConc(product shop_competition.Product) error {
	ch := make(chan error, 1)
	go func() {
		ch <- products.ModifyProduct(product)
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(DURATION):
		//err = TimeoutError
		return TimeoutError
	}
}

func (products *Products) ModifyProduct(product shop_competition.Product) error {

	err := productCheck(product)
	if err != nil {
		return err
	}
	products.Lock()
	defer products.Unlock()

	if _, ok := products.Products[product.Name]; !ok {
		return errors.New("product not found")
	}
	products.Products[product.Name] = product

	return nil
}

func (products *Products) RemoveProductConc(name string) error {
	ch := make(chan error, 1)
	go func() {
		ch <- products.RemoveProduct(name)
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(DURATION):
		//err = TimeoutError
		return TimeoutError
	}
}
func (products *Products) RemoveProduct(name string) error {
	products.Lock()
	defer products.Unlock()
	if _, ok := products.Products[name]; !ok {
		return errors.New("product not found")
	}
	delete(products.Products, name)
	return nil
}

func (products *Products) ImportProductsCSV(data []byte) error {
	var buf bytes.Buffer
	buf.Write(data)
	reader := csv.NewReader(&buf)
	//ctx, cancel := context.WithCancel()
	eg, ctx := errgroup.WithContext(context.Background())
	productsListMap := make([]map[string]shop_competition.Product, 0)
	records := make([][]string, 0)
	mutex := sync.Mutex{}
	i := 0
EOF:
	for {
		select {
		case <-ctx.Done():
			break EOF
		default:

		}
		readPackage := ImportNone
		record, err := reader.Read()
		if err == io.EOF {
			readPackage = ImportEof
		} else if err != nil {
			return err
		}

		if readPackage == ImportNone {
			records = append(records, record)
			i++
			if i%PACKAGE_SIZE == 0 {
				readPackage = ImportProcessing
			}
		}
		switch readPackage {
		case ImportProcessing, ImportEof:
			if readPackage != ImportEof || i%PACKAGE_SIZE != 0 {
				//fmt.Println(i)
				recordsGo := records
				records = make([][]string, 0)
				eg.Go(func() error {
					productsRes, errImport := ImportProductsCSVConc(recordsGo, eg, ctx)
					mutex.Lock()
					defer mutex.Unlock()

					if errImport != nil {
						return errImport
					}

					productsListMap = append(productsListMap, productsRes)
					return nil
				})
			}
			if readPackage == ImportEof {
				break EOF
			}
		}
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	productsRes := make(map[string]shop_competition.Product, i)
	for _, productsMap := range productsListMap {
		for key, value := range productsMap {
			productsRes[key] = value
		}
	}
	products.Lock()
	products.Products = productsRes
	products.Unlock()
	return nil
}
func ImportProductsCSVConc(data [][]string, wg *errgroup.Group, ctx context.Context) (map[string]shop_competition.Product, error) {
	products := make(map[string]shop_competition.Product)
	for _, line := range data {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		var product shop_competition.Product
		name := line[0]
		product.Name = name
		price, err := strconv.ParseFloat(line[1], 32)
		if err != nil {
			return nil, &ImportProductError{
				Product: product,
				Err:     err,
			}
		}
		product.Price = float32(price)
		typeProduct, err := strconv.ParseInt(line[2], 0, strconv.IntSize)
		if err != nil {
			return nil, &ImportProductError{
				Product: product,
				Err:     err,
			}
		}
		product.Type = shop_competition.ProductType(typeProduct)
		products[name] = product
	}
	return products, nil
}

func (products *Products) ExportProductsCSV() []byte {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	record := make([]string, 3)
	for _, product := range products.Products {
		record[0] = product.Name
		record[1] = fmt.Sprintf("%.2f", product.Price)
		record[2] = fmt.Sprintf("%v", product.Type)
		err := writer.Write(record)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	if err := writer.Error(); err != nil {
		panic(err)
	}

	return buf.Bytes()

}
func productCheck(product shop_competition.Product) error {
	if product.Type == shop_competition.ProductSample {
		if product.Price != 0 {
			return errors.New("sample was free in bundle")
		}
	} else {
		if product.Price <= 0 {
			return fmt.Errorf("product price %.2f not valid", product.Price)
		}
	}
	return nil
}
