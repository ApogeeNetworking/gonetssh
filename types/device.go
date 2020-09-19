package types

// Device ...
type Device interface {
	Connect(retries int) error
	Disconnect()
	SendCmd(cmd string) (string, error)
	// ExecEnable()
}
