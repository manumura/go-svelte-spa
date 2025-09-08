package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"spa/common"
	"spa/ui"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	// start gin
	r := gin.Default()
	r.SetTrustedProxies(nil)

	// setup cors middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	// health route
	r.GET("/health", handleHealth)

	// API routes
	api := r.Group("/api")
	api.GET("/users", handleUsers)
	api.GET("/users/:id", handleUserByID)

	isProd := common.IsProduction()

	if isProd {
		fs, err := static.EmbedFolder(ui.DistDir, "dist")
		if err != nil {
			log.Fatal(err)
		}
		r.Use(static.Serve("/", fs))

		r.NoRoute(func(c *gin.Context) {
			fmt.Printf("%s doesn't exists, serving index.html\n", c.Request.URL.Path)
			c.File(ui.IndexFilePath)
		})
	} else {
		if !common.IsViteServerRunning() {
			log.Fatalln("Vite dev server is not running. Please start it with 'pnpm dev' or 'npm run dev'")
		}
		// Static assets and frontend routes
		r.Any("/", staticAssetHandler)
	}

	fmt.Println("📡 API endpoints: /api/*")
	fmt.Println("💚 Health check: /health")
	if isProd {
		fmt.Println("🌐 Serving static files from ui/dist directory")
	} else {
		fmt.Println("🔄 Proxying to Vite dev server at http://localhost:5173")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Starting server on :" + port)
	log.Fatal(r.Run(":" + port))
}

func handleHealth(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"status":  "healthy",
		"service": "go-backend",
	})
}

func handleUsers(c *gin.Context) {
	c.JSON(http.StatusOK, common.TestUsers)
}

func handleUserByID(c *gin.Context) {
	id := c.Param("id")
	for _, user := range common.TestUsers {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
}

func staticAssetHandler(c *gin.Context) {
	viteProxy := common.GetViteProxy()
	// Proxy to Vite for frontend assets and routes
	viteProxy.ServeHTTP(c.Writer, c.Request)
}
