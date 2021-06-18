package tests

import (
	"github.com/stretchr/testify/assert"
	sсhedules "schedules/internal/schedules"
	"testing"
)

func TestShedules(t *testing.T) {
	s := sсhedules.NewStore()
	all, err := s.AllSchedules()
	assert.Nil(t, err)
	assert.Equal(t, len(all), 0)
	err = s.SaveSchedule(&sсhedules.Schedule{Id: "a", Person: sсhedules.Person{Name: "Anton"}})
	assert.Nil(t, err)
	all, err = s.AllSchedules()
	assert.Nil(t, err)
	assert.Equal(t, len(all), 1)
	sch := all[0]
	assert.Equal(t, sch.Id, "a")
	assert.Equal(t, sch.Person.Name, "Anton")
	err = s.SaveSchedule(&sсhedules.Schedule{Id: "a", Person: sсhedules.Person{Name: "Nikita"}})
	assert.Nil(t, err)
	all, err = s.AllSchedules()
	assert.Nil(t, err)
	assert.Equal(t, len(all), 1)
	sch = all[0]
	assert.Equal(t, sch.Id, "a")
	assert.Equal(t, sch.Person.Name, "Nikita")
}

func TestVacations(t *testing.T) {
	s := sсhedules.NewStore()
	all, err := s.AllVacations()
	assert.Nil(t, err)
	assert.Equal(t, len(all), 0)
	err = s.SaveVacation(&sсhedules.Vacation{})
	assert.Nil(t, err)
	all, err = s.AllVacations()
	assert.Nil(t, err)
	assert.Equal(t, len(all), 1)
}
