package util_test

import (
	"testing"
	"urlshorterner/internal/pkg/util"

	"github.com/stretchr/testify/assert"
)

func TestPaddingLeadingZero(t *testing.T) {
	expected := "0000011111"
	assert.Equal(t, expected, util.PaddingLeadingZero(11111))
}

func TestRemoveLeadingZero_Normal(t *testing.T) {
	var expected uint = 1
	num, err := util.RemoveLeadingZero("0000000001")
	assert.Nil(t, err)
	assert.Equal(t, expected, num)
}

func TestRemoveLeadingZero_NegativeNumber(t *testing.T) {
	var expected uint = 0
	num, err := util.RemoveLeadingZero("-0000000001")
	assert.NotNil(t, err)
	assert.Equal(t, expected, num)
}

func TestRemoveLeadingZero_NotNumber(t *testing.T) {
	var expected uint = 0
	num, err := util.RemoveLeadingZero("00abc")
	assert.NotNil(t, err)
	assert.Equal(t, expected, num)
}
