package account_test

import (
	"testing"

	"github.com/Nas-virat/PFin-personal-finance/account"
	"github.com/Nas-virat/PFin-personal-finance/errs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAccountRepo struct {
	mock.Mock
}

func (m *MockAccountRepo) CreateAccount(accountModel account.Account) (*account.Account, error) {
	args := m.Called()
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*account.Account), args.Error(1)
}
func (m *MockAccountRepo) GetAccountById(id int) (*account.Account, error) {
	args := m.Called()
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*account.Account), args.Error(1)
}

func (m *MockAccountRepo) GetAccounts() ([]account.Account, error) {
	args := m.Called()
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]account.Account), args.Error(1)
}
func (m *MockAccountRepo) EditAccountInfo(accountModel account.Account, id int) (*account.Account, error) {
	args := m.Called()
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*account.Account), args.Error(1)
}

func TestCreateAccountService(t *testing.T) {
	t.Run("should create a new account", func(t *testing.T) {

		accountRequest := account.NewAccountRequest{
			Name:        "Nas-virat",
			Type:        "bank",
			Amount:      1.0,
			Description: "",
			Currency:    "",
		}
		mockRepo := MockAccountRepo{}
		mockRepo.On("CreateAccount").Return(&account.Account{Amount: accountRequest.Amount}, nil)
		accountService := account.NewAccountService(&mockRepo)

		result, err := accountService.CreateAccount(accountRequest)

		assert.Equal(t, result.Amount, accountRequest.Amount)
		assert.Equal(t, err, nil)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return err if the account type is not bank or investment", func(t *testing.T) {

		accountRequest := account.NewAccountRequest{
			Name:        "Nas-virat",
			Type:        "",
			Amount:      1.0,
			Description: "",
			Currency:    "",
		}
		mockRepo := MockAccountRepo{}
		mockRepo.On("CreateAccount").Return(&account.Account{Amount: accountRequest.Amount}, nil)
		accountService := account.NewAccountService(&mockRepo)

		result, err := accountService.CreateAccount(accountRequest)

		assert.Nil(t, result, nil)
		assert.NotEqual(t, err, nil)
	})

	t.Run("should return err if account amount less than 0", func(t *testing.T) {

		accountRequest := account.NewAccountRequest{
			Name:        "Nas-virat",
			Type:        "bank",
			Amount:      -1.0,
			Description: "",
			Currency:    "",
		}
		mockRepo := MockAccountRepo{}
		mockRepo.On("CreateAccount").Return(&account.Account{Amount: accountRequest.Amount}, nil)
		accountService := account.NewAccountService(&mockRepo)

		result, err := accountService.CreateAccount(accountRequest)

		assert.Nil(t, result, nil)
		assert.NotEqual(t, err, nil)
	})

	t.Run("should return err if accountRepo throw err", func(t *testing.T) {

		accountRequest := account.NewAccountRequest{
			Name:        "Nas-virat",
			Type:        "bank",
			Amount:      1.0,
			Description: "",
			Currency:    "",
		}
		mockRepo := MockAccountRepo{}
		mockRepo.On("CreateAccount").Return(nil, errs.NewUnexpectedError())
		accountService := account.NewAccountService(&mockRepo)

		result, err := accountService.CreateAccount(accountRequest)

		assert.Nil(t, result, nil)
		assert.NotEqual(t, err, nil)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetAccountByIdService(t *testing.T) {

	t.Run("should get account by id if Repo can get account", func(t *testing.T) {

		accountId := 1

		mockAccount := account.Account{
			AccountName: "Nas-virat",
			Type:        "bank",
			Amount:      100.0,
			Description: "",
			Currency:    "",
		}
		mockRepo := MockAccountRepo{}
		mockRepo.On("GetAccountById").Return(&mockAccount, nil)
		accountService := account.NewAccountService(&mockRepo)

		result, err := accountService.GetAccountById(accountId)

		assert.Equal(t, result.AccountID, accountId)
		assert.Equal(t, result.AccountName, mockAccount.AccountName)
		assert.Equal(t, result.Type, mockAccount.Type)
		assert.Equal(t, result.Amount, mockAccount.Amount)
		assert.Equal(t, result.Description, mockAccount.Description)
		assert.Equal(t, err, nil)
	})

	t.Run("should get account by id if Repo cannot get account", func(t *testing.T) {

		accountId := 1

		mockRepo := MockAccountRepo{}
		mockRepo.On("GetAccountById").Return(nil, errs.NewUnexpectedError())
		accountService := account.NewAccountService(&mockRepo)

		result, err := accountService.GetAccountById(accountId)

		assert.Nil(t, result)
		assert.NotEqual(t, err, nil)

	})
}

func TestGetAccount(t *testing.T) {

	t.Run("should get account by id if Repo can get account", func(t *testing.T) {

		mockAccount := []account.Account{
			{
				AccountName: "Nas-virat",
				Type:        "bank",
				Amount:      100.0,
				Description: "",
				Currency:    "",
			},
			{
				AccountName: "Nas-virat1",
				Type:        "bank",
				Amount:      200.0,
				Description: "Hello world",
				Currency:    "",
			},
		}
		mockRepo := MockAccountRepo{}
		mockRepo.On("GetAccounts").Return(mockAccount, nil)
		accountService := account.NewAccountService(&mockRepo)

		result, err := accountService.GetAccounts()

		assert.Equal(t, mockAccount[0].AccountName, result[0].AccountName)
		assert.Equal(t, mockAccount[0].Type, result[0].Type)
		assert.Equal(t, mockAccount[0].Amount, result[0].Amount)
		assert.Equal(t, mockAccount[0].Description, result[0].Description)
		assert.Equal(t, mockAccount[1].AccountName, result[1].AccountName)
		assert.Equal(t, mockAccount[1].Type, result[1].Type)
		assert.Equal(t, mockAccount[1].Amount, result[1].Amount)
		assert.Equal(t, mockAccount[1].Description, result[1].Description)
		assert.Equal(t, err, nil)
		mockRepo.MethodCalled("GetAccounts")
	})

	t.Run("should return error if Repo cannot get account", func(t *testing.T) {

		mockRepo := MockAccountRepo{}
		mockRepo.On("GetAccounts").Return(nil, errs.NewUnexpectedError())
		accountService := account.NewAccountService(&mockRepo)

		result, err := accountService.GetAccounts()

		assert.Nil(t, result)
		assert.NotEqual(t, err, nil)
	})
}

func TestEditAccountInfo(t *testing.T) {

	t.Run("should return error if Name is empty", func(t *testing.T) {

		mockAccountRequest := account.NewAccountRequest{
			Name: "",
		}

		id := 1

		mockAccount := account.Account{
			AccountName: "",
			Type:        "bank",
			Amount:      100.0,
			Description: "",
			Currency:    "",
		}
		mockRepo := MockAccountRepo{}
		mockRepo.On("EditAccountInfo").Return(mockAccount, nil)
		accountService := account.NewAccountService(&mockRepo)

		result, err := accountService.EditAccountInfo(mockAccountRequest, id)

		assert.Nil(t, result)
		assert.NotNil(t, err)
	})

	t.Run("should return error if account amount is less than 0", func(t *testing.T) {

		mockAccountRequest := account.NewAccountRequest{
			Amount: -1,
		}

		id := 1

		mockAccount := account.Account{
			AccountName: "",
			Type:        "bank",
			Amount:      -1,
			Description: "",
			Currency:    "",
		}
		mockRepo := MockAccountRepo{}
		mockRepo.On("EditAccountInfo").Return(mockAccount, nil)
		accountService := account.NewAccountService(&mockRepo)

		result, err := accountService.EditAccountInfo(mockAccountRequest, id)

		assert.Nil(t, result)
		assert.NotNil(t, err)
	})

	t.Run("should return error if account type is not bank or investment ", func(t *testing.T) {

		mockAccountRequest := account.NewAccountRequest{
			Type: "h",
		}

		id := 1

		mockAccount := account.Account{
			AccountName: "Nas-virat",
			Type:        "h",
			Amount:      100,
			Description: "",
			Currency:    "",
		}
		mockRepo := MockAccountRepo{}
		mockRepo.On("EditAccountInfo").Return(mockAccount, nil)
		accountService := account.NewAccountService(&mockRepo)

		result, err := accountService.EditAccountInfo(mockAccountRequest, id)

		assert.Nil(t, result)
		assert.NotNil(t, err)
	})

	t.Run("should return error if account repo return err", func(t *testing.T) {

		mockAccountRequest := account.NewAccountRequest{
			Name:   "Nas-virat",
			Amount: 100,
			Type:   "bank",
		}

		mockAccount := account.Account{
			AccountName: "Nas-virat",
			Type:        "bank",
			Amount:      100.0,
			Description: "",
			Currency:    "",
		}

		id := 1

		mockRepo := MockAccountRepo{}
		mockRepo.On("GetAccountById").Return(&mockAccount, nil)
		mockRepo.On("EditAccountInfo").Return(nil, errs.NewUnexpectedError())
		accountService := account.NewAccountService(&mockRepo)

		result, err := accountService.EditAccountInfo(mockAccountRequest, id)

		assert.Nil(t, result)
		assert.NotNil(t, err)
	})
}
