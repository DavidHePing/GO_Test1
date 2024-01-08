package weekday

type WeekDay int

const (
	Sunday    WeekDay = 0
	Monday    WeekDay = 1
	Tuesday   WeekDay = 2
	Wednesday WeekDay = 3
	Thursday  WeekDay = 4
	Friday    WeekDay = 5
	Saturday  WeekDay = 6
)

func (w WeekDay) GetDesciption() string {
	switch w {
	case Sunday:
		return "Sunday"
	case Monday:
		return "Monday"
	case Tuesday:
		return "Tuesday"
	case Wednesday:
		return "Wednesday"
	case Thursday:
		return "Thursday"
	case Friday:
		return "Friday"
	case Saturday:
		return "Saturday"
	default:
		return ""
	}
}
