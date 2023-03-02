package restaurantlikemodel

// To list which Users has liked RestaurantId
type Filter struct {
	RestaurantId int `json:"-" form:"restaurant_id"`
	UserId       int `json:"-" form:"user_id"`
}
