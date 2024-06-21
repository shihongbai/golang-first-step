package client

// 客户端要遵守与服务端一样的契约

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}
