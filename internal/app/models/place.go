package models

type Place struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Rating      int    `json:"rating"`
}
