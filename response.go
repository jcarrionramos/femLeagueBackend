package main

type Response struct {
	Status int         `json:"-"`
	Data   interface{} `json:"data,omitempty"`
	Meta   interface{} `json:"meta,omitempty"`
}
