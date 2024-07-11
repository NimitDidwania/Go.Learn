package models

type GetAllItemsRequest struct {
	PageNo       int    `json:"pageNo"`
	PageSize     int    `json:"pageSize"`
	SearchString string `json:"searchString"`
}
