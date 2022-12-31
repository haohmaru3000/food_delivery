package userstore

import (
	"context"
	"github.com/0xThomas3000/food_delivery/common"
	"github.com/0xThomas3000/food_delivery/modules/user/model"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
	db := s.db.Begin() // Mở 1 Transaction ra (nếu thực hiện nhiều lệnh bên dưới), rồi Commit nó.

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		// Có lỗi mà ko Rollback, conn ko nhả. Quá nhìu TX mở ra mà ko rollback/commit -> too many conn to DB -> crash DB
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
