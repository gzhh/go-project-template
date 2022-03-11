package webapp

import (
	"context"
	"demo/cmd/webapp/handler"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"net/http"
	"time"
)

func init() {
	rootCmd.AddCommand(webCmd)
}

var webCmd = &cobra.Command{
	Use:   "webapp",
	Short: "web server",
	Long:  `web server, handler client request`,
	Run:   web,
}

func timeoutMiddleware(timeout time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		// wrap the request context with a timeout
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)

		defer func() {
			// check if context timeout was reached
			if ctx.Err() == context.DeadlineExceeded {
				c.Abort()
			}
			//cancel to clear resources after finished
			cancel()
		}()

		// replace request with context wrapped request
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func web(cmd *cobra.Command, args []string) {
	router := gin.New()
	router.Use(timeoutMiddleware(time.Second*60)).GET("/api/test", handler.Test)

	srv := &http.Server{
		Addr:    ":8890",
		Handler: router,
	}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
