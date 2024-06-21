package server

import "errors"

// Go RPC的函数只有符合下面的条件才能被远程访问，不然会被忽略，详细的要求如下：
// 1. 函数必须是导出的(首字母大写)
// 2. 必须有两个导出类型的参数
// 3. 第一个参数是接受的参数，第二个参数是返回给客户端的参数，第二个参数必须是指针类型
// 4. 返回值必须要有个error

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
