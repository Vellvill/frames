package config

type errCode uint8

const (
	ErrNotFound errCode = iota
)

type ConfigErr struct {
	message    string
	statusCode errCode
}

func (c ConfigErr) Error() string {
	return c.message
}

func newConfigError(m string, code errCode) ConfigErr {
	return ConfigErr{
		message:    m,
		statusCode: code,
	}
}
