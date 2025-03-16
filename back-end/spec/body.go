package spec

type (
	CommentBody struct {
		Text      string `json:"text" bson:"text"`
		CreatedBy string `json:"created_by" bson:"created_by"`
	}
)
