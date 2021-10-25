package redis

// redis key注意使用命名空间的方式,方便查询和拆分
//使用:分隔命名空间（bluebell为项目名，post为和帖子相关的）
//投票是固定不变的，要定义成常量

const (
	//redis key定义
	//方法1：（给每个key加上注释说明常量类型为zset）
	//KeyPostTime        = "post:time"   // zset;贴子及发帖时间
	//方法2：（直接在定义常量名时写上类型zset）
	//KeyPostTimeZSet    = "post:time"

	Prefix             = "bluebell:"   // 项目key前缀
	KeyPostTimeZSet    = "post:time"   // zset;贴子及发帖时间
	KeyPostScoreZSet   = "post:score"  // zset;贴子及投票的分数
	KeyPostVotedZSetPF = "post:voted:" // zset;记录用户及投票类型;参数是post id  (PF代表不完整的key)

	KeyCommunitySetPF = "community:" // set;保存每个分区下帖子的id
)

// 给redis key加上前缀
func getRedisKey(key string) string {
	return Prefix + key
}
