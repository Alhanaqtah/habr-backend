package article

type storage interface {
	GetAll() *[]Article
	GetFlow(flow string) *[]Article
}
