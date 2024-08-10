package account_test

import (
	"testing"

	"github.com/Nas-virat/PFin-personal-finance/account"
	"github.com/stretchr/testify/mock"
)

type MockAccountRepo struct {
	mock.Mock
}

func (m *MockAccountRepo) CreateAccount(accountModel account.Account) (*account.Account, error) {
	args := m.Called()
	return args.Get(0).(*account.Account), args.Error(1)
}
func (m *MockAccountRepo) GetAccountById(id int) (*account.Account, error) {
	args := m.Called()
	return args.Get(0).(*account.Account), args.Error(1)
}

func (m *MockAccountRepo) GetAccounts() ([]account.Account, error) {
	args := m.Called()
	return args.Get(0).([]account.Account), args.Error(1)
}
func (m *MockAccountRepo) EditAccountInfo(accountModel account.Account, id int) (*account.Account, error) {
	args := m.Called()
	return args.Get(0).(*account.Account), args.Error(1)
}

func TestService(t *testing.T) {
	accountRequest := account.NewAccountRequest{
		Name:        "",
		Type:        "bank",
		Amount:      1.0,
		Description: "",
		Currency:    "",
	}
	mockRepo := MockAccountRepo{}
	mockRepo.On("CreateAccount").Return(&account.Account{}, nil)
	accountService := account.NewAccountService(&mockRepo)

	_, err := accountService.CreateAccount(accountRequest)

	if err != nil {
		t.Error(err)
	}

}

