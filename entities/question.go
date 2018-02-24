package entities

type QuestionRepository interface {
	Store()
	Read()
	Update()
	Delete()
}

type Question struct {
	ID   string
	Text string
}
