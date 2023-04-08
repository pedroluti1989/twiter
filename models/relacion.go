package models

type Relacion struct {
	UserID         string `bson:"userid, omitempty" json:"userId"`
	UserRelacionID string `bson:"userrelacionid, omitempty" json:"userRelacionId"`
}
