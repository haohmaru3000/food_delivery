package restaurantmodel

import (
	"errors"
	"strings"

	"github.com/0xThomas3000/food_delivery/common"
)

type RestaurantType string

const TypeNormal RestaurantType = "normal"
const TypePremium RestaurantType = "premium"
const EntityName = "Restaurant" // Tạo EntityName vì lỗi này dc dùng đi dùng lại

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	Addr            string         `json:"addr" gorm:"column:addr;"`
	Type            RestaurantType `json:"type" gorm:"column:type;"` // kiểu enum (như options cho các quyền...)
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

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if data.Name == "" {
		return ErrNameIsEmpty // Chỉ cần unit test Validate() về chính xác giá trị này
	}

	return nil
}

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

var (
	ErrNameIsEmpty = errors.New("name cannot be empty") // Ko dc ghép phía trên cho Unit test => vì luôn về new pointer (ko thể so sánh dc)
)
