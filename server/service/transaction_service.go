package service

import (
	"github.com/Nas-virat/PFin-personal-finance/errs"
	"github.com/Nas-virat/PFin-personal-finance/model"
	"github.com/Nas-virat/PFin-personal-finance/repository"
)

type transactionService struct {
	transactionRepo repository.TransactionRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository) TransactionService {
	return &transactionService{transactionRepo: transactionRepo}
}


func(s transactionService) CreateTransaction(transactionRequest model.NewTransactionRequest) (*model.TransactionResponse, error){

	// Validate Transaction
	if transactionRequest.Amount <= 0 {
		return nil, errs.NewVaildationError("Amount must be greater than 0")
	}

	transaction := model.Transaction{
		TransactionType: transactionRequest.TransactionType,
		Catergory: transactionRequest.Catergory,
		Description: transactionRequest.Description,
		Amount: transactionRequest.Amount,
	}

	transactionResult, err := s.transactionRepo.CreateTransaction(transaction)

	if err != nil {
		return nil, errs.NewUnexpectedError()
	}

	transactionResponse := model.TransactionResponse{
		CreateAt: transactionResult.CreatedAt,
		TransactionType: transactionResult.TransactionType,
		Catergory:transactionResult.Catergory,
		Description: transactionResult.Description,
		Amount: transactionResult.Amount,
	}


	return &transactionResponse,nil
}

func(s transactionService) GetTransactionByID(id uint) (*model.TransactionResponse, error){

	transaction, err := s.transactionRepo.GetTransactionByID(id)

	if err != nil {
		return nil, errs.NewUnexpectedError()
	}

	transactionResponse := model.TransactionResponse{
		CreateAt: transaction.CreatedAt,
		TransactionType: transaction.TransactionType,
		Catergory: transaction.Catergory,
		Description: transaction.Description,
		Amount: transaction.Amount,
	}

	return &transactionResponse,nil
}

func(s transactionService) GetTransactions() ([]model.TransactionResponse, error){

	transactions, err := s.transactionRepo.GetTransactions()

	if err != nil{
		return nil, errs.NewUnexpectedError()
	}

	transactionResponses := []model.TransactionResponse{}

	for _ , transaction := range transactions{
		transactionResponses = append(transactionResponses,
			model.TransactionResponse{
				CreateAt: transaction.CreatedAt,
				TransactionType: transaction.TransactionType,
				Catergory: transaction.Catergory,
				Description: transaction.Description,
				Amount: transaction.Amount,
			},
		)
	}

	return transactionResponses,nil
}

func(s transactionService) UpdateTransaction(id uint, newInfo model.Transaction) (*model.TransactionResponse, error){
	
	// Validate newInfo transaction
	if newInfo.Amount <= 0{
		return nil, errs.NewVaildationError("Amount need to more than 0")
	}

	transaction, err := s.transactionRepo.UpdateTransaction(id,newInfo)

	if err != nil{
		return nil, errs.NewUnexpectedError()
	}

	transactionResponse := model.TransactionResponse{
		CreateAt: transaction.CreatedAt,
		TransactionType: transaction.TransactionType,
		Catergory: transaction.Catergory,
		Description: transaction.Description,
		Amount: transaction.Amount,
	}

	return &transactionResponse, nil
}

func(s transactionService) DeleteTransaction(id uint) (*model.TransactionResponse, error){
	transaction, err := s.transactionRepo.DeleteTransaction(id)

	if err != nil{
		return nil, errs.NewUnexpectedError()
	}

	transactionResponse := model.TransactionResponse{
		CreateAt: transaction.CreatedAt,
		TransactionType: transaction.TransactionType,
		Catergory: transaction.Catergory,
		Description: transaction.Description,
		Amount: transaction.Amount,
	}

	return &transactionResponse, nil
}


