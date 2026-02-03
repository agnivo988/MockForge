package parser

import (
	"context"
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

func Validate(spec *openapi3.T) error {
	if err := spec.Validate(context.Background()); err != nil {
		return fmt.Errorf("spec validation failed: %w", err)
	}
	return nil
}
