package dto

// 分页默认值与全局硬上限
const (
	DefaultPageSize  = 10
	MaxPageSizeLimit = 10000 //全局分页硬上限（防御性，正常路径由 GetPageSize 限制）
)

type Pagination struct {
	PageIndex int `form:"pageIndex"`
	PageSize  int `form:"pageSize"`
	// PageSizeLimit 单页条数上限（默认100；导出等场景由服务端放宽，客户端不可传）
	PageSizeLimit int `form:"-"`
}

func (m *Pagination) GetPageIndex() int {
	if m.PageIndex <= 0 {
		m.PageIndex = 1
	}
	return m.PageIndex
}

// GetPageSize 单页条数：默认10，上限默认100（导出等场景可服务端设置 PageSizeLimit 放宽）
func (m *Pagination) GetPageSize() int {
	limit := m.PageSizeLimit
	if limit <= 0 {
		limit = 100
	}
	if m.PageSize <= 0 {
		if DefaultPageSize > limit {
			return limit
		}
		return DefaultPageSize
	}
	if m.PageSize > limit {
		return limit
	}
	return m.PageSize
}
