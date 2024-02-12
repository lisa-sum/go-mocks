package db

// DB is fake database interface.
//
//go:generate mockery --name DB
type DB interface {
	FetchMessage(lang string) (string, error)
	FetchDefaultMessage() (string, error)
}
type greeter struct {
	database DB
	lang     string
}

func (g greeter) Greet() string {
	msg, _ := g.database.FetchMessage(g.lang) // call database to get the message based on the lang
	return "Message is: " + msg
}

func (g greeter) GreetInDefaultMsg() string {
	msg, _ := g.database.FetchDefaultMessage() // call database to get the default message
	return "Message is: " + msg
}

// GreeterService is service to greet your friends.
//
//go:generate mockery --name GreeterService
type GreeterService interface {
	Greet() string
	GreetInDefaultMsg() string
}
