package main

import (
    "encoding/json"
    "html/template"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
)

type LogEntry struct {
    LogDate  string `json:"log_date"`
    Thoughts string `json:"thoughts"`
    Category string `json:"category"`
    Tools    string `json:"tools"`
    Day      string `json:"day"`
}

type BugMetrics struct {
    BugSeverity map[string]int `json:"bug_severity"`
    BugStatus   map[string]int `json:"bug_status"`
}

func loadMetrics() (BugMetrics, error) {
    var metrics BugMetrics
    file, err := os.Open("logs/metrics.json")
    if err != nil {
        // Return empty metrics if file doesn't exist
        return BugMetrics{
            BugSeverity: map[string]int{},
            BugStatus:   map[string]int{},
        }, nil
    }
    defer file.Close()
    err = json.NewDecoder(file).Decode(&metrics)
    return metrics, err
}

func saveMetrics(metrics BugMetrics) error {
    file, err := os.Create("logs/metrics.json")
    if err != nil {
        return err
    }
    defer file.Close()
    return json.NewEncoder(file).Encode(metrics)
}

func loadLogs() ([]LogEntry, error) {
    file, err := os.Open("logs/logs.json")
    if err != nil {
        return []LogEntry{}, nil
    }
    defer file.Close()

    var logs []LogEntry
    err = json.NewDecoder(file).Decode(&logs)
    return logs, err
}

func saveLog(entry LogEntry) error {
    logs, _ := loadLogs()
    logs = append(logs, entry)

    file, err := os.Create("logs/logs.json")
    if err != nil {
        return err
    }
    defer file.Close()

    return json.NewEncoder(file).Encode(logs)
}

func main() {
    router := gin.Default()
    router.Static("/static", "./static")

    tmpl := template.Must(template.ParseFiles(
        "templates/index.html",
        "templates/about.html",
        "templates/log.html",
        "templates/audit_logs.html",
        "templates/bug_logs.html",
        "templates/update.html", // Include update page
    ))
    router.SetHTMLTemplate(tmpl)

    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", nil)
    })

    router.GET("/about", func(c *gin.Context) {
        c.HTML(http.StatusOK, "about.html", nil)
    })

    router.GET("/daily-log", func(c *gin.Context) {
        c.HTML(http.StatusOK, "log.html", nil)
    })

    router.POST("/submit-log", func(c *gin.Context) {
        entry := LogEntry{
            LogDate:  c.PostForm("log_date"),
            Thoughts: c.PostForm("thoughts"),
            Category: c.PostForm("category"),
            Tools:    c.PostForm("tools"),
            Day:      c.PostForm("day"),
        }
        if err := saveLog(entry); err != nil {
            c.String(http.StatusInternalServerError, "Failed to save log")
            return
        }
        c.Redirect(http.StatusSeeOther, "/daily-log")
    })

    router.GET("/audit-logs", func(c *gin.Context) {
        logs, _ := loadLogs()
        var auditLogs []LogEntry
        for _, log := range logs {
            if log.Category == "audit" {
                auditLogs = append(auditLogs, log)
            }
        }
        c.HTML(http.StatusOK, "audit_logs.html", gin.H{"Logs": auditLogs})
    })

    router.GET("/bug-logs", func(c *gin.Context) {
        logs, _ := loadLogs()
        var bugLogs []LogEntry
        for _, log := range logs {
            if log.Category == "bug" {
                bugLogs = append(bugLogs, log)
            }
        }
        c.HTML(http.StatusOK, "bug_logs.html", gin.H{"Logs": bugLogs})
    })

    router.GET("/calendar", func(c *gin.Context) {
        tmpl := template.Must(template.ParseFiles("templates/calendar.html"))
        tmpl.ExecuteTemplate(c.Writer, "calendar.html", nil)
    })

    router.StaticFile("/download-logs", "./logs/logs.json")

    // 📊 Metrics API
    router.GET("/api/metrics", func(c *gin.Context) {
        metrics, err := loadMetrics()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load metrics"})
            return
        }
        c.JSON(http.StatusOK, metrics)
    })

    router.POST("/api/update-metrics", func(c *gin.Context) {
        var newData BugMetrics
        if err := c.ShouldBindJSON(&newData); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
            return
        }

        // Overwrite existing metrics instead of accumulating
        updated := BugMetrics{
            BugSeverity: newData.BugSeverity,
            BugStatus:   newData.BugStatus,
        }

        if err := saveMetrics(updated); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save metrics"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "Metrics updated"})
    })

    // 📝 Update Page
    router.GET("/update", func(c *gin.Context) {
        c.HTML(http.StatusOK, "update.html", nil)
    })

    router.Run(":8080")
}
