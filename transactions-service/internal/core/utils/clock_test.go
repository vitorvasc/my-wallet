package utils

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ClockTestSuite struct {
	suite.Suite
}

func TestClock(t *testing.T) {
	suite.Run(t, new(ClockTestSuite))
}

func (suite *ClockTestSuite) TestNow() {
	clock := NewClock()
	now := clock.Now()
	suite.NotZero(now)
}
