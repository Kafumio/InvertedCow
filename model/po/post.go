package po

import "gorm.io/gorm"

// 发布流程：
// 1. 用户选择发布动态
// 2. 收到返回的Token，动态state处于发布中
// 3. 客户端使用token上传文件
// 4. 上传完成后，source触发回调，关联视频-动态，同时更改动态state为已发布
// 5. source返回动态发布地址/发布成功事件，客户端接收 | 客户端轮询动态state
// ! 设置超时为发布失败，state为发布失败
// ! other...
// TODO: 上传图片：多个数量，可能需要维护一个cnt，来保证所有图片上传后才更改为发布成功。如果上传失败。。。。。

// Post 动态
type Post struct {
	gorm.Model
	OriginUrl  string `gorm:"column:origin_url;comment:唯一标识，用作地址链接" json:"origin_url"` // 动态原始地址, 接一个唯一 随机 UID
	State      int    `gorm:"column:state;comment:1发布中 2发布成功 3发布失败" json:"state"`
	Publisher  int64  `gorm:"column:publisher;comment:发布者" json:"publisher"`        // 发布者
	OriginText string `gorm:"column:origin_text;comment:原始动态内容" json:"origin_text"` // 原始动态内容，后期可能做转义处理
	// TODO: tags, like, star, forward
	// other list
}
