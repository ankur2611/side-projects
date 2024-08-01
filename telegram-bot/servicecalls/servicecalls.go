package servicecalls

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceCall struct {
}

func NewServiceCall(ctx context.Context) ServiceCall {
	return ServiceCall{}
}

func (sc ServiceCall) HTTPRequest(ctx *gin.Context, method, endpoint string, body interface{}, headers map[string]string) ([]byte, error) {

	client := &http.Client{}
	requestBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, endpoint, bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, err
	}

	for key, header := range headers {
		request.Header.Set(key, header)
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	statusCode := response.StatusCode

	if err != nil {
		return nil, err
	}

	bytesResponse, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	if statusCode < 200 || statusCode >= 400 {
		return bytesResponse, fmt.Errorf("HTTP request failed with status code %s", string(bytesResponse))
	}

	return bytesResponse, nil
}

func (sc ServiceCall) Get(ctx *gin.Context, endpoint string, headers map[string]string) ([]byte, error) {

	response, err := sc.HTTPRequest(ctx, http.MethodGet, endpoint, nil, headers)

	if err != nil {

		return response, err
	}

	return response, nil
}

func (sc ServiceCall) Post(ctx *gin.Context, endpoint string, body interface{}, headers map[string]string) ([]byte, error) {

	response, err := sc.HTTPRequest(ctx, http.MethodPost, endpoint, body, headers)

	if err != nil {
		return response, err
	}

	return response, nil
}

func (sc ServiceCall) Put(ctx *gin.Context, endpoint string, body interface{}, headers map[string]string) ([]byte, error) {

	response, err := sc.HTTPRequest(ctx, http.MethodPut, endpoint, nil, headers)

	if err != nil {
		return response, err
	}

	return response, nil
}

func (sc ServiceCall) Delete(ctx *gin.Context, endpoint string, body interface{}, headers map[string]string) ([]byte, error) {

	response, err := sc.HTTPRequest(ctx, http.MethodDelete, endpoint, nil, headers)

	if err != nil {
		return response, err
	}

	return response, nil
}

func (sc ServiceCall) Patch(ctx *gin.Context, endpoint string, body interface{}, headers map[string]string) ([]byte, error) {

	response, err := sc.HTTPRequest(ctx, http.MethodPatch, endpoint, nil, headers)

	if err != nil {
		return response, err
	}

	return response, nil
}
