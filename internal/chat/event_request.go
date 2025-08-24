package chat

type Event struct {
	ID        ID        `json:"_id"`
	Entry     []Entry   `json:"entry"`
	HTTP      HTTP      `json:"http"`
	Object    string    `json:"object"`
}

type ID struct {
	Oid string `json:"$oid"`
}

type Headers struct {
	Accept           string `json:"accept"`
	ContentType      string `json:"content-type"`
	XHubSignature    string `json:"x-hub-signature"`
	UserAgent        string `json:"user-agent"`
	Host             string `json:"host"`
	CfConnectingIP   string `json:"cf-connecting-ip"`
	XForwardedProto  string `json:"x-forwarded-proto"`
	CfRay            string `json:"cf-ray"`
	CdnLoop          string `json:"cdn-loop"`
	XForwardedFor    string `json:"x-forwarded-for"`
	XRequestID       string `json:"x-request-id"`
	XHubSignature256 string `json:"x-hub-signature-256"`
	CfVisitor        string `json:"cf-visitor"`
	CfIpcountry      string `json:"cf-ipcountry"`
	AcceptEncoding   string `json:"accept-encoding"`
}
type HTTP struct {
	Path    string  `json:"path"`
	Method  string  `json:"method"`
	Headers Headers `json:"headers"`
}
