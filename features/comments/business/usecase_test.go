package business

import (
	"antonio/features/comments"
	"antonio/features/contents/mocks"
	"testing"
)

var (
	commentService 		comments.Business
	commentRepository	mocks.Data
	commentData			comments.Core
	commentsData		comments.Core
)

func TestMain(m *testing.M){
	
}