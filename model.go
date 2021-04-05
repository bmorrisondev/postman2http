package main

type PostmanCollection struct {
	Info  Info   `json:"info"`
	Items []Item `json:"item"`
}

type Info struct {
	PostmanId string `json:"_postman_id"`
	Name      string `json:"name"`
	Schema    string `json:"schema"`
}

type Item struct {
	Name    string  `json:"name"`
	Request Request `json:"request"`
}

type Request struct {
	Method  string   `json:"method"`
	Headers []Header `json:"header"`
	Url     Url      `json:"url"`
	Body    Body     `json:"body"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Url struct {
	Raw string `json:"raw"`
}

type Body struct {
	Raw string `json:"raw"`
}
