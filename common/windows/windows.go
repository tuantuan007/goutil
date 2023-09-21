//go:build windows

package windows

import (
	"github.com/yusufpapurcu/wmi"
	"log"
)

func Query[T any]() ([]T, error) {
	var dst []T
	query := wmi.CreateQuery(&dst, "")
	if err := wmi.Query(query, &dst); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return dst, nil
}
