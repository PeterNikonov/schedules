package sсhedules

// storages хранилища, используемые фасадом
type storages struct {
	schedule scheduleStorage
	vacation vacationStorage
}

// store работа с хранилищами графиков и отпусков
type store struct {
	storages *storages
}

// SaveSchedule сохранить график
func (p *store) SaveSchedule(s *Schedule) error {
	return p.storages.schedule.save(s)
}

// SaveVacation сохранить отпуск
func (p *store) SaveVacation(v *Vacation) error {
	return p.storages.vacation.save(v)
}

// AllSchedules получить все графики
func (p *store) AllSchedules() ([]*Schedule, error) {
	return p.storages.schedule.all()
}

// AllVacations получить все отпуска
func (p *store) AllVacations() ([]*Vacation, error) {
	return p.storages.vacation.all()
}

// NewStore Constructor
func NewStore() *store {
	return &store{storages: &storages{
		schedule: newInMemoryScheduleStorage(),
		vacation: newInMemoryVacationStorage(),
	}}
}
