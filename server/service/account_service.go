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

func (s accountService) CreateAccount(account model.NewAccountRequest) (*model.NewAccountResponse, error) {

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

	response := model.NewAccountResponse{
		AccountID: int(createAccount.ID),
		Opendate:  createAccount.CreatedAt,
		Type:	  createAccount.Type,
		Amount:   createAccount.Amount,
		Status:   true,
	}

	return &response, nil
}

func (s accountService) GetAccountById(accountID int) (*model.AccountResponse, error) {

	account, err := s.accRepo.GetAccountById(accountID)

	if err != nil {
		return nil, errs.NewUnexpectedError()
	}
	accountResponse := model.AccountResponse{
		AccountID: int(account.ID),
		AccountName: account.AccountName,
		Type: account.Type,
		Amount: account.Amount,
		Description: account.Description,
		Status: account.Status,
	}

	return &accountResponse, nil
}

func (s accountService) GetAccounts() ([]model.AccountResponse, error) {

	accounts, err := s.accRepo.GetAccounts()
	if err != nil{
		return nil, errs.NewUnexpectedError()
	}

	accountResponses := []model.AccountResponse{}

	for _, account := range accounts{
		accountResponses = append(accountResponses, 
			model.AccountResponse{
				AccountID: int(account.ID),
				AccountName: account.AccountName,
				Type: account.Type,
				Amount: account.Amount,
				Description: account.Description,
				Status: account.Status,
			},
		)
	}

	return accountResponses, nil
}

func (s accountService) EditAccountInfo(account model.NewAccountRequest, id int) (*model.NewAccountResponse, error){

	// check account is negative or not
	if account.Amount < 0 {
		return nil, errs.NewVaildationError("Account intial Balance can not less than 0")
	}

	// check if Type is Bank or investment or not
	if strings.ToLower(account.Type) != "bank" && strings.ToLower(account.Type) != "investment" {
		return nil, errs.NewVaildationError("Account can not be " + account.Type)
	}

	// check if this id is exist or not
	accountCheck, err := s.accRepo.GetAccountById(id)
	if accountCheck == nil{
		return nil, errs.NewVaildationError("This account id is not exist")
	}

	accountRepoInsert := model.Account{
		AccountName: accountCheck.AccountName,
		Type: account.Type,
		Amount: account.Amount,
		Description: account.Description,
		Status: accountCheck.Status,
	}


	// update account
	updatedAccount, err := s.accRepo.EditAccountInfo(accountRepoInsert,id)

	if err != nil{
		return nil, errs.NewUnexpectedError()
	}

	accResponse := model.NewAccountResponse{
		AccountID: int(updatedAccount.ID),
		Opendate:  updatedAccount.CreatedAt,
		Type:	  updatedAccount.Type,
		Amount:   updatedAccount.Amount,
		Status:   true,
	}


	return &accResponse, nil
}