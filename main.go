package main

import (
    "encoding/json"
    "github.com/gin-gonic/gin"
    "html/template"
    "net/http"
    "os"
    
)

type LogEntry struct {
    LogDate  string `json:"log_date"`
    Thoughts string `json:"thoughts"`
    Category string `json:"category"`
    Tools    string `json:"tools"`
    Day      string `json:"day"`
}

func loadLogs() ([]LogEntry, error) {
    file, err := os.Open("logs/logs.json")
    if err != nil {
        return []LogEntry{}, nil // return empty if file doesn't exist
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
        "templates/layout.html",
        "templates/index.html",
        "templates/about.html",
        "templates/log.html",
        "templates/audit_logs.html",
        "templates/bug_logs.html",
        
    ))
    router.SetHTMLTemplate(tmpl)

    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", nil)
    })

    router.GET("/about", func(c *gin.Context) {
        c.HTML(http.StatusOK, "about.html", nil)
    })

    // router.GET("/about", func(c *gin.Context) {
    //     tmpl := template.Must(template.ParseFiles(
    //         "templates/layout.html",
    //         "templates/about.html",
    //         ))
    //         tmpl.ExecuteTemplate(c.Writer , "about",nil)
    // })  

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

    router.Run(":8080")
}
