package ginupload

import (
	"github.com/gin-gonic/gin"
	_ "image/jpeg"
	_ "image/png"

	"github.com/0xThomas3000/food_delivery/common"
	"github.com/0xThomas3000/food_delivery/component/appctx"
	"github.com/0xThomas3000/food_delivery/module/upload/uploadbusiness"
)

func Upload(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		//db := appCtx.GetMainDBConnection()

		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// dùng key lấy là 'folder', value nếu ko có thì lấy 'img' | nếu có thì lấy đúng cái dc truyền vào
		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close() // đây là 1 stream file => nên close nó nếu có bất cứ vấn đề gì

		dataBytes := make([]byte, fileHeader.Size)      // data bytes dc tạo có kích thước = fileHeader.Size
		if _, err := file.Read(dataBytes); err != nil { // đổ hết nội dung file vào mảng data bytes phía trên
			panic(common.ErrInvalidRequest(err)) // khi gặp panic nó sẽ chạy defer phía trên để close
		}

		//imgStore := uploadstorage.NewSQLStore(db)
		// làm cách nào vận chuyển UploadProvider vào tới Transport này? => dùng "appCtx"
		biz := uploadbusiness.NewUploadBiz(appCtx.UploadProvider(), nil)                    // business lên
		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename) // lấy mảng bytes(toàn bộ dữ liệu của file) gọi qua business

		if err != nil {
			panic(err)
		}
		c.JSON(200, common.SimpleSuccessResponse(img))
	}
}
