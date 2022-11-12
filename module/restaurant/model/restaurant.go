package restaurantmodel

import (
	"github.com/0xThomas3000/food_delivery/common"
	"strings"
)

const EntityName = "Restaurant" // Tạo EntityName vì lỗi này dc dùng đi dùng lại

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	Addr            string `json:"addr" gorm:"column:addr;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	// Nếu là Admin: ko cần che giấu info nhiều | là user: cần che vài info
	// VD: CMS lấy API thì đầy đủ, người b thường lấy API thì bị hạn chế
	r.GenUID(common.DbTypeRestaurant)
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	Addr            string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (data *RestaurantCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
}

func (res *RestaurantCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return ErrNameCannotBeEmpty
	}

	return nil
}

var (
	ErrNameCannotBeEmpty = common.NewCustomError(nil, "restaurant name can't be blank", "ErrNameCannotBeEmpty")
)

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}
