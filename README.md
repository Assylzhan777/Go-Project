Көлік дөңгелектері REST API
Бұл құжат Көлік дөңгелектері жобасына арналған REST API түсіндірмесін береді. Жоба Go тілінде жазылған, Gin фреймворкі және GORM кітапханасы арқылы жүзеге асырылған. Дерекқор ретінде PostgreSQL пайдаланылады.

1. Gin фреймворкі
   Gin — Go тілінде жазылған жеңіл әрі жылдам HTTP сервер фреймворкі. Ол API жасау үшін өте ыңғайлы.

Артықшылықтары:
Жылдам жұмыс істейді (net/http негізінде)

JSON-пен жұмыс істеу оңай

Middleware қосуға болады

Орнату:
bash
Копировать
Редактировать
go get -u github.com/gin-gonic/gin
Қарапайым сервер мысалы:
go
Копировать
Редактировать
package main

import "github.com/gin-gonic/gin"

func main() {
r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "API жұмыс істеп тұр"})
    })

    r.Run(":8080")
}
2. GORM — Go үшін ORM
   GORM — Go тілінде дерекқорлармен жұмыс істеуге арналған ORM кітапханасы. Бұл кітапхана дерекқорды басқаруды жеңілдетеді.

Орнату:
bash
Копировать
Редактировать
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
Дерекқорға қосылу:
go
Копировать
Редактировать
import (
"gorm.io/driver/postgres"
"gorm.io/gorm"
"log"
)

func main() {
dsn := "host=localhost user=postgres password=7982 dbname=postgres port=5432 sslmode=disable"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
if err != nil {
log.Fatal("Қосылу қатесі:", err)
}
}
3. Өнім моделі мен кесте құру
   Модель:

go
Копировать
Редактировать
type Tire struct {
ID          uint   `gorm:"primaryKey"`
Brand       string `json:"brand"`
Model       string `json:"model"`
Diameter    float64 `json:"diameter"`
Width       int    `json:"width"`
Price       uint   `json:"price"`
}
Миграция:

go
Копировать
Редактировать
db.AutoMigrate(&Tire{})
4. Негізгі CRUD маршруттар
   Барлық дөңгелектерді алу
   go
   Копировать
   Редактировать
   r.GET("/tires", func(c *gin.Context) {
   var tires []Tire
   db.Find(&tires)
   c.JSON(200, tires)
   })
   Бір дөңгелекті көру
   go
   Копировать
   Редактировать
   r.GET("/tires/:id", func(c *gin.Context) {
   var tire Tire
   if err := db.First(&tire, c.Param("id")).Error; err != nil {
   c.JSON(404, gin.H{"error": "Дөңгелек табылмады"})
   return
   }
   c.JSON(200, tire)
   })
   Жаңа дөңгелек қосу
   go
   Копировать
   Редактировать
   r.POST("/tires", func(c *gin.Context) {
   var tire Tire
   if err := c.ShouldBindJSON(&tire); err != nil {
   c.JSON(400, gin.H{"error": err.Error()})
   return
   }
   db.Create(&tire)
   c.JSON(201, tire)
   })
   Дөңгелекті жаңарту
   go
   Копировать
   Редактировать
   r.PUT("/tires/:id", func(c *gin.Context) {
   var tire Tire
   if err := db.First(&tire, c.Param("id")).Error; err != nil {
   c.JSON(404, gin.H{"error": "Дөңгелек табылмады"})
   return
   }
   var input Tire
   if err := c.ShouldBindJSON(&input); err != nil {
   c.JSON(400, gin.H{"error": err.Error()})
   return
   }
   db.Model(&tire).Updates(input)
   c.JSON(200, tire)
   })
   Дөңгелекті жою
   go
   Копировать
   Редактировать
   r.DELETE("/tires/:id", func(c *gin.Context) {
   var tire Tire
   if err := db.First(&tire, c.Param("id")).Error; err != nil {
   c.JSON(404, gin.H{"error": "Дөңгелек табылмады"})
   return
   }
   db.Delete(&tire)
   c.JSON(200, gin.H{"message": "Дөңгелек жойылды"})
   })
5. Тексеру мысалы (Postman арқылы)
   GET http://localhost:8080/tires

POST http://localhost:8080/tires

json
Копировать
Редактировать
{
"brand": "Michelin",
"model": "Pilot Sport 4",
"diameter": 18,
"width": 225,
"price": 85000
}
6. Қорытынды
   Бұл жоба Көлік дөңгелектері интернет дүкеніне арналған REST API құрды. Gin және GORM кітапханаларының көмегімен дөңгелектерді қосу, көру, жаңарту және жою сияқты CRUD операцияларын жасау мүмкіндігі бар.