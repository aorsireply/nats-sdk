package pkg

type INats interface {
	Connect(url string, options ...Option) (*Conn, error)
	Publish(subj string, data []byte) error
	PublishMsg(m *Msg) error
	Subscribe(subj string, cb MsgHandler) (*Subscription, error)
}
