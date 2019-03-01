package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	defaultPort          = 8080
	defaultGeneratedFile = "example.json"
)

//Settings describe the settings.json file contents
type Settings struct {
	ListenPort int    `json:"port"`
	Mocks      []Mock `json:"mocks"`
}

// Mock describes what path that should return what status & body
type Mock struct {
	Path         string            `json:"path"`
	Method       string            `json:"method"`
	ResponseCode int               `json:"response_code"`
	ResponseBody string            `json:"response_body"`
	Headers      map[string]string `json:"headers"`
}

// NewSettings returns a new settings file containing the endpoints to mock
func NewSettings(f *os.File) (s *Settings, err error) {
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &s)
	return
}

// Port returns the set port number or returns the default port if none is set
func (s *Settings) Port() int {
	if s.ListenPort > 0 {
		return s.ListenPort
	}
	return defaultPort
}

func createDefault() error {
	f, err := os.Create(defaultGeneratedFile)
	if err != nil {
		return err
	}
	s := Settings{
		ListenPort: 8080,
		Mocks: []Mock{
			Mock{
				Path:         "/",
				Method:       "GET",
				ResponseBody: "{\"hello\":\"world\"}",
				ResponseCode: 200,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			},
		},
	}
	enc := json.NewEncoder(f)
	enc.SetIndent("", "\t")
	return enc.Encode(&s)
}
