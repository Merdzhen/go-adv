package main


// https://pkg.go.dev/io#Reader
type Reader interface {
	Read(p []byte) (n int, err error)
}

// Зачем такая сигнатура:
// Внешний буфер передается, чтобы управлять памятью эффективно.
// Избегается создание внутреннего объекта, который позже нужно будет чистить, 
// что связано с особенностями работы стека и кучи.

type Reader2 interface {
	Read() (p []byte, err error)
	//в этом случае p ушел бы на кучу
}
