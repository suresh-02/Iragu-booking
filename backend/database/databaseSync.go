package database

import "github.com/suresh-02/Iragu-booking/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.UserCreds{})
}
