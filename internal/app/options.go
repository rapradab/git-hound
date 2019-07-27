package app

// Flags stores the program options.
type Flags struct {
	SubdomainFile string
	Debug         bool
	Dig           bool
	APIKeys       bool
}

var flags Flags

// GetFlags is a singleton that returns the program flags.
func GetFlags() *Flags {
	return &flags
}
