package newsfeed

type Store interface {
	Get() []Item
	Set(Item) bool
}
