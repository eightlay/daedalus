package daedalus

type Data interface {
	GetName() string
	CopyFrom(data Data)
}
