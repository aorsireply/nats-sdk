package pkg

type ConnHandler func(*Conn)

type Conn struct {
	// Keep all members for which we use atomic at the beginning of the
	// struct and make sure they are all 64bits (or use padding if necessary).
	// atomic.* functions crash on 32bit machines if operand is not aligned
	// at 64bit. See https://github.com/golang/go/issues/599
	Statistics

	// Opts holds the configuration of the Conn.
	// Modifying the configuration of a running Conn is a race.
	Opts Options
	// contains filtered or unexported fields
}

type Statistics struct {
	InMsgs     uint64
	OutMsgs    uint64
	InBytes    uint64
	OutBytes   uint64
	Reconnects uint64
}
