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
	Logo            *common.Image  `json:"logo" gorm:"logo;"`
	Cover           *common.Images `json:"cover" gorm:"cover;"` // Có thể là 1 dạng chạy slide các ảnh...
	// Cover           []common.Images `json:"cover" gorm:"cover;"` || ko dc sd như này
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
	Name            string         `json:"name" gorm:"column:name;"`
	Addr            string         `json:"addr" gorm:"column:addr;"`
	Logo            *common.Image  `json:"logo" gorm:"logo;"`
	Cover           *common.Images `json:"cover" gorm:"cover;"`
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
	Name  *string        `json:"name" gorm:"column:name;"`
	Addr  *string        `json:"addr" gorm:"column:addr;"`
	Logo  *common.Image  `json:"logo" gorm:"logo;"` // Phải là pointer (vì nếu ko có thì nên về nil, đừng về struct rỗng)
	Cover *common.Images `json:"cover" gorm:"cover;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

var (
	ErrNameIsEmpty = errors.New("name cannot be empty") // Ko dc ghép phía trên cho Unit test => vì luôn về new pointer (ko thể so sánh dc)
)
