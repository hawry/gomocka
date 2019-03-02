package settings

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

//New returns a new configuration from the given file
func New(f *os.File) (s *Settings, err error) {
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &s)
	return
}

//Port returns the port number from the configuration, or a default port number if not set
func (s *Settings) Port() int {
	if s.ListenPort > 0 {
		return s.ListenPort
	}
	return defaultPort
}

//CreateDefault creates a default configuration and returns the struct representation as well, default filename is example.json
func CreateDefault() (*Settings, error) {
	f, err := os.Create(defaultGeneratedFile)
	if err != nil {
		return nil, err
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
			Mock{
				Path:         "/users/{userid}",
				Method:       "GET",
				ResponseBody: "{\"user\":\"{userid}\"}",
				ResponseCode: 200,
			},
		},
	}
	enc := json.NewEncoder(f)
	enc.SetIndent("", "\t")
	err = enc.Encode(&s)
	return &s, err
}
