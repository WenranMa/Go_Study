package main

type AddParam struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type AddResult struct {
	Code int `json:"code"`
	Data int `json:"data"`
}
