package sсhedules

// vacationStorage хранилище данных о графиках
type vacationStorage interface {
	// save сохранить график
	save(v *Vacation) error
	// all получить все графики
	all() ([]*Vacation, error)
	// find находит график
	find(v *Vacation) *Vacation
	// update обновляет график
	update(target *Vacation, actual *Vacation) error
}

type inMemoryVacationStorage struct {
	data []*Vacation
}

func newInMemoryVacationStorage() *inMemoryVacationStorage {
	return &inMemoryVacationStorage{}
}

func (s *inMemoryVacationStorage) save(v *Vacation) error {
	existed := s.find(v)
	if existed != nil {
		s.update(existed, v)
		return nil
	}
	s.data = append(s.data, v)
	return nil
}

func (s *inMemoryVacationStorage) all() ([]*Vacation, error) {
	return s.data, nil
}

func (s *inMemoryVacationStorage) find(v *Vacation) *Vacation {
	for _, item := range s.data {
		if item.Id == v.Id {
			return item
		}
	}
	return nil
}

func (s *inMemoryVacationStorage) update(target *Vacation, actual *Vacation) error {
	target.Person = actual.Person
	target.From = actual.From
	target.To = actual.To
	return nil
}
