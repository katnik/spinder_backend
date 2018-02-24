package entities

type PlayerRepository interface {
	Store()
	Read()
	Update()
	Delete()
}

type Player struct {
	ID   int
	Name string
	Team string
}
