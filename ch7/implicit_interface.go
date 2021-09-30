package main

type Logic1 interface {
	Process(data string) string
}

// LogicProvider implicitly implements Logic Interface
type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
	return ""
}

// only the Client (Caller) knows about the interface
type Client struct {
	L Logic1
}

// Program method is on the client, just to be able to call the Process on Logic
func (c Client) Program() {
	data := "s" // get data from somewhere
	c.L.Process(data)
}

func main() {
	c := Client{
		L: LogicProvider{}, // new LogicProvider can be provided in the future
	}
	c.Program()
}