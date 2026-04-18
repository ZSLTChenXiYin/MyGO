package httpserv

import (
	"fmt"
	"net/http"

	"github.com/ZSLTChenXiYin/MyGO/configure"
	"github.com/ZSLTChenXiYin/MyGO/logger"
	"github.com/ZSLTChenXiYin/MyGO/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Engine struct {
	eng *gin.Engine
}

func NewEngine(conf configure.Configuration) *Engine {
	// 设置运行模式
	if !conf.Server().Debug() {
		gin.SetMode(gin.ReleaseMode)
	}

	return &Engine{eng: gin.New()}
}

func (e *Engine) Init(std_logger *logger.StdLogger, zap_logger *zap.Logger, allow_headers string, expose_headers string) error {
	e.eng.Use(middleware.StdLogger(std_logger)).
		Use(middleware.ZapLogger(zap_logger)).
		Use(gin.Recovery()).
		Use(middleware.Cors(allow_headers, expose_headers))

	return nil
}

func (e *Engine) Engine() *gin.Engine {
	return e.eng
}

func (e *Engine) HTTPServer(port int) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: e.eng,
	}
}
