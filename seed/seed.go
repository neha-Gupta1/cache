package seed

import (
	"log"

	"github.com/cache/model"
	"github.com/jinzhu/gorm"
)

// var values = []model.Data{
// 	model.Data{
// 		ID:    1,
// 		Value: "first cache value",
// 	},
// 	model.Data{
// 		ID:    10,
// 		Value: "value at 10",
// 	},
// 	model.Data{
// 		ID:    15,
// 		Value: "value at 15",
// 	},
// 	model.Data{
// 		ID:    16,
// 		Value: "value at 16",
// 	},
// 	model.Data{
// 		ID:    26,
// 		Value: "value at 26 goes in same row as 16",
// 	},
// 	model.Data{
// 		ID:    11,
// 		Value: "value at 11",
// 	},
// }

// Load is the inital seeder
func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&model.Data{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&model.Data{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	// for i := range values {
	// 	err = db.Debug().Model(&models.Data{}).Create(&values[i]).Error
	// 	if err != nil {
	// 		log.Fatalf("cannot seed users table: %v", err)
	// 	}

	// }
}
