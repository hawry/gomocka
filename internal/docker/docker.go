package docker

import (
	"fmt"
	"os"
	"time"
)

var template = `
# generated %s
FROM scratch
ADD ./gomocka /
ADD %s /
ENTRYPOINT ["/gomocka"]
CMD ["--config", "%s"]
`

//Create generates a new dockerfile with the given config file as the settings file, this will overwrite any already existing dockerfiles in the default location ("./generated/Dockerfile")
func Create(configFile string) error {
	f, err := os.Create("dockerfile")
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf(template, time.Now().Format("2006-01-02 15:04:05"), configFile, configFile))
	if err != nil {
		return err
	}
	return nil
}
