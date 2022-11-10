package internal

type Service interface {
	GetAll() ([]Transaction, error)
	Get(id int64) (Transaction, error)
	GetOne(code string) (Transaction, error)
	Store(code, currency string, price float64, emitter, receiver, date string) (int64, error)
	Update(id int, code, currency string, price float64, emitter, receiver, date string) error
	PartialUpdate(id int, code string, price float64) (Transaction, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]Transaction, error) {
	rep, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return rep, nil
}

func (s *service) Get(id int64) (Transaction, error) {
	return s.repository.Get(id)
}

func (s *service) GetOne(code string) (Transaction, error) {
	return s.repository.GetOne(code)
}
func (s *service) Store(code, currency string, price float64, emitter, receiver, date string) (int64, error) {
	return s.repository.Store(code, currency, price, emitter, receiver, date)
}

func (s *service) Update(id int, code, currency string, price float64, emmiter, receiver, date string) error {

	return s.repository.Update(id, code, currency, price, emmiter, receiver, date)
}

func (s *service) PartialUpdate(id int, code string, price float64) (Transaction, error) {
	return s.repository.PartialUpdate(id, code, price)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
