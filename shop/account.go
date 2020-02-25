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
	sorting "sort"
	"strconv"
	"sync"
	"time"
)

type Accounts struct {
	Accounts map[string]shop_competition.Account
	sync.RWMutex
}

func (accounts *Accounts) RegisterConc(name string) error {
	ch := make(chan error, 1)
	go func() {
		ch <- accounts.Register(name)
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(DURATION):
		return TimeoutError
	}
}

func (accounts *Accounts) Register(username string) error {
	accounts.Lock()
	defer accounts.Unlock()
	_, okay := accounts.Accounts[username]
	if okay {
		return errors.New("user already registered")
	}

	accounts.Accounts[username] = shop_competition.Account{
		Name:        username,
		Balance:     0,
		AccountType: shop_competition.AccountNormal,
	}

	return nil
}

func (accounts *Accounts) NewAccountConc(username string, name string, balance float32, accountType shop_competition.AccountType) error {
	ch := make(chan error, 1)
	go func() {
		ch <- accounts.NewAccount(username, name, balance, accountType)
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(DURATION):
		return TimeoutError
	}
}
func (accounts *Accounts) NewAccount(username string, name string, balance float32, accountType shop_competition.AccountType) error {
	accounts.Lock()
	defer accounts.Unlock()

	err := accounts.Register(username)
	if err != nil {
		return err
	}
	accounts.Accounts[username] = shop_competition.Account{
		Name:        name,
		Balance:     balance,
		AccountType: accountType,
	}
	return nil
}

func (accounts *Accounts) AddBalanceConc(username string, sum float32) error {
	ch := make(chan error, 1)
	go func() {
		ch <- accounts.AddBalance(username, sum)
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(DURATION):
		return TimeoutError
	}
}
func (accounts *Accounts) AddBalance(username string, sum float32) error {
	if sum <= 0 {
		return errors.New("sum only positive value")
	}
	accounts.Lock()
	defer accounts.Unlock()
	user, okay := accounts.Accounts[username]
	if okay {
		return errors.New("user not found")
	}

	user.Balance = (ToMoney(user.Balance) + ToMoney(sum)).Float32()
	accounts.Accounts[username] = user

	return nil
}

func (accounts *Accounts) GetBalanceConc(name string) (float32, error) {
	res, ch := make(chan float32, 1), make(chan error, 1)
	go func() {
		balance, err := accounts.GetBalance(name)
		ch <- err
		res <- balance
	}()
	select {
	case err := <-ch:
		balance := <-res
		return balance, err
	case <-time.After(DURATION):
		return 0, TimeoutError
	}
}

func (accounts *Accounts) GetBalance(username string) (float32, error) {
	accounts.RLock()
	defer accounts.RUnlock()
	user, okay := accounts.Accounts[username]
	if okay {
		return 0, errors.New("user not found")
	}
	return user.Balance, nil
}
func (accounts *Accounts) GetAccountsConc(sort shop_competition.AccountSortType) []shop_competition.Account {
	ch := make(chan []shop_competition.Account, 1)
	go func() {
		ch <- accounts.GetAccounts(sort)
	}()
	select {
	case accounts := <-ch:
		return accounts
	case <-time.After(DURATION):
		return make([]shop_competition.Account, 0)
	}
}

func (accounts *Accounts) GetAccounts(sort shop_competition.AccountSortType) []shop_competition.Account {
	accounts.RLock()
	accountsRes := make([]shop_competition.Account, len(accounts.Accounts))
	i := 0
	for _, v := range accounts.Accounts {
		accountsRes[i] = v
		i++
	}
	accounts.RUnlock()
	switch sort {
	case shop_competition.SortByName:
		sorting.Slice(accountsRes, func(i, j int) bool {
			return accountsRes[i].Name < accountsRes[j].Name
		})
	case shop_competition.SortByNameReverse:
		sorting.Slice(accountsRes, func(i, j int) bool {
			return accountsRes[i].Name > accountsRes[j].Name
		})
	case shop_competition.SortByBalance:
		sorting.Slice(accountsRes, func(i, j int) bool {
			return accountsRes[i].Balance > accountsRes[j].Balance
		})
	}
	return accountsRes

}

func ImportAccountsCSVConc(data [][]string, wg *errgroup.Group, ctx context.Context) (map[string]shop_competition.Account, error) {
	accounts := make(map[string]shop_competition.Account)
	for _, line := range data {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		var account shop_competition.Account
		name := line[0]
		account.Name = name
		balance, err := strconv.ParseFloat(line[1], 32)
		if err != nil {
			return nil, &ImportAccountError{
				Account: account,
				Err:     err,
			}
		}
		account.Balance = float32(balance)
		typeAccount, err := strconv.ParseInt(line[2], 0, strconv.IntSize)
		if err != nil {
			return nil, &ImportAccountError{
				Account: account,
				Err:     err,
			}
		}
		account.AccountType = shop_competition.AccountType(typeAccount)
		accounts[name] = account
	}
	return accounts, nil
}
func (accounts *Accounts) ImportAccountsCSV(data []byte) error {
	var buf bytes.Buffer
	buf.Write(data)
	reader := csv.NewReader(&buf)
	//ctx, cancel := context.WithCancel()
	eg, ctx := errgroup.WithContext(context.Background())
	accountListMap := make([]map[string]shop_competition.Account, 0)
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
					accountsLRes, errImport := ImportAccountsCSVConc(recordsGo, eg, ctx)
					mutex.Lock()
					defer mutex.Unlock()

					if errImport != nil {
						return errImport
					}

					accountListMap = append(accountListMap, accountsLRes)
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

	accountsRes := make(map[string]shop_competition.Account, i)
	for _, accountMap := range accountListMap {
		for key, value := range accountMap {
			accountsRes[key] = value
		}
	}
	accounts.Lock()
	accounts.Accounts = accountsRes
	accounts.Unlock()
	return nil
}
func (accounts *Accounts) ExportAccountsCSV() []byte {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	record := make([]string, 3)
	for _, account := range accounts.Accounts {
		record[0] = account.Name
		record[1] = fmt.Sprintf("%.2f", account.Balance)
		record[2] = fmt.Sprintf("%v", account.AccountType)
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
