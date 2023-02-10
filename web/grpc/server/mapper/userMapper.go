package mapper

import (
	"fmt"
	"server/mapper/entity"
)

type userMapper struct {
}

func (userMapper) getByUserName(username string) entity.User {
	var user entity.User
	result := db.First(&user, "username = ?", username)
	fmt.Println(result.RowsAffected) // 返回插入记录的条数
	if result.Error == nil {         // 检测error
		fmt.Println("no errors!!!")
	}
	return user
}
