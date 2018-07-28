package main

// Meta is a HTTP header to be set in response to GET object
type Meta struct {
	// Name is a header name
	Name string `json:"name"`
	// Value is a heder value
	Value string `json:"value"`
}

// ObjectContainer is a wrapper to store object data and meta in Redis
type ObjectContainer struct {
	Meta []Meta `json:"meta"`
	// Value is object bytes
	Value []byte `json:"value"`
}
