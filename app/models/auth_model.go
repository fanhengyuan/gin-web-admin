package model

import (
    "gin-test/utils"
    "github.com/jinzhu/gorm"
    "time"
)

type Auth struct {
    BaseModel
    RoleType int `gorm:"Size:4" json:"role_type"`
    LoggedInAt JSONTime `json:"logged_in_at"`
    Username string `gorm:"Size:20" json:"user_name"`
    Password string `gorm:"Size:50" json:"-"`
}

func CheckAuth(username, password string) (bool, uint) {
    var auth Auth
    db.Select("id").Where(Auth{
        Username : username,
        Password : utils.EncodeMD5(password),
    }).First(&auth)

    if auth.ID > 0 {
        // 记录登录时间
        db.Model(&auth).Update("logged_in_at", time.Now())
        return true, auth.ID
    }

    return false, 0
}

func GetUserTotal(maps interface{}) (int, error) {
    var count int
    if err := db.Model(&Auth{}).Where(maps).Count(&count).Error; err != nil {
        return 0, err
    }

    return count, nil
}

// GetTestUsers gets a list of users based on paging constraints
func GetUsers(pageNum int, pageSize int, maps interface{}) ([]*Auth, error) {
    var user [] *Auth
    err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&user).Error
    if err != nil && err != gorm.ErrRecordNotFound {
        return nil, err
    }

    return user, nil
}
