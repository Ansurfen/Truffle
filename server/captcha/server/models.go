package server

type ProofRequest struct {
	Id  string `json:"id" form:"id"`
	Val string `json:"val" form:"val"`
}
