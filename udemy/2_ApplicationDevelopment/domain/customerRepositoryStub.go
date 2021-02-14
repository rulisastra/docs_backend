package domain

// the adapter
type CustomerRepositoryStub struct {
	customers []Customer
}

// we want a func and attach it to the CustomerRepositoryStub |
// the fuction used will be FindALL who takes nothing () |
// returned a slice of Customer dan error
func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

// nama function | return CustomerRepositoryStub
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{ // variable simpan dummy customers
		{"1001", "Ruli", "Palembang", "30163", "2002-01-02", "1"},
		{"1002", "Sastra", "Palembang", "30163", "2002-02-02", "2"},
	}
	return CustomerRepositoryStub{customers} // rerurn list dari variabel
}
