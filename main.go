package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
    engine:= gin.Default()
    ua := ""
     // ミドルウェアを使用
    engine.Use(func(c *gin.Context) {
        ua = c.GetHeader("User-Agent")
        c.Next()
    })
    engine.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message":    "hello world",
            "User-Agent": ua,
        })
    })
    engine.Run(":3000")
}

/*User-AgentはHTTPリクエストのヘッダーの一つで、リクエストを送信するクライアント
（ブラウザ、クローラーなど）のアプリケーションの種類、オペレーティングシステム、
バージョンなどの情報を含んでいます。サーバー側では、このUser-Agentを読み取って、
クライアントに応じた最適なレスポンスを提供することができます。たとえば、
特定のブラウザに最適化されたコンテンツを提供したり、モバイルデバイスからの
アクセスに応じたレイアウトを適用したりする場合に利用されます。*/