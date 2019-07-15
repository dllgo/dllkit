package gins

// 分页请求数据
type Paging struct {
	PageNo     int `json:"pageNo"`     // 页码
	PageSize   int `json:"pageSize"`   // 每页条数
	TotalPage  int `json:"totalPage"`  // 总数据条数
	TotalCount int `json:"totalCount"` // 总数据条数
}

func NewPaging() *Paging {
	return &Paging{}
}
func (p *Paging) Offset() int {
	offset := 0
	if p.PageNo > 0 {
		offset = (p.PageNo - 1) * p.PageSize
	}
	return offset
}

func (p *Paging) TotalPages() int {
	if p.TotalCount == 0 || p.PageSize == 0 {
		return 0
	}
	totalPage := p.TotalCount / p.PageSize
	if p.TotalCount%p.PageSize > 0 {
		totalPage = totalPage + 1
	}
	return totalPage
}
