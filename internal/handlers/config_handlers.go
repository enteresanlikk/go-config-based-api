package handlers

import (
	"go-config-based-api/internal/config"
	customHttp "go-config-based-api/internal/http"

	"github.com/valyala/fasthttp"
)

// GetAllConfigs handles the request to get all configs
func GetAllConfigs(ctx *fasthttp.RequestCtx) {
	configLoader := config.GetInstance()
	configs := configLoader.GetAllConfigs()

	if err := customHttp.JsonResponse(ctx, configs); err != nil {
		customHttp.JsonResponseWithStatus(ctx, fasthttp.StatusInternalServerError, map[string]string{
			"error": "Failed to encode response",
		})
	}
}

// GetConfigByID handles the request to get a specific config by ID
func GetConfigByID(ctx *fasthttp.RequestCtx) {
	configLoader := config.GetInstance()
	configID := ctx.UserValue("id").(string)

	if config, exists := configLoader.GetConfig(configID); exists {
		if err := customHttp.JsonResponse(ctx, config); err != nil {
			customHttp.JsonResponseWithStatus(ctx, fasthttp.StatusInternalServerError, map[string]string{
				"error": "Failed to encode response",
			})
		}
	} else {
		customHttp.JsonResponseWithStatus(ctx, fasthttp.StatusNotFound, map[string]string{
			"error": "Config not found",
		})
	}
}
