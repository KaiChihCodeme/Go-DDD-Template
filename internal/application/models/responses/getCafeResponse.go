package responses

type GetCafeResponse struct {
	Uid     uint64 `json:"uid"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
