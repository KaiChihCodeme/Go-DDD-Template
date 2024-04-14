package models

type GetCafeRequest struct {
	Name string
}

type Cafe struct {
	Uid     uint64
	Name    string
	Address string
}
