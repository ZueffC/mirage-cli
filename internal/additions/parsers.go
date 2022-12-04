package additions

import (
	"github.com/BurntSushi/toml"
)

func ParseNodes(path string) []string {
	var nodes []string
	var nodeStruct struct {
		Nodes []string
	}

	_, err := toml.DecodeFile("nodes.toml", &nodeStruct)
	if err != nil {
		panic(err)
	}

	nodes = append(nodes, nodeStruct.Nodes...)

	return nodes
}

func ParseDependencies(path string) []string {
	var deps []string
	var dependencies struct {
		Dependencies []string
	}

	_, err := toml.DecodeFile(path, &dependencies)
	if err != nil {
		panic(err)
	}

	deps = append(deps, dependencies.Dependencies...)

	return deps
}

func ParseRunScript(path string) []string {
	var scripts []string
	var scriptStruct struct {
		AfterInstallScripts []string
	}

	_, err := toml.DecodeFile(path, &scriptStruct)
	if err != nil {
		panic(err)
	}

	scripts = append(scripts, scriptStruct.AfterInstallScripts...)

	return scripts
}
