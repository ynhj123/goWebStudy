package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	gogpt "github.com/sashabaranov/go-gpt3"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type SteamRequest struct {
	Msg   string `form:"msg" binding:"required"`
	Chats string `form:"chats" binding:"required"`
}
type Chat struct {
	A string `form:"a"`
	Q string `form:"q"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.GET("/stream", func(c *gin.Context) {
		var request SteamRequest
		err := c.ShouldBindQuery(&request)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Println("request" + request.Msg)
		var chats Chat
		errs := json.Unmarshal([]byte(request.Chats), &chats)
		if errs != nil {
			fmt.Println("json unmarshal error:", errs)
		}
		//chanStream := make(chan int, 10)
		//go func() {
		//	defer close(chanStream)
		//	for i := 0; i < 5; i++ {
		//		chanStream <- i
		//		time.Sleep(time.Second * 1)
		//	}
		//}()
		//c.Stream(func(w io.Writer) bool {
		//	if msg, ok := <-chanStream; ok {
		//		c.SSEvent("message", msg)
		//		if msg == 4 {
		//			c.SSEvent("stop", msg)
		//		}
		//		return true
		//	}
		//	return false
		//})
		client := gogpt.NewClient("you token")
		ctx := context.Background()

		req := gogpt.CompletionRequest{
			Model:     gogpt.GPT3TextDavinci003,
			MaxTokens: 5,
			Prompt:    request.Msg,
			Stream:    true,
		}
		stream, err := client.CreateCompletionStream(ctx, req)
		if err != nil {
			return
		}

		chanStream := make(chan string, 100)
		go func() {
			defer stream.Close()
			defer close(chanStream)
			for {
				response, err := stream.Recv()
				if errors.Is(err, io.EOF) {
					fmt.Println("Stream finished")
					chanStream <- "<!finish>"
					return
				}

				if err != nil {
					fmt.Printf("Stream error: %v\n", err)
					chanStream <- "<!error>"
					return
				}
				if len(response.Choices) == 0 {
					fmt.Println("Stream finished")
					chanStream <- "<!finish>"
					return
				}
				data, err := json.Marshal(response.Choices[0])
				chanStream <- string(data)
				fmt.Printf("Stream response: %v\n", response.Choices[0])
			}
		}()

		c.Stream(func(w io.Writer) bool {
			if msg, ok := <-chanStream; ok {
				if msg == "<!finish>" {
					c.SSEvent("stop", "finish")
				}
				if msg == "<!error>" {
					c.SSEvent("stop", "error")
				}
				c.SSEvent("message", msg)
				fmt.Printf("message: %v\n", msg)

				return true
			}
			return false
		})

	})
	return r
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := setupRouter()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
