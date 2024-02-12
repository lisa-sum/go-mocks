package main

type db struct{}

// DB is fake database interface.
//
//go:generate mockery --name DB
type DB interface {
	FetchMessage(lang string) (string, error)
	FetchDefaultMessage() (string, error)
}

func main() {

}
