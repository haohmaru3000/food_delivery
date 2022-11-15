package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Image struct {
	Id        int    `json:"id" gorm:"column:id;"`
	Url       string `json:"url" gorm:"column:url;"`
	Width     int    `json:"width" gorm:"column:width;"`
	Height    int    `json:"height" gorm:"column:height;"`
	CloudName string `json:"cloud_name,omitempty" gorm:"-"`
	Extension string `json:"extension,omitempty" gorm:"-"` // Phần đuôi cho client hiển thị icon
}

func (Image) TableName() string { return "images" }

// Đi ngược từ DB đi ra
func (j *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte) // Đổi value thành mảng byte
	if !ok {                    // check coi có mảng byte ko, nếu ko thì báo lỗi
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var img Image // Tạo ra 1 structure mới image

	if err := json.Unmarshal(bytes, &img); err != nil { // Chuyển từ mảng bytes sang structure image
		return err
	}

	*j = img // Thay đổi value của con trỏ về image
	return nil
}

// Value return json value, implement driver.Valuer interface (từ structure đi vô DB)
func (j *Image) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

type Images []Image

func (j *Images) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var img []Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*j = img
	return nil
}

// Value return json value, implement driver.Valuer interface
func (j *Images) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}
