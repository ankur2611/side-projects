package interfaces

import "github.com/gin-gonic/gin"

type IServiceCalls interface {
	HTTPRequest(ctx *gin.Context, method, endpoint string, body interface{}, headers map[string]string) ([]byte, error)
	Get(ctx *gin.Context, endpoint string, headers map[string]string) ([]byte, error)
	Post(ctx *gin.Context, endpoint string, body interface{}, headers map[string]string) ([]byte, error)
	Put(ctx *gin.Context, endpoint string, body interface{}, headers map[string]string) ([]byte, error)
	Delete(ctx *gin.Context, endpoint string, body interface{}, headers map[string]string) ([]byte, error)
	Patch(ctx *gin.Context, endpoint string, body interface{}, headers map[string]string) ([]byte, error)
}
