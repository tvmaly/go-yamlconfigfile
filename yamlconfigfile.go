package yamlconfigfile

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// LoadYAMLFile loads a single level yaml config file to a map
func LoadYAMLFile(file string) (map[string]string, error) {

	var results map[string]string

	yamlfile, err := ioutil.ReadFile(file)

	if err != nil {
		return results, err
	}

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal(yamlfile, &m)

	if err != nil {
		return results, err
	}

	results = make(map[string]string)

	for k, v := range m {

		var key string

		key, ok := k.(string)

		if !ok {
			continue
		}

		switch kt := k.(type) {
		case string:
			key = kt
		default:
			key = fmt.Sprintf("%v", kt)
		}

		switch t := v.(type) {
		case string:
			results[key] = t
		default:
			results[key] = fmt.Sprintf("%v", t)
		}

	}

	return results, nil
}
