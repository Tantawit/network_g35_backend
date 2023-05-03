package subscriber

type Subscriber interface {
	Listen()
	Close()
}
