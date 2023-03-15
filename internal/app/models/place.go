package models

type Place struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Rating  int    `json:"rating"`
}
