package restaurantrepo

import (
	"context"

	"github.com/0xThomas3000/food_delivery/common"
	"github.com/0xThomas3000/food_delivery/modules/restaurant/model"
)

type ListRestaurantStore interface {
	ListDataWithCondition(ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

// type LikeRestaurantStore interface {
// 	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
// }

type listRestaurantRepo struct {
	store ListRestaurantStore
	// likeStore LikeRestaurantStore
}

func NewListRestaurantRepo(store ListRestaurantStore) *listRestaurantRepo {
	return &listRestaurantRepo{
		store: store,
	}
}

func (repo *listRestaurantRepo) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := repo.store.ListDataWithCondition(ctx, filter, paging, "User")

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	// ids := make([]int, len(result))

	// for i := range result {
	// 	ids[i] = result[i].Id
	// }

	// likeMap, err := repo.likeStore.GetRestaurantLikes(ctx, ids)

	// if err != nil {
	// 	log.Println("Cannot get restaurant likes:", err)
	// 	return result, nil
	// }

	// if v := likeMap; v != nil {
	// 	for i, item := range result {
	// 		result[i].LikedCount = likeMap[item.Id]
	// 	}
	// }

	// List restaurant only has liked_count > 10

	return result, nil
}
