package ircc

type Capability struct {
	ID     int
	Result []poots
}

type poots struct {
	Overview overview  `json:",omitempty"`
	Command  []command `json:",omitempty"`
}

type overview struct {
	Bundled bool   `json:"bundled,omitempty"`
	Type    string `json:"type,omitempty"`
}

type command struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
