package service

import "2_ApplicationDevelopment/domain"

type CustomerService interface {
	// one of the behaviour ada GetAllCustumer dan return domain.Customer
	GetAllCustomer() ([]domain.Customer, error)
}

// the implementation
type DefaultCustomerService struct {
	// lets make a repository
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustumer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
