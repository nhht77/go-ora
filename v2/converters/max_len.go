package converters

import "sync"

var (
	MAX_LEN_VARCHAR2  int = 0x7FFF
	MAX_LEN_NVARCHAR2     = 0x7FFF
	MAX_LEN_RAW           = 0x7FFF
	MAX_LEN_NUMBER        = 0x16
	MAX_LEN_DATE          = 0xB
	MAX_LEN_TIMESTAMP     = 0xB
)

var Mutex sync.RWMutex
