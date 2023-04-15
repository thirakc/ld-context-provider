package ldcontext

type LDContextArg struct {
	Url         string `json:"url" bson:"url"`
	DocumentUrl string `json:"documentUrl" bson:"documentUrl"`
	Content     any    `json:"content" bson:"content"`
}

type LDContextEntity struct {
	Id string `bson:"_id"`
	*LDContextArg
}
