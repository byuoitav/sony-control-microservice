package system

import "github.com/byuoitav/hateoas"

// Commands is list of commands
type Commands struct {
	Commands []Command `json:"commands"`
}

// Command is the key-value pair of the command and its ircc equivalent
type Command struct {
	Links   []hateoas.Link `json:"links,omitempty"`
	Command string         `json:"command"`
}

type systemResponse struct {
	Result systemResult `json:"result"`
}

type systemResult struct {
	Bundled bool     `json:"bundled"`
	Type    string   `json:"type"`
	Signals []signal `json:"signals"`
	ID      int      `json:"id"`
}

type signal struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
