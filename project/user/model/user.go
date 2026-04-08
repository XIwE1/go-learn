package model

type User struct {
	// uri 结构体标签将 URI 路径参数直接绑定到结构体中
	Name string `uri:"name" json:"name" binding:"required"`
	Id   int    `uri:"id" json:"id" binding:"required"`

	// `xx:"yy"` = 结构体标签（Struct Tag）。它是一种元数据（关于数据的数据），用来为结构体的字段提供额外的信息
	// Name string `json:"name"`
	// 当把 User 结构体转换成 JSON 字符串（序列化）时，字段 Name 在 JSON 中应该使用键名 "name"，而不是默认的字段名 "Name"。
	// 类似的还有
	// gorm:"column:user_name;type:varchar(100)" 用来指定数据库表中的列名和类型
}
