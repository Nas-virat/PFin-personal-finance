package service

import (
	"strings"

	"github.com/Nas-virat/PFin-personal-finance/errs"
	"github.com/Nas-virat/PFin-personal-finance/model"
	"github.com/Nas-virat/PFin-personal-finance/repository"
)

type accountService struct {
	accRepo repository.AccountRepository
}

func NewAccountService(accRepo repository.AccountRepository) AccountService {
	return accountService{accRepo: accRepo}
}

func (s accountService) CreateAccount(account model.NewAccountRequest) (*model.AccountResponse, error) {

	// check account is negative or not
	if account.Amount < 0 {
		return nil, errs.NewVaildationError("Account intial Balance can not less than 0")
	}

	// check if Type is Bank or investment or not
	if strings.ToLower(account.Type) != "bank" && strings.ToLower(account.Type) != "investment" {
		return nil, errs.NewVaildationError("Account can not be " + account.Type)
	}

	newAccount := model.Account{
		AccountName: account.Name,
		Type:        account.Type,
		Amount:      account.Amount,
		Description: account.Description,
	}

	createAccount, err := s.accRepo.CreateAccount(newAccount)

	if err != nil {
		return nil, errs.NewUnexpectedError()
	}

	response := model.AccountResponse{
		AccountID: int(createAccount.ID),
		Opendate:  createAccount.CreatedAt,
	}

	return &response, nil
}

func (s accountService) GetAccountById(accountID int) (*model.AccountResponse, error) {
	return nil, nil
}

func (s accountService) GetAccounts() ([]model.AccountResponse, error) {
	return nil, nil
}
