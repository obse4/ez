package request

type FindPagingCommonRequest struct {
	// 分页大小
	PageSize int `json:"page_size"`
	// 分页索引
	PageIndex int `json:"page_index"`
}
