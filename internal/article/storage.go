package article

type storage interface {
	GetAll() *[]Article
	GetByID(id int) *Article
	GetFlow(flow string) *[]Article
}
