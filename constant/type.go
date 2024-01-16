package constant

type ClientType string

const (
	CLIENT_TYPE_HTTP = "http"
	CLIENT_TYPE_WS   = "ws"
	CLIENT_TYPE_TCP  = "tcp"
	CLIENT_TYPE_RPC  = "rpc"
)

type SOURCE_TYPE string

const (
	ST_FILE   SOURCE_TYPE = "file"
	ST_ORIGIN SOURCE_TYPE = "origin"
)

type READ_METHOD string

const (
	RM_ALL  READ_METHOD = "all"  // read all one time
	RM_LINE READ_METHOD = "line" // read one line one time
)
