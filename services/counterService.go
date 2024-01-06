package services

import (
	"github.com/sCuz12/ancient-greek-quote-api/models"
	"gorm.io/gorm"
)

type CounterService struct {
	DB *gorm.DB
}


func ConstructCounterService(db *gorm.DB) *CounterService {
	return &CounterService{DB:db}
}

func (cs *CounterService) IncrementCounter() error  {
	
	var counter models.Counter
    if err := cs.DB.FirstOrCreate(&counter, models.Counter{Id: 1}).Error; err != nil {
        return err
    }

    counter.Count++
    return cs.DB.Save(&counter).Error
	
}