package domain

type service struct {
	productRepo Repository
}

// create a new service along with a given repository
func NewProductService(productRepo Repository) Service {
	return &service{productRepo: productRepo}
}

func (s *service) Find(code string) (*Product, error) {
	return s.productRepo.Find(code)
}

func (s *service) Store(product *Product) error {
	return s.productRepo.Store(product)
}

func (s *service) Update(product *Product) error {
	return s.productRepo.Update(product)
}

func (s *service) FindAll() ([]*Product, error) {
	return s.productRepo.FindAll()
}

func (s *service) Delete(code string) error {
	return s.productRepo.Delete(code)
}
