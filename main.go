package main

import (
	"LearningGin/controller"
	"LearningGin/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
    engine := gin.Default()
    // ミドルウェア
    engine.Use(middleware.RecordUaAndTime)
    // CRUD 書籍
    bookEngine := engine.Group("/book")
    {
        v1 := bookEngine.Group("/v1")
        {
            v1.POST("/add", controller.BookAdd)
            v1.GET("/list", controller.BookList)
            v1.PUT("/update", controller.BookUpdate)
            v1.DELETE("/delete", controller.BookDelete)
        }
    }
    engine.Run(":3000")
}

/*User-AgentはHTTPリクエストのヘッダーの一つで、リクエストを送信するクライアント
（ブラウザ、クローラーなど）のアプリケーションの種類、オペレーティングシステム、
バージョンなどの情報を含んでいます。サーバー側では、このUser-Agentを読み取って、
クライアントに応じた最適なレスポンスを提供することができます。たとえば、
特定のブラウザに最適化されたコンテンツを提供したり、モバイルデバイスからの
アクセスに応じたレイアウトを適用したりする場合に利用されます。*/