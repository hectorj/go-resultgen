package tests

//go:generate go run ../main.go Dummy --tags=!strict
//go:generate go run ../main.go Dummy --strict --tags=strict --output=dummy_result_strict.go
type Dummy struct {
	ID int
}
