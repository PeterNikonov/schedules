package sсhedules

// scheduleStorage хранилище данных о графиках
type scheduleStorage interface {
	// save сохранить график
	save(sch *Schedule) error
	// all получить все графики
	all() ([]*Schedule, error)
	// find находит график
	find(sch *Schedule) *Schedule
	// update обновляет график
	update(target *Schedule, actual *Schedule) error
}

type inMemoryScheduleStorage struct {
	data []*Schedule
}

func newInMemoryScheduleStorage() *inMemoryScheduleStorage {
	return &inMemoryScheduleStorage{}
}

func (s *inMemoryScheduleStorage) save(sch *Schedule) error {
	existed := s.find(sch)
	if existed != nil {
		s.update(existed, sch)
		return nil
	}
	s.data = append(s.data, sch)
	return nil
}

func (s *inMemoryScheduleStorage) all() ([]*Schedule, error) {
	return s.data, nil
}

func (s *inMemoryScheduleStorage) find(sch *Schedule) *Schedule {
	for _, item := range s.data {
		if item.Id == sch.Id {
			return item
		}
	}
	return nil
}

func (s *inMemoryScheduleStorage) update(target *Schedule, actual *Schedule) error {
	target.Person = actual.Person
	target.Specialty = actual.Specialty
	target.Reception = actual.Reception
	return nil
}
