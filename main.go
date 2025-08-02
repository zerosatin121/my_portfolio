package main

import (
    "github.com/gin-gonic/gin"
    "html/template"
    "net/http"
    "path/filepath"
)

func main() {
    // Create Gin router with default middleware (Logger + Recovery)
    router := gin.Default()

    // Serve static files from ./static at /static
    router.Static("/static", "./static")

    // Load templates
    tmpl := template.Must(template.ParseFiles(
        filepath.Join("templates", "index.html"),
        filepath.Join("templates", "about.html"),
    ))
    router.SetHTMLTemplate(tmpl)

    // Route: Landing page
    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", nil)
    })

    // Route: About page
    router.GET("/about", func(c *gin.Context) {
        c.HTML(http.StatusOK, "about.html", nil)
    })

    // Start server
    router.Run(":8080") // Gin logs each request automatically
}
