package requests

type PageBase struct {
	Page  uint64 `json:"page,default=1"`
	Limit uint64 `json:"limit,default=10"`
	Off   bool   `json:"off,default=false"`
}

func (c *PageBase) ParsePage() (uint64, uint64) {
	if c.Page == 0 {
		c.Page = 1
	}
	if c.Limit == 0 || c.Limit > 100 {
		c.Limit = 10
	}
	// 关闭分页
	if c.Off == true {
		c.Limit = 1000000
	}
	return (c.Page - 1) * c.Limit, c.Limit
}
