package store

import "github.com/getkin/kin-openapi/openapi3"

var LoadedSpec *openapi3.T

func SetSpec(spec *openapi3.T){
	LoadedSpec = spec
}

func GetSpec() *openapi3.T {
	return LoadedSpec
}