// Package paginator 处理分页逻辑
package paginator

import (
	"fmt"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/configYaml"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

// Paging 分页数据
type Paging struct {
	CurrentPage int    // 当前页
	PerPage     int    // 每页条数
	TotalPage   int    // 总页数
	TotalCount  int64  // 总条数
	NextPageURL string // 下一页的链接
	PrevPageURL string // 上一个的链接
}

// Paginator 分页操作类
type Paginator struct {
	BaseURL    string // 用以拼接URL
	PerPage    int    // 每页条数
	Page       int    // 当前页
	Offset     int    // 数据库读取数据的 Offset 偏移量
	TotalCount int64  // 总条数
	TotalPage  int    // 总页数 = TotalCount / PerPage
	Sort       string // 排序规则
	Order      string // 排序顺序

	query *gorm.DB     // db query 句柄
	ctx   *gin.Context // gin.Context, 方便调用
}

// Paginate 分页
// c —— gin.context 用来获取分页的 URL 参数
// db —— GORM 查询句柄，用以查询数据集和获取数据总数
// baseURL —— 用以分页链接
// data —— 模型数组，传址获取数据
// PerPage —— 每页条数，优先从 url 参数里取，否则使用 perPage 的值
// 用法:
//
//	query := database.DB.Model(Topic{}).Where("category_id = ?", cid)
//	var topics []Topic
//	paging := paginator.Paginate(
//	    c,
//	    query,
//	    &topics,
//	    app.APIURL(database.TableName(&Topic{})),
//	    perPage,
//	)
func Paginate(c *gin.Context, db *gorm.DB, data interface{}, baseURL string, perPage int) Paging {

	// 初始化 Paginator 实例
	p := &Paginator{
		query: db,
		ctx:   c,
	}

	p.initProperties(perPage, baseURL)

	// 查询数据库
	err := p.query.Preload(clause.Associations). // 读取关联
							Order(p.Sort + " " + p.Order). // 排序
							Limit(p.PerPage).
							Offset(p.Offset).
							Find(data).Error

	// 数据库出错
	if err != nil {
		logger.LogIf(err)
		return Paging{}
	}

	return Paging{
		CurrentPage: p.Page,
		PerPage:     p.PerPage,
		TotalPage:   p.TotalPage,
		TotalCount:  p.TotalCount,
		NextPageURL: p.getNextPageURL(),
		PrevPageURL: p.getPrevPageURL(),
	}
}

// 初始化分页必须用到的属性，基于这些属性查询数据库
func (p *Paginator) initProperties(perPage int, baseURL string) {

	p.BaseURL = p.formatBaseURL(baseURL)
	p.PerPage = p.getPerPage(perPage)

	// 排序参数（控制器中以验证过这些参数，可放心使用）
	p.Sort = p.ctx.DefaultQuery(configYaml.Gohub_Config.Paging.UrlQuerySort, "id")
	p.Order = p.ctx.DefaultQuery(configYaml.Gohub_Config.Paging.UrlQueryOrder, "asc")

	p.TotalCount = p.getTotalCount()
	p.TotalPage = p.getTotalPage()
	p.Page = p.getCurrentPage()
	p.Offset = (p.Page - 1) * p.PerPage
}

// 兼容 URL 带与不带 `?` 的情况
func (p *Paginator) formatBaseURL(baseURL string) string {
	if strings.Contains(baseURL, "?") {
		baseURL = baseURL + "&" + configYaml.Gohub_Config.Paging.UrlQueryPage + "="
	} else {
		baseURL = baseURL + "?" + configYaml.Gohub_Config.Paging.UrlQueryPage + "="
	}
	return baseURL
}

// getPerPage 获取每页条数
func (p *Paginator) getPerPage(perPage int) int {
	// 优先使用请求 per_page 参数
	queryPerPage := p.ctx.Query(configYaml.Gohub_Config.Paging.UrlQueryPerPage)
	if len(queryPerPage) > 0 {
		perPage = cast.ToInt(queryPerPage)
	}

	// 没有传参, 使用默认页数
	if perPage <= 0 {
		perPage = configYaml.Gohub_Config.Paging.PerPage
	}

	return perPage
}

// getTotalCount 返回的是数据库里的条数
func (p *Paginator) getTotalCount() int64 {
	var count int64
	if err := p.query.Count(&count).Error; err != nil {
		return 0
	}
	return count
}

// getTotalPage 计算总页数
func (p *Paginator) getTotalPage() int {
	if p.TotalCount == 0 {
		return 0
	}

	nums := int64(float64(p.TotalCount) / float64(p.PerPage))
	if p.TotalCount%int64(p.PerPage) != 0 {
		nums++
	}

	return int(nums)
}

// getCurrentPage 返回当前页码
func (p *Paginator) getCurrentPage() int {
	// 优先取用户请求的 page
	page := cast.ToInt(p.ctx.Query(configYaml.Gohub_Config.Paging.UrlQueryPage))
	if page <= 0 {
		// 默认为 1
		page = 1
	}
	// TotalPage 等于 0 ，意味着数据不够分页
	if p.TotalPage == 0 {
		return 0
	}
	// 请求页数大于总页数，返回总页数
	if page > p.TotalPage {
		return p.TotalPage
	}
	return page
}

// 拼接分页链接
func (p *Paginator) getPageLink(page int) string {
	return fmt.Sprintf("%v%v&%s=%s&%s=%s&%s=%v",
		p.BaseURL,
		page,
		configYaml.Gohub_Config.Paging.UrlQuerySort,
		p.Sort,
		configYaml.Gohub_Config.Paging.UrlQueryOrder,
		p.Order,
		configYaml.Gohub_Config.Paging.UrlQueryPerPage,
		p.PerPage,
	)
}

// getNextPageURL 返回下一页的链接
func (p *Paginator) getNextPageURL() string {
	if p.TotalPage > p.Page {
		return p.getPageLink(p.Page + 1)
	}
	return ""
}

// getPrevPageURL 返回上一页的链接
func (p *Paginator) getPrevPageURL() string {
	if p.Page <= 1 || p.Page > p.TotalPage {
		return ""
	}
	return p.getPageLink(p.Page - 1)
}
