package utils

import (
	"log"
	service "staycation/internal/services"

	"github.com/robfig/cron/v3"
)

type CronJobService struct {
	roomService service.RoomService
}

func NewCronJobService(roomService service.RoomService) *CronJobService {
	return &CronJobService{roomService}
}

func (c *CronJobService) Start() {
	cronJob := cron.New()

	cronJob.AddFunc("@every 1h", func() {
		err := c.roomService.UpdateRoomStatus()
		if err != nil {
			log.Println("Error updating room status:", err)
		} else {
			log.Println("Room statuses updated successfully.")
		}
	})

	cronJob.Start()
}
