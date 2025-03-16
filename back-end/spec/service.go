package spec

type (
	CommentService interface {
		Insert(model CommentModel) error
		Find() ([]CommentModel, error)
	}
)
