package mock

import (
	"net/http"

	"github.com/agnivo988/MockForge/backend/internal/scenario"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
)

// NOTE: Later this will be loaded from JSON/YAML
var rules []scenario.Rule

func RegisterMockRoutes(r *gin.Engine, spec *openapi3.T) {
	for path, pathItem := range spec.Paths.Map() {
		registerOperation(r, http.MethodGet, path, pathItem.Get)
		registerOperation(r, http.MethodPost, path, pathItem.Post)
		registerOperation(r, http.MethodPut, path, pathItem.Put)
		registerOperation(r, http.MethodDelete, path, pathItem.Delete)
	}
}

func registerOperation(
	r *gin.Engine,
	method string,
	path string,
	op *openapi3.Operation,
) {
	if op == nil {
		return
	}

	r.Handle(method, path, func(ctx *gin.Context) {

		//Scenario rule matching (FIRST)
		if rule := scenario.MatchRule(ctx, rules); rule != nil {
			ctx.JSON(rule.StatusCode, rule.Response)
			return
		}

		//OpenAPI example response (FALLBACK)
		for code, resp := range op.Responses.Map() {
			if resp.Value == nil || resp.Value.Content == nil {
				continue
			}

			for _, media := range resp.Value.Content {
				if media.Example != nil {
					ctx.JSON(parseStatus(code), media.Example)
					return
				}
			}
		}

		//Default response
		ctx.JSON(http.StatusOK, gin.H{
			"message": "mock response",
		})
	})
}

func parseStatus(code string) int {
	switch code {
	case "200":
		return 200
	case "201":
		return 201
	case "400":
		return 400
	case "404":
		return 404
	case "500":
		return 500
	default:
		return 200
	}
}
