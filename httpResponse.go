package main

// Quotes Quotes contains the quote and other info
type Quotes struct {
	Author     string   `json:"author"`
	Background string   `json:"background"`
	Category   string   `json:"category"`
	Date       string   `json:"date"`
	ID         string   `json:"id"`
	Length     int      `json:"length"`
	Permalink  string   `json:"permalink"`
	Quote      string   `json:"quote"`
	Tags       []string `json:"tags"`
	Title      string   `json:"title"`
}

// Contents Contents needs to be here to contain quotes, :|
type Contents struct {
	Quotes []Quotes `json:"quotes"`
}

// QuoteResponse idk why i need it... oh yeah  just contains everything
type QuoteResponse struct {
	Contents Contents `json:"contents"`
	Success  []string `json:"success"`
}

// SomeError error struct for error messages
type SomeError struct {
	Error TheActualError `json:"error"`
}

// TheActualError the error content struct
type TheActualError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
