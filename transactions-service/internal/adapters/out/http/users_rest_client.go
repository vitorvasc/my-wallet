package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"transactions-service/internal/core/domain"
)

const (
	defaultHost = "http://account-balance-service:8080"

	pathGetUserByID       = "/v1/users/%d"
	pathGetAccountBalance = "/v1/users/%d/account_balance"
)

type UsersRestClient struct {
	client http.Client
	Host   string
}

func NewUsersRestClient() *UsersRestClient {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	return &UsersRestClient{
		client: client,
		Host:   defaultHost,
	}
}

func (c *UsersRestClient) GetUserByID(userID uint64) (*domain.User, error) {
	formattedPath := fmt.Sprintf(pathGetUserByID, userID)

	res, err := c.client.Get(fmt.Sprintf("%s%s", c.Host, formattedPath))
	if err != nil {
		log.Printf("[GetUserByID] error obtaining user: %v", err)
		return nil, domain.ErrObtainingUserByID
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, domain.ErrUserNotFound
	}

	defer res.Body.Close()
	user := new(domain.User)
	err = json.NewDecoder(res.Body).Decode(user)
	if err != nil {
		log.Printf("[GetUserByID] error parsing user: %v", err)
		return nil, domain.ErrParsingUserResponse
	}

	return user, nil
}

func (c *UsersRestClient) GetAccountBalance(userID uint64) (*domain.AccountBalance, error) {
	formattedPath := fmt.Sprintf(pathGetAccountBalance, userID)

	res, err := c.client.Get(fmt.Sprintf("%s%s", c.Host, formattedPath))
	if err != nil {
		log.Printf("[GetAccountBalance] error obtaining user: %v", err)
		return nil, domain.ErrObtainingAccountBalance
	}
	defer res.Body.Close()

	accountBalance := new(domain.AccountBalance)
	err = json.NewDecoder(res.Body).Decode(accountBalance)
	if err != nil {
		log.Printf("[GetAccountBalance] error parsing account balance: %v", err)
		return nil, domain.ErrParsingAccountBalanceResponse
	}

	return accountBalance, nil
}
