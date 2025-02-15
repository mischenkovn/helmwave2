package action_test

import (
	"testing"

	"github.com/helmwave/helmwave/pkg/action"
	"github.com/stretchr/testify/suite"
)

type ListTestSuite struct {
	suite.Suite
}

func (ts *ListTestSuite) TestImplementsAction() {
	ts.Require().Implements((*action.Action)(nil), &action.List{})
}

func TestListTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ListTestSuite))
}
