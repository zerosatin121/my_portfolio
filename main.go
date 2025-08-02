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
        filepath.Join("templates", "audit.html"),
        filepath.Join("templates", "bugs.html"),
        filepath.Join("templates", "log.html"),
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
// Route: Audit challenge page
    router.GET("/audit-challenge", func(c *gin.Context) {
    c.HTML(http.StatusOK, "audit.html", nil)
})

//route to bug hunting page
router.GET("/bug-hunting", func(c *gin.Context) {
    c.HTML(http.StatusOK, "bugs.html", nil)
})

// Route: Daily log page
router.GET("/daily-log", func(c *gin.Context) {
    c.HTML(http.StatusOK, "log.html", nil)
})


    // Start server
    router.Run(":8080") // Gin logs each request automatically
}
