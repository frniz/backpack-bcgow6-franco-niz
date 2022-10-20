package internal

type MockStore struct {
	ReadWasCalled bool
	TrBefore      Transaction
	db            []Transaction
	errorEsperado error
}

func (s *MockStore) Read(data interface{}) error {
	s.ReadWasCalled = true
	t1 := Transaction{
		ID:       1,
		Code:     "BeforeUpdate",
		Currency: "CurrencyPrueba1",
		Price:    1.0,
		Emitter:  "Prueba1",
		Receiver: "Prueba1Receiver",
		Date:     "UnaFecha",
	}
	t2 := Transaction{
		ID:       2,
		Code:     "001",
		Currency: "CurrencyPrueba2",
		Price:    2.0,
		Emitter:  "Prueba2",
		Receiver: "Prueba2Receiver",
		Date:     "UnaFecha2",
	}

	ts := []Transaction{
		t1,
		t2,
	}

	dataStub := data.(*[]Transaction)
	*dataStub = ts

	return nil
}
func (s *MockStore) Write(data interface{}) error {
	s.db = data.([]Transaction)
	return nil
}
