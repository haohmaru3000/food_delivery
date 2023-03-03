package uploadprovider

import (
	"context"

	"github.com/0xThomas3000/food_delivery/common"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
