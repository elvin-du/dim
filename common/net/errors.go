package libnet

import (
	"errors"
)

// errors
var (
	SendToClosedError          = errors.New("Send to closed session")
	BlockingError              = errors.New("Blocking happened")
	PacketTooLargeError        = errors.New("Packet too large")
	SessionDuplicateStartError = errors.New("Session duplicate start")
	ServerDuplicateStartError  = errors.New("Server duplicate start")
	SyncSendTimeoutError       = errors.New("Sync send timeout")
	TimeoutBlockingError       = errors.New("Timeout happens when blocking")
	CloseBlockingError         = errors.New("Session closed after blocking")
	DiscardBlockingError       = errors.New("Message is discarded after blocking")
	PacketHeadLenError         = errors.New("Invalid packet head length")
)
