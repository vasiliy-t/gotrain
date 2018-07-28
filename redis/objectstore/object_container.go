package main

type Meta struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ObjectContainer struct {
	Meta  []Meta `json:"meta"`
	Value []byte `json:"value"`
}
