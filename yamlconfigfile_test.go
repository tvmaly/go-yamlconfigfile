package yamlconfigfile

import (
	"fmt"
	"testing"
)

func TestLoadYAMLFile(t *testing.T) {

	yaml, err := LoadYAMLFile("test.yaml")

	if err != nil {
		t.Fatalf("Error loading yaml test file: %s", err)
	} else {
		fmt.Printf("Loaded\n: %v\n", yaml)
	}

}
