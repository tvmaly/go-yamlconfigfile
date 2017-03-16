# go-yamlconfigfile
yaml config file reader in Go

I wanted a simple way to load a single level yaml file with config values into my project as a map[string]string.

I did not want to put API keys in code, so the plan is to keep them in a config file that is kept out of source control

YAML is a pretty easy format to read, so I chose this format for the config.

For example, say you have a file test.yaml that contains:

--------------------------------------
foo: 12345
baz: abc123
bar: "hello world"
pi: 3.1415926
--------------------------------------


yaml, err := LoadYAMLFile("test.yaml")


yaml is a map[string]string

if you were to print this with 

fmt.Printf("%v\n", yaml)

You would see

map[bar:hello world pi:3.1415926 foo:12345 baz:abc123]

