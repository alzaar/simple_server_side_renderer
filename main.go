package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TaskForm struct {
	Title     string `form:"title" binding:"required"`
	Submitted bool
	Tasks     []Task
}

type Task struct {
	ID        int
	Title     string
	CreatedAt time.Time
}

var tasks []Task

func main() {
	router := gin.Default()

	router.Static("/assets", "./assets")
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/submit", func(c *gin.Context) {
		var form TaskForm
		// Bind form data from the request
		if err := c.ShouldBind(&form); err != nil {
			c.String(http.StatusBadRequest, "Bad Request")
			return
		}

		// Set the Submitted field to true
		form.Submitted = true

		// Render the form with the submitted data
		renderForm(c, form)
	})

	router.Run(":8080")
}

func renderForm(c *gin.Context, data TaskForm) {
	template, err := template.ParseFiles("templates/index.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	addTask(data.Title)
	data.Tasks = tasks
	err = template.Execute(c.Writer, data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
}

func addTask(title string) {
	tasks = append(tasks, Task{len(tasks) + 1, title, time.Now()})
}
