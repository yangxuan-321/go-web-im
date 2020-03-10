package util

import (
	"github.com/satori/go.uuid"
)

func GetUUID()(uuid_ string) {
	// 创建 UUID v4
	u1 := uuid.Must(uuid.NewV4())
	return u1.String()
}
