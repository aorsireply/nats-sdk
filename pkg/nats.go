package pkg

import (
	"errors"
	"strings"
	"time"
)

const (
	_EMPTY_ = ""
)

// Default Constants
const (
	Version                   = "1.24.0"
	DefaultURL                = "nats://127.0.0.1:4222"
	DefaultPort               = 4222
	DefaultMaxReconnect       = 60
	DefaultReconnectWait      = 2 * time.Second
	DefaultReconnectJitter    = 100 * time.Millisecond
	DefaultReconnectJitterTLS = time.Second
	DefaultTimeout            = 2 * time.Second
	DefaultPingInterval       = 2 * time.Minute
	DefaultMaxPingOut         = 2
	DefaultMaxChanLen         = 64 * 1024       // 64k
	DefaultReconnectBufSize   = 8 * 1024 * 1024 // 8MB
	RequestChanLen            = 8
	DefaultDrainTimeout       = 30 * time.Second
	LangString                = "go"
)

type NatsMessageBus struct {
}

func (n NatsMessageBus) Connect(url string, options ...Option) (*Conn, error) {
	opts := GetDefaultOptions()
	opts.Servers = processUrlString(url)
	for _, opt := range options {
		if opt != nil {
			if err := opt(&opts); err != nil {
				return nil, err
			}
		}
	}
	return opts.Connect()
}

// GetDefaultOptions returns default configuration options for the client.
func GetDefaultOptions() Options {
	return Options{
		AllowReconnect:     true,
		MaxReconnect:       DefaultMaxReconnect,
		ReconnectWait:      DefaultReconnectWait,
		ReconnectJitter:    DefaultReconnectJitter,
		ReconnectJitterTLS: DefaultReconnectJitterTLS,
		Timeout:            DefaultTimeout,
		PingInterval:       DefaultPingInterval,
		MaxPingsOut:        DefaultMaxPingOut,
		SubChanLen:         DefaultMaxChanLen,
		ReconnectBufSize:   DefaultReconnectBufSize,
		DrainTimeout:       DefaultDrainTimeout,
	}
}

// Publish publishes the data argument to the given subject. The data
// argument is left untouched and needs to be correctly interpreted on
// the receiver.
func (nc *Conn) Publish(subj string, data []byte) error {
	return nc.publish(subj, _EMPTY_, nil, data)
}

// PublishMsg publishes the Msg structure, which includes the
// Subject, an optional Reply and an optional Data field.
func (nc *Conn) PublishMsg(m *Msg) error {
	if m == nil {
		return errors.New("nats: invalid message or message nil")
	}
	hdr, err := m.headerBytes()
	if err != nil {
		return err
	}
	return nc.publish(m.Subject, m.Reply, hdr, m.Data)
}

// Subscribe will express interest in the given subject. The subject
// can have wildcards.
// There are two type of wildcards: * for partial, and > for full.
// A subscription on subject time.*.east would receive messages sent to time.us.east and time.eu.east.
// A subscription on subject time.us.> would receive messages sent to
// time.us.east and time.us.east.atlanta, while time.us.* would only match time.us.east
// since it can't match more than one token.
// Messages will be delivered to the associated MsgHandler.
func (nc *Conn) Subscribe(subj string, cb MsgHandler) (*Subscription, error) {
	return nc.subscribe(subj, _EMPTY_, cb, nil, false, nil)
}

func (nc *Conn) publish(subj, reply string, hdr, data []byte) error {
	//TODO
	return nil
}

// subscribe is the internal subscribe function that indicates interest in a subject.
func (nc *Conn) subscribe(subj, queue string, cb MsgHandler, ch chan *Msg, isSync bool, js *jsSub) (*Subscription, error) {
	//TODO
	return &Subscription{}, nil
}

// Connect will attempt to connect to a NATS server with multiple options.
func (o Options) Connect() (*Conn, error) {
	//TODO
	return &Conn{}, nil
}

// Process the url string argument to Connect.
// Return an array of urls, even if only one.
func processUrlString(url string) []string {
	urls := strings.Split(url, ",")
	var j int
	for _, s := range urls {
		u := strings.TrimSpace(s)
		if len(u) > 0 {
			urls[j] = u
			j++
		}
	}
	return urls[:j]
}
