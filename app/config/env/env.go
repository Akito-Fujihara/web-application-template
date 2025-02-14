package env

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(
	MysqlConfig,
	CacheConfig,
	CORSConfig,
)
