package parser

import (
	"context"
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

func LoadSpec(path string) (*openapi3.T, error) {
	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true

	spec, err := loader.LoadFromFile(path)
	if(err != nil){
		return nil, err
	}

	if err := spec.Validate(context.Background());
	err != nil{
		return nil, fmt.Errorf("Invalid OpenAPI spec: %w",err)
	}

	return spec, nil
}