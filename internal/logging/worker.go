package logging

import (
	"gorm.io/gorm"
)

// var LogQueue = make(chan models.RequestLog, 500)

func StartLogWorker(db *gorm.DB) {
	go func() {
		// for x := range LogQueue {
		// data := models.RequestLog{
		// 	Path:      x.Path,
		// 	Method:    x.Method,
		// 	Username:  x.Username,
		// 	Request:   x.Request,
		// 	IP:        x.IP,
		// 	CreatedAt: x.CreatedAt,
		// }

		// if err := db.Create(&data).Error; err != nil {
		// 	fmt.Println("error insert log:", err)
		// }
		// }
	}()
}
