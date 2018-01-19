package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gildasch/upspin-music/album"
	"github.com/gildasch/upspin-music/upspin"
	"github.com/gin-gonic/gin"
	"upspin.io/client"
	"upspin.io/config"
	_ "upspin.io/transports"
)

type Accesser interface {
	List(path string) ([]*album.Album, error)
	Get(path string) (io.Reader, error)
}

func main() {
	cfg, err := config.FromFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	client := client.New(cfg)
	if client == nil {
		fmt.Println("client could be initialized")
	}

	accesser := upspin.Accesser{client}

	router := gin.Default()

	router.Static("/static", "./static")
	router.LoadHTMLFiles("templates/index.html")

	router.GET("/listen/*path", func(c *gin.Context) {
		albums, err := accesser.List(c.Param("path"))
		if err != nil {
			fmt.Println("accesser.List:", err)
			c.Status(http.StatusNotFound)
			return
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"albums": albums,
		})
	})

	router.GET("/get/*path", func(c *gin.Context) {
		reader, err := accesser.Get(c.Param("path"))
		if err != nil {
			fmt.Println("accesser.Get:", err)
			c.Status(http.StatusNotFound)
			return
		}

		c.Stream(func(w io.Writer) bool {
			_, err := io.CopyN(w, reader, 1024*1024)
			return err == nil
		})
	})

	router.Run()
}