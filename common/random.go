package common

import (
	"math/rand"
	"time"
)

// Khai báo mảng kí tự
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSequence(n int) string {
	b := make([]rune, n)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Cần phải chạy thêm rand.Seed(), nếu ko sẽ luôn trả về giá trị giống nhau cho các lần chạy kế tiếp.
	// rand.Intn(10000)

	for i := range b { // Khi đã có gói Random với Source mới, loop để lấy ra random kí tự trong []letters.
		index := r1.Intn(99999) % len(letters) // Chỉ số ngẫu nhiên trong phạm vi giới hạn mảng letters.
		b[i] = letters[index]
	}
	return string(b)
}

func GenSalt(length int) string {
	if length < 0 {
		length = 50
	}
	return randSequence(length)
}
