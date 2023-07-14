package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	router *gin.Engine
	port   string
}

type HttpServerConfig struct {
	Port string `json:"port" default:":8080"`
	Mode string `json:"mode" default:"debug" example:"debug/release/test"`
}

func NewHttpServer(conf *HttpServerConfig) HttpServer {
	gin.SetMode(conf.Mode)
	router := gin.New()

	// 写入recover重新拉起
	router.Use(Recover)

	// 写入跨域中间件
	router.Use(Cors())

	// debug模式和test模式开启日志
	if conf.Mode == "debug" || conf.Mode == "test" {
		router.Use(Logger())
	}

	// 端口
	if conf.Port == "" {
		conf.Port = ":8080"
	}

	return HttpServer{
		router: router,
		port:   conf.Port,
	}
}

// 路由
func (h *HttpServer) Router() *gin.Engine {
	return h.router
}

// 初始化
// 在注册中间件及路由后运行初始化
// ! 请在主函数的结束位置完成http服务初始化，否则将导致初始化后的函数失效
func (h *HttpServer) Init() {
	srv := &http.Server{
		Addr:    h.port,
		Handler: h.router,
	}

	LogInfo("Server", "Listen:%s", h.port)

	go func() {
		// 开启goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			LogFatal("Server", "Fatal: %v", err)
		}
		LogInfo("Server", "Server Start...")
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1)
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 阻塞在此，当接收到上述两种信号时才会往下执行
	<-quit
	LogInfo("Server", "Shutting Down Server...")
	// 创建一个10秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 10秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过10秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		LogFatal("Server", "Server Forced To Shutdown: %v", err)
	}

	LogInfo("Server", "Server Exiting")
	defer LogInfo("Server", "Server Closed")
}
