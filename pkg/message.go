package pkg

import (
	"bytes"
	"errors"
	"net/http"
)

const (
	hdrLine = "NATS/1.0\r\n"
	crlf    = "\r\n"
)

type Header map[string][]string

type MsgHandler func(msg *Msg)

type Msg struct {
	Subject string
	Reply   string
	Header  Header
	Data    []byte
	Sub     *Subscription
	// contains filtered or unexported fields
}

func (m *Msg) headerBytes() ([]byte, error) {
	var hdr []byte
	if len(m.Header) == 0 {
		return hdr, nil
	}

	var b bytes.Buffer
	_, err := b.WriteString(hdrLine)
	if err != nil {
		return nil, errors.New("nats: message could not decode headers")
	}

	err = http.Header(m.Header).Write(&b)
	if err != nil {
		return nil, errors.New("nats: message could not decode headers")
	}

	_, err = b.WriteString(crlf)
	if err != nil {
		return nil, errors.New("nats: message could not decode headers")
	}

	return b.Bytes(), nil
}
