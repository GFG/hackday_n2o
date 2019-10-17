package srv

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type (
	downloadHandler struct {

	}
)

func DownloadHandler(server *gin.Engine) {
	u := &downloadHandler{}
	server.GET("/download", u.Download)
}

func (u *downloadHandler) Download(c *gin.Context) {
	fmt.Println("Download working")
}


