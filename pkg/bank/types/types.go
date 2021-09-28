package types

type Money int64

type Currency string

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

type Card struct{
	ID int
	PAN string
	Balance Money
	MinBalance Money
	Currency Currency
	Color string
	Name string
	Active bool
}