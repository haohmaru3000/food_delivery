package uploadbusiness

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"io"
	"log"
	"path/filepath"
	"time"

	"github.com/0xThomas3000/food_delivery/common"
	"github.com/0xThomas3000/food_delivery/components/uploadprovider"
	"github.com/0xThomas3000/food_delivery/module/upload/uploadmodel"
)

type CreateImageStorage interface {
	CreateImage(context context.Context, data *common.Image) error
}

type uploadBiz struct {
	// chỉ là Interface cung cấp hàm SaveFileUpload phía dưới
	provider uploadprovider.UploadProvider // -> che dấu cách ta lưu trữ hình ở đâu (cloud, write to our own service = cách implement this interface)
	imgStore CreateImageStorage
}

func NewUploadBiz(provider uploadprovider.UploadProvider, imgStore CreateImageStorage) *uploadBiz {
	return &uploadBiz{provider: provider, imgStore: imgStore}
}

func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	// io reader có nhiều loại (bytes.NewBuffer là buffer thuần, means đọc buffer ra mà ko cần mở đi đâu hết)
	fileBytes := bytes.NewBuffer(data) // Mở NewBuffer file ra vì có hàm getImageDimension

	w, h, err := getImageDimension(fileBytes)

	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage(err)
	}

	fileExt := filepath.Ext(fileName)                                // "img.jpg" => ".jpg"
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt) // 9129324893248.jpg

	img, err := biz.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))

	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	img.Width = w
	img.Height = h
	//img.CloudName = "s3" // should be set in provider
	img.Extension = fileExt

	// Các hàm sau là lưu xuống DB (nhưng ko cần => comment out)
	//if err := biz.imgStore.CreateImage(ctx, img); err != nil {
	//	// delete img on S3
	//	return nil, uploadmodel.ErrCannotSaveFile(err)
	//}

	return img, nil
}

// Lấy thông tin Width, Height và gọi thông qua hàm DecodeConfig
func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("err: ", err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
