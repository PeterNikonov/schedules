package sсhedules

const (
	monday    Weekday = "Mon"
	tuesday   Weekday = "Tue"
	wednesday Weekday = "Wed"
	thursday  Weekday = "Thu"
	friday    Weekday = "Fri"
	saturday  Weekday = "Sat"
	sunday    Weekday = "Sun"
)

// Schedule рабочий график
type Schedule struct {
	Id        string    // идентификатор
	Person    Person    // сотрудник
	Specialty Specialty // специализация
	Reception struct {
		Weekdays []*ReceiptWeekday // рабочие дни недели
		Dates    []*ReceiptDate    // рабочие даты
	}
}

// собственные типы для наглядности
type Year int
type Month int
type Day int
type Hour int
type Minute int
type Weekday string

// представление даты
type Date struct {
	Y Year
	M Month
	D Day
}

// When часы и минуты, описывающие ограничение приемного времени
type When struct {
	H Hour
	M Minute
}

// Worktime приемное время
type Worktime struct {
	From     When   // время первого приема
	To       When   // время последнего приема
	Interval Minute // минут на один прием
}

// Person сотрудник
type Person struct {
	Name string
}

// ReceiptDate приемная дата
type ReceiptDate struct {
	Date     Date
	Worktime Worktime
}

// приемный день недели
type ReceiptWeekday struct {
	Weekday  Weekday
	Worktime Worktime
}

// Specialty специальность
type Specialty struct {
	Name string
}

// Vacation отпуск
type Vacation struct {
	Id     string
	Person Person
	From   Date
	To     Date
}
