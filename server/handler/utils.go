package handler

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func respondJSON(ctx echo.Context, statusCode int, data interface{}) error {

	var response []byte
	var err error
	if ctx.Echo().Debug {
		response, err = json.MarshalIndent(data, "", "  ")
	} else {
		response, err = json.Marshal(data)
	}
	if err != nil {
		return err
	}

	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return respond(ctx, statusCode, response)
}

func respond(ctx echo.Context, statusCode int, data []byte) error {

	writeHeader(ctx, statusCode)
	_, err := ctx.Response().Write(data)
	if err != nil {
		ctx.Logger().Errorf("Error writing response body: %s", err.Error())
	}

	return nil
}

func respondGenericError(ctx echo.Context, err error) error {
	return respondJSON(ctx, http.StatusInternalServerError, err)
}

func writeHeader(ctx echo.Context, statusCode int) {
	res := ctx.Response()
	res.WriteHeader(statusCode)
}
