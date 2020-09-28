package types

type Money int

type Category string

type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status Status
}

type Status string

const (
	StatusOk Status = "OK"
	StatusFail Status = "FAIL"
	StatusProgress Status = "INPROGRESS"
)