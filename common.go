package gripp

type WithSearchName[T any] struct {
	ID         T      `json:"id"`
	SearchName string `json:"searchname"`
	Discr      string `json:"discr,omitempty"`
}
