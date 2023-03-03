package pkg

import (
	"crypto/tls"
	"net"
	"time"
)

type Option func(*Options) error

type Options struct {

	// Url represents a single NATS server url to which the client
	// will be connecting. If the Servers option is also set, it
	// then becomes the first server in the Servers array.
	Url string

	// InProcessServer represents a NATS server running within the
	// same process. If this is set then we will attempt to connect
	// to the server directly rather than using external TCP conns.
	InProcessServer InProcessConnProvider

	// Servers is a configured set of servers which this client
	// will use when attempting to connect.
	Servers []string

	// NoRandomize configures whether we will randomize the
	// server pool.
	NoRandomize bool

	// NoEcho configures whether the server will echo back messages
	// that are sent on this connection if we also have matching subscriptions.
	// Note this is supported on servers >= version 1.2. Proto 1 or greater.
	NoEcho bool

	// Name is an optional name label which will be sent to the server
	// on CONNECT to identify the client.
	Name string

	// Verbose signals the server to send an OK ack for commands
	// successfully processed by the server.
	Verbose bool

	// Pedantic signals the server whether it should be doing further
	// validation of subjects.
	Pedantic bool

	// Secure enables TLS secure connections that skip server
	// verification by default. NOT RECOMMENDED.
	Secure bool

	// TLSConfig is a custom TLS configuration to use for secure
	// transports.
	TLSConfig *tls.Config

	// AllowReconnect enables reconnection logic to be used when we
	// encounter a disconnect from the current server.
	AllowReconnect bool

	// MaxReconnect sets the number of reconnect attempts that will be
	// tried before giving up. If negative, then it will never give up
	// trying to reconnect.
	// Defaults to 60.
	MaxReconnect int

	// ReconnectWait sets the time to backoff after attempting a reconnect
	// to a server that we were already connected to previously.
	// Defaults to 2s.
	ReconnectWait time.Duration

	// CustomReconnectDelayCB is invoked after the library tried every
	// URL in the server list and failed to reconnect. It passes to the
	// user the current number of attempts. This function returns the
	// amount of time the library will sleep before attempting to reconnect
	// again. It is strongly recommended that this value contains some
	// jitter to prevent all connections to attempt reconnecting at the same time.
	CustomReconnectDelayCB ReconnectDelayHandler

	// ReconnectJitter sets the upper bound for a random delay added to
	// ReconnectWait during a reconnect when no TLS is used.
	// Defaults to 100ms.
	ReconnectJitter time.Duration

	// ReconnectJitterTLS sets the upper bound for a random delay added to
	// ReconnectWait during a reconnect when TLS is used.
	// Defaults to 1s.
	ReconnectJitterTLS time.Duration

	// Timeout sets the timeout for a Dial operation on a connection.
	// Defaults to 2s.
	Timeout time.Duration

	// DrainTimeout sets the timeout for a Drain Operation to complete.
	// Defaults to 30s.
	DrainTimeout time.Duration

	// FlusherTimeout is the maximum time to wait for write operations
	// to the underlying connection to complete (including the flusher loop).
	FlusherTimeout time.Duration

	// PingInterval is the period at which the client will be sending ping
	// commands to the server, disabled if 0 or negative.
	// Defaults to 2m.
	PingInterval time.Duration

	// MaxPingsOut is the maximum number of pending ping commands that can
	// be awaiting a response before raising an ErrStaleConnection error.
	// Defaults to 2.
	MaxPingsOut int

	// ClosedCB sets the closed handler that is called when a client will
	// no longer be connected.
	ClosedCB ConnHandler

	// DisconnectedCB sets the disconnected handler that is called
	// whenever the connection is disconnected.
	// Will not be called if DisconnectedErrCB is set
	// DEPRECATED. Use DisconnectedErrCB which passes error that caused
	// the disconnect event.
	DisconnectedCB ConnHandler

	// DisconnectedErrCB sets the disconnected error handler that is called
	// whenever the connection is disconnected.
	// Disconnected error could be nil, for instance when user explicitly closes the connection.
	// DisconnectedCB will not be called if DisconnectedErrCB is set
	DisconnectedErrCB ConnErrHandler

	// ConnectedCB sets the connected handler called when the initial connection
	// is established. It is not invoked on successful reconnects - for reconnections,
	// use ReconnectedCB. ConnectedCB can be used in conjunction with RetryOnFailedConnect
	// to detect whether the initial connect was successful.
	ConnectedCB ConnHandler

	// ReconnectedCB sets the reconnected handler called whenever
	// the connection is successfully reconnected.
	ReconnectedCB ConnHandler

	// DiscoveredServersCB sets the callback that is invoked whenever a new
	// server has joined the cluster.
	DiscoveredServersCB ConnHandler

	// AsyncErrorCB sets the async error handler (e.g. slow consumer errors)
	AsyncErrorCB ErrHandler

	// ReconnectBufSize is the size of the backing bufio during reconnect.
	// Once this has been exhausted publish operations will return an error.
	// Defaults to 8388608 bytes (8MB).
	ReconnectBufSize int

	// SubChanLen is the size of the buffered channel used between the socket
	// Go routine and the message delivery for SyncSubscriptions.
	// NOTE: This does not affect AsyncSubscriptions which are
	// dictated by PendingLimits()
	// Defaults to 65536.
	SubChanLen int

	// UserJWT sets the callback handler that will fetch a user's JWT.
	UserJWT UserJWTHandler

	// Nkey sets the public nkey that will be used to authenticate
	// when connecting to the server. UserJWT and Nkey are mutually exclusive
	// and if defined, UserJWT will take precedence.
	Nkey string

	// SignatureCB designates the function used to sign the nonce
	// presented from the server.
	SignatureCB SignatureHandler

	// User sets the username to be used when connecting to the server.
	User string

	// Password sets the password to be used when connecting to a server.
	Password string

	// Token sets the token to be used when connecting to a server.
	Token string

	// TokenHandler designates the function used to generate the token to be used when connecting to a server.
	TokenHandler AuthTokenHandler

	// Dialer allows a custom net.Dialer when forming connections.
	// DEPRECATED: should use CustomDialer instead.
	Dialer *net.Dialer

	// CustomDialer allows to specify a custom dialer (not necessarily
	// a *net.Dialer).
	CustomDialer CustomDialer

	// UseOldRequestStyle forces the old method of Requests that utilize
	// a new Inbox and a new Subscription for each request.
	UseOldRequestStyle bool

	// NoCallbacksAfterClientClose allows preventing the invocation of
	// callbacks after Close() is called. Client won't receive notifications
	// when Close is invoked by user code. Default is to invoke the callbacks.
	NoCallbacksAfterClientClose bool

	// LameDuckModeHandler sets the callback to invoke when the server notifies
	// the connection that it entered lame duck mode, that is, going to
	// gradually disconnect all its connections before shutting down. This is
	// often used in deployments when upgrading NATS Servers.
	LameDuckModeHandler ConnHandler

	// RetryOnFailedConnect sets the connection in reconnecting state right
	// away if it can't connect to a server in the initial set. The
	// MaxReconnect and ReconnectWait options are used for this process,
	// similarly to when an established connection is disconnected.
	// If a ReconnectHandler is set, it will be invoked on the first
	// successful reconnect attempt (if the initial connect fails),
	// and if a ClosedHandler is set, it will be invoked if
	// it fails to connect (after exhausting the MaxReconnect attempts).
	RetryOnFailedConnect bool

	// For websocket connections, indicates to the server that the connection
	// supports compression. If the server does too, then data will be compressed.
	Compression bool

	// For websocket connections, adds a path to connections url.
	// This is useful when connecting to NATS behind a proxy.
	ProxyPath string

	// InboxPrefix allows the default _INBOX prefix to be customized
	InboxPrefix string

	// IgnoreAuthErrorAbort - if set to true, client opts out of the default connect behavior of aborting
	// subsequent reconnect attempts if server returns the same auth error twice (regardless of reconnect policy).
	IgnoreAuthErrorAbort bool

	// SkipHostLookup skips the DNS lookup for the server hostname.
	SkipHostLookup bool
}

type ReconnectDelayHandler func(attempts int) time.Duration

type InProcessConnProvider interface {
	InProcessConn() (net.Conn, error)
}

type ConnErrHandler func(*Conn, error)

type ErrHandler func(*Conn, *Subscription, error)

type UserJWTHandler func() (string, error)

type SignatureHandler func([]byte) ([]byte, error)

type AuthTokenHandler func() string

type CustomDialer interface {
	Dial(network, address string) (net.Conn, error)
}
