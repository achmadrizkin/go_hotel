package kamar

type Service interface {
	FindAll() ([]Kamar, error)
	FindByNamaKamar(nama string) ([]Kamar, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Kamar, error) {
	kamar, err := s.repository.FindAll()
	return kamar, err

	// return s.repository.FindAll()
}

func (s *service) FindByNamaKamar(nama string) ([]Kamar, error) {
	kamar, err := s.repository.FindByNamaKamar(nama)
	return kamar, err
	// return s.repository.FindAll()
}
