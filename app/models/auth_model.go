package model

import (
	"gin-web-admin/utils"
	"github.com/jinzhu/gorm"
)

type Auth struct {
	BaseModel
	RoleId     int      `gorm:"Size:4;DEFAULT:0;NOT NULL;" json:"role_id"`
	Status     int      `gorm:"type:int(1);DEFAULT:0;NOT NULL;" json:"status"`
	LoggedInAt JSONTime `json:"logged_in_at"`
	Username   string   `gorm:"Size:20;UNIQUE_INDEX;NOT NULL;" json:"user_name"`
	Password   string   `gorm:"Size:50;NOT NULL;" json:"-"`
	RoleName   string   `gorm:"-" json:"role_name"`
	Role       Role     `gorm:"foreignkey:RoleId;association_foreignkey:RoleId;" json:"-"`
}

func (Auth) TableName() string {
	return TablePrefix + "auth"
}

func CheckAuth(username string, password string) (bool, uint, string, bool) {
	var auth Auth
	db.Select([]string{"id", "role_id"}).Where(Auth{
		Username: username,
		Password: utils.EncodeUserPassword(password),
	}).Preload("Role").First(&auth)

	if auth.ID > 0 {
		return true, auth.ID, auth.Role.RoleKey, auth.Role.IsAdmin
	}

	return false, 0, "", false
}

func CreatUser(auth Auth) error {
	db.NewRecord(auth)
	res := db.Create(&auth)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

// GetTestUsers gets a list of users based on paging constraints
func GetUsers(pageNum int, pageSize int, maps interface{}) ([]*Auth, error) {
	var user []*Auth
	err := db.Select("*").Where(maps).Offset(pageNum).Limit(pageSize).Preload(
		"Role", func(db *gorm.DB) *gorm.DB {
			return db.Select("role_id,role_name")
		}).Find(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return user, nil
}
