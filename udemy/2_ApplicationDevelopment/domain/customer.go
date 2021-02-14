package domain

type Customer struct {
	Id           string
	Name         string
	City         string
	Zipcode      string
	DateofBitrth string
	Status       string
}

type CustomerRepository interface { // biasanya orang kasih "I" di depan nama interface mereka

	// take nothing, but it will return a slice of cust and an error
	FindAll() ([]Customer, error)
}
