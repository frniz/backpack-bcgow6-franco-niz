package internal

// Se debe separar la estructura del proyecto y como primer paso generando el paquete internal, en el paquete internal deben
// estar todas las funcionalidades que no dependan de paquetes externos.

// Dentro del paquete deben estar las capas:
// Servicio, debe contener la lógica de nuestra aplicación.
// Se debe crear el archivo service.go.
// Se debe generar la interface Service con todos sus métodos.
// Se debe generar la estructura service que contenga el repositorio.
// Se debe generar una función que devuelva el Servicio.
// Se deben implementar todos los métodos correspondientes a las operaciones a realizar (GetAll, Store, etc..).
// Repositorio, debe tener el acceso a la variable guardada en memoria.
// Se debe crear el archivo repository.go
// Se debe crear la estructura de la entidad
// Se deben crear las variables globales donde guardar las entidades
// Se debe generar la interface Repository con todos sus métodos
// Se debe generar la estructura repository
// Se debe generar una función que devuelva el Repositorio
// Se deben implementar todos los métodos correspondientes a las operaciones a realizar (GetAll, Store, etc..)

type Service interface {
	GetAll() ([]Transaction, error)
	Store(code, currency string, price float64, emmiter, receiver, date string) (Transaction, error)
	Update(id int, code, currency string, price float64, emmiter, receiver, date string) (Transaction, error)
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
func (s *service) Store(code, currency string, price float64, emmiter, receiver, date string) (Transaction, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Transaction{}, err
	}

	lastID++

	transaction, err := s.repository.Store(lastID, code, currency, price, emmiter, receiver, date)

	if err != nil {
		return Transaction{}, err
	}

	return transaction, nil
}

func (s *service) Update(id int, code, currency string, price float64, emmiter, receiver, date string) (Transaction, error) {

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
