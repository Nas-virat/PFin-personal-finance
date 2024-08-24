package account_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Nas-virat/PFin-personal-finance/account"
	"github.com/Nas-virat/PFin-personal-finance/errs"
	"github.com/Nas-virat/PFin-personal-finance/response"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAccountService struct {
	mock.Mock
}

func (m *MockAccountService) CreateAccount(accountRequest account.NewAccountRequest) (*account.NewAccountResponse, error) {
	args := m.Called(accountRequest)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*account.NewAccountResponse), args.Error(1)
}

func (m *MockAccountService) GetAccountById(accountID int) (*account.AccountResponse, error) {
	args := m.Called(accountID)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*account.AccountResponse), args.Error(1)
}

func (m *MockAccountService) GetAccounts() ([]account.AccountResponse, error) {
	args := m.Called()
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]account.AccountResponse), args.Error(1)
}

func (m *MockAccountService) EditAccountInfo(accountRequest account.NewAccountRequest, id int) (*account.NewAccountResponse, error) {
	args := m.Called(accountRequest, id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*account.NewAccountResponse), args.Error(1)
}

func TestCreateAccountHandler(t *testing.T) {

	t.Run("should create Account", func(t *testing.T) {
		//arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		request := account.NewAccountRequest{
			Name:        "test",
			Type:        "savings",
			Amount:      1000,
			Description: "test account",
			Currency:    "USD",
		}

		mockAccountService := &MockAccountService{}
		mockResponse := account.NewAccountResponse{
			AccountID: 1,
			Opendate:  time.Time{},
			Type:      "savings",
			Amount:    1000,
			Status:    true,
		}

		mockAccountService.On("CreateAccount", request).Return(&mockResponse, nil)

		handler := account.NewAccountHandler(mockAccountService)
		body, _ := json.Marshal(request)

		want, _ := json.Marshal(response.Response{
			Success: true,
			Message: "insert successfully",
			Data:    mockResponse,
		})

		//act
		c.Request, _ = http.NewRequest(http.MethodPost, "/api/account/create", bytes.NewBuffer(body))
		handler.CreateAccountHandler(c)

		//assert
		assert.Equal(t, http.StatusCreated, w.Code)
		assert.JSONEq(t, string(want), w.Body.String())
	})

	t.Run("should return error if request can not bind", func(t *testing.T) {
		//arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		request := account.NewAccountRequest{
			Name:        "test",
			Type:        "savings",
			Amount:      1000,
			Description: "test account",
			Currency:    "USD",
		}
		fake := `{"key":"1","name":"test","description":"test","logo":"test","tags":["tag"]}}`
		fakeRequest, _ := json.Marshal(fake)

		mockAccountService := &MockAccountService{}
		mockResponse := account.NewAccountResponse{
			AccountID: 1,
			Opendate:  time.Time{},
			Type:      "savings",
			Amount:    1000,
			Status:    true,
		}
		mockAccountService.On("CreateAccount", request).Return(&mockResponse, nil)

		handler := account.NewAccountHandler(mockAccountService)

		//act
		c.Request, _ = http.NewRequest(http.MethodPost, "/api/account/create", bytes.NewBuffer(fakeRequest))
		handler.CreateAccountHandler(c)

		//assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return error if service return error", func(t *testing.T) {
		//arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		request := account.NewAccountRequest{
			Name:        "test",
			Type:        "savings",
			Amount:      1000,
			Description: "test account",
			Currency:    "USD",
		}

		mockAccountService := &MockAccountService{}
		mockAccountService.On("CreateAccount", request).Return(nil, errs.NewVaildationError(""))

		handler := account.NewAccountHandler(mockAccountService)
		body, _ := json.Marshal(request)

		//act
		c.Request, _ = http.NewRequest(http.MethodPost, "/api/account/create", bytes.NewBuffer(body))
		handler.CreateAccountHandler(c)

		//assert
		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})
}

func TestGetAccountByIdHandler(t *testing.T) {

	t.Run("should return Account when AccountId is exist", func(t *testing.T) {

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		id := 1

		accountResponse := &account.AccountResponse{
			AccountID:   1,
			AccountName: "Name",
			Type:        "bank",
			Amount:      300.0,
			Description: "Hello",
			Status:      true,
		}

		// Set the route parameters manually
		c.Params = gin.Params{
			{Key: "id", Value: "1"},
		}

		mockAccountService := &MockAccountService{}
		mockAccountService.On("GetAccountById", id).Return(accountResponse, nil)
		handler := account.NewAccountHandler(mockAccountService)

		want, _ := json.Marshal(response.Response{
			Success: true,
			Message: "get account by id",
			Data:    accountResponse,
		})

		//act
		c.Request, _ = http.NewRequest(http.MethodGet, "/api/account/1", nil)
		handler.GetAccountByIdHandler(c)

		//assert
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, string(want), w.Body.String())
	})

	t.Run("should return Error when AccountId is not integer", func(t *testing.T) {

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		mockAccountService := &MockAccountService{}
		handler := account.NewAccountHandler(mockAccountService)

		//act
		c.Request, _ = http.NewRequest(http.MethodGet, "/api/account/abc", nil)
		handler.GetAccountByIdHandler(c)

		//assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestGetAccountHandler(t *testing.T) {

	t.Run("should return accounts if repo successfully", func(t *testing.T) {

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		accountResponse := []account.AccountResponse{
			{
				AccountID:   1,
				AccountName: "Name",
				Type:        "bank",
				Amount:      300.0,
				Description: "Hello",
				Status:      true,
			},
			{
				AccountID:   2,
				AccountName: "Find",
				Type:        "bank",
				Amount:      300.0,
				Description: "Hello world",
				Status:      true,
			},
		}
		mockAccountService := &MockAccountService{}
		mockAccountService.On("GetAccounts").Return(accountResponse, nil)
		handler := account.NewAccountHandler(mockAccountService)

		want, _ := json.Marshal(response.Response{
			Success: true,
			Message: "get all accounts",
			Data:    accountResponse,
		})

		c.Request, _ = http.NewRequest(http.MethodGet, "/api/account", nil)
		handler.GetAccountsHandler(c)

		//assert
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, string(want), w.Body.String())
	})

	t.Run("should return error if repo fail to get data", func(t *testing.T) {

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		mockAccountService := &MockAccountService{}
		mockAccountService.On("GetAccounts").Return(nil, errs.NewUnexpectedError())
		handler := account.NewAccountHandler(mockAccountService)

		c.Request, _ = http.NewRequest(http.MethodGet, "/api/account", nil)
		handler.GetAccountsHandler(c)

		//assert
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
