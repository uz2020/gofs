package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"github.com/uz2020/gofs/internal/cmd/controllers"
	log "unknwon.dev/clog/v2"
)

var Web = cli.Command{
	Name:        "web",
	Usage:       "Start web server",
	Description: "start gofs",
	Action:      runWeb,
	Flags: []cli.Flag{
		stringFlag("port, p", "3000", "Temporary port number to prevent conflict"),
		stringFlag("config, c", "", "Custom configuration file path"),
	},
}

func runWeb(c *cli.Context) error {
	port := c.String("port")
	addr := fmt.Sprintf(":%s", port)
	log.Info("start gofs on port: %s", port)

	router := gin.Default()
	g := router.Group("/file")
	{
		g.POST("/:id", controllers.UploadFile)
		g.GET("/", controllers.AcquireFileName)
		g.GET("/:id", controllers.GetFile)
		g.GET("/:id/meta", controllers.GetFileMeta)
		g.DELETE("/:id", controllers.DeleteFile)
	}

	router.Run(addr)

	return nil
}
