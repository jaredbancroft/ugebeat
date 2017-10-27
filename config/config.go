// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

// Config - configuration options for beat
type Config struct {
	Period  time.Duration `config:"period"`
	Ugeroot string        `config:"ugeroot"`
	Ugecell string        `config:"ugecell"`
}

// DefaultConfig - configurations if no configuration provided
var DefaultConfig = Config{
	Period:  1 * time.Second,
	Ugeroot: ".",
	Ugecell: "default",
}
