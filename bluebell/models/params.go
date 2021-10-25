package models

//定义与请求相关的结构体

const (
	OrderTime  = "time"
	OrderScore = "score"
)

//定义注册请求的参数结构体
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"` //使用 binding tag标识相关校验规则
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

//定义登录请求的参数结构体
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//投票数据
type ParamVoteData struct {
	// UserID 从请求中获取当前的用户
	PostID    string `json:"post_id" binding:"required"`              // 贴子id
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` // 赞成票(1)还是反对票(-1)取消投票(0)
	//这里若写required，当填入0时会报错，因为required默认会取出0值或布尔值为false的
}

// ParamPostList 获取帖子列表query string参数
type ParamPostList struct {
	CommunityID int64  `json:"community_id" form:"community_id"`   // 可以为空,若为空，则默认按照所有id排序，不涉及社区
	Page        int64  `json:"page" form:"page" example:"1"`       // 页码
	Size        int64  `json:"size" form:"size" example:"10"`      // 每页数据量
	Order       string `json:"order" form:"order" example:"score"` // 排序依据
}

//example用来在生成swagger文档时作为预填值

////按社区获取帖子列表query string参数(已和上方结构体合并)
//type ParamCommunityPostList struct {
//	*ParamPostList
//	CommunityID int64 `json:"community_id" form:"community_id"`
//}
