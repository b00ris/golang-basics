package shop

import (
	"errors"
	"github.com/youricorocks/shop_competition"
	sorting "sort"
)

type Accounts map[string]shop_competition.Account

func (accounts Accounts) Register(username string) error {
	_, okay := accounts[username]
	if okay {
		return errors.New("user already registered")
	}

	accounts[username] = shop_competition.Account{
		Name:        username,
		Balance:     0,
		AccountType: shop_competition.AccountNormal,
	}

	return nil
}

func (accounts Accounts) NewAccount(username string, name string, balance float32, accountType shop_competition.AccountType) error {
	err := accounts.Register(username)
	if err != nil {
		return err
	}
	accounts[username] = shop_competition.Account{
		Name:        name,
		Balance:     balance,
		AccountType: accountType,
	}
	return nil
}

func (accounts Accounts) AddBalance(username string, sum float32) error {
	if sum <= 0 {
		return errors.New("sum only positive value")
	}
	user, okay := accounts[username]
	if okay {
		return errors.New("user not found")
	}

	user.Balance = (ToMoney(user.Balance) + ToMoney(sum)).Float32()
	accounts[username] = user

	return nil
}

func (accounts Accounts) GetBalance(username string) (float32, error) {
	user, okay := accounts[username]
	if okay {
		return 0, errors.New("user not found")
	}
	return user.Balance, nil
}

func (accounts Accounts) GetAccounts(sort shop_competition.AccountSortType) []shop_competition.Account {
	accountsRes := make([]shop_competition.Account, len(accounts))

	i := 0
	for _, v := range accounts {
		accountsRes[i] = v
		i++
	}
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
