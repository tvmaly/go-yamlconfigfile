# go-yamlconfigfile
a simple yaml config file reader in Go that converts a single level of yaml values into a map[string]string

I wanted a simple way to load a single level yaml file with config values for API keys etc into my project.

I did not want to put API keys in code, so the plan is to keep them in a config file that is kept out of source control

YAML is a pretty easy format to read, so I chose this format for the config.

For example, say you have a file test.yaml that contains:

>
> foo: 12345
>
> baz: abc123
>
> bar: "hello world"
>
> pi: 3.1415926
>


yaml, err := LoadYAMLFile("test.yaml")


yaml is a map[string]string

if you were to print this with 

fmt.Printf("%v\n", yaml)

You would see

map[bar:hello world pi:3.1415926 foo:12345 baz:abc123]

-------------------------

go test 

Type
: map[string]string
Loaded Values
: map[pi:3.1415926 foo:12345 baz:abc123 bar:hello world]
PASS
ok  	go-yamlconfigfile	0.007s

