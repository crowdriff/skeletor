package skeletor

// Conf represents all the runtime flags passed to the binary.
type Conf struct {
	Prod    bool
	Version string
}

var conf *Conf

// SaveConf saves the parsed conf.
func SaveConf(c *Conf) {
	conf = c
}

// GetConf returns a copy of the parsed conf.
func GetConf() Conf {
	return *conf
}

// Version returns the version
func Version() string {
	return conf.Version
}
