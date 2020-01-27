package shop

import (
	"errors"
	"github.com/youricorocks/shop_competition"
	sorting "sort"
)

type MyAccount shop_competition.Account

func (shop Shop) Register(username string) error {
	_, okay := shop.Accounts[username]
	if okay {
		return errors.New("user already registered")
	}
	shop.Accounts[username] = shop_competition.Account{}
	return nil
}

func (shop Shop) NewAccount(username string, name string, balance float32, accountType shop_competition.AccountType) error {
	err := shop.Register(username)
	if err != nil {
		return err
	}
	shop.Accounts[username] = shop_competition.Account{
		Name:        name,
		Balance:     balance,
		AccountType: accountType,
	}
	return nil
}

func (shop Shop) AddBalance(username string, sum float32) error {
	if sum <= 0.0 {
		return errors.New("sum very small")
	}
	user, okay := shop.Accounts[username]
	if okay {
		return errors.New("user not found")
	}
	user.Balance += sum
	shop.Accounts[username] = user

	return nil
}

func (shop Shop) GetBalance(username string) (float32, error) {
	user, okay := shop.Accounts[username]
	if okay {
		return 0, errors.New("user not found")
	}
	return user.Balance, nil
}

func (shop Shop) GetAccounts(sort shop_competition.AccountSortType) []shop_competition.Account {
	accounts := make([]shop_competition.Account, 0, len(shop.Accounts))

	for _, v := range shop.Accounts {
		accounts = append(accounts, v)
	}
	switch sort {
	case shop_competition.SortByName:
		sorting.Slice(accounts, func(i, j int) bool {
			return accounts[i].Name < accounts[j].Name
		})
	case shop_competition.SortByNameReverse:
		sorting.Slice(accounts, func(i, j int) bool {
			return accounts[i].Name > accounts[j].Name
		})
	case shop_competition.SortByBalance:
		sorting.Slice(accounts, func(i, j int) bool {
			return accounts[i].Balance < accounts[j].Balance
		})
	}
	return accounts

}
