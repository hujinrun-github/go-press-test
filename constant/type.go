package constant

type ClientType string

const (
	CLIENT_TYPE_HTTP = "http"
	CLIENT_TYPE_WS   = "ws"
	CLIENT_TYPE_TCP  = "tcp"
	CLIENT_TYPE_RPC  = "rpc"
)
