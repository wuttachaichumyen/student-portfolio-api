package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wuttachaichumyen/student-portfolio-api/handlers" // แก้ YOUR_USERNAME
)

func main() {
	r := gin.Default()

	// ตั้งค่า CORS (สำคัญมาก เพื่อให้ Frontend React เรียกใช้งานได้)
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Group Route เพื่อความสวยงาม เช่น /api/v1/students
	api := r.Group("/api/v1")
	{
		api.GET("/students", handlers.GetStudents)
		api.GET("/students/:id", handlers.GetStudentByID)
		api.POST("/students", handlers.CreateStudent)
	}

	// รัน Server ที่ Port 8080
	r.Run(":8080")
}