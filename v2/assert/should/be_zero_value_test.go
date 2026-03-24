package should_test

import (
	"testing"

	"github.com/smarty/gunit/v2/assert/should"
)

func TestShouldBeZeroValue(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.BeZeroValue, "EXTRA")

	assert.Pass(nil, should.BeZeroValue)
	assert.Pass(0, should.BeZeroValue)
	assert.Pass("", should.BeZeroValue)
	assert.Pass(false, should.BeZeroValue)
	assert.Pass(0.0, should.BeZeroValue)
	assert.Pass(complex(0, 0), should.BeZeroValue)
	assert.Pass(struct{}{}, should.BeZeroValue)
	assert.Pass((*string)(nil), should.BeZeroValue)
	assert.Pass([]string(nil), should.BeZeroValue)
	assert.Pass(map[string]string(nil), should.BeZeroValue)
	assert.Pass((chan string)(nil), should.BeZeroValue)
	assert.Pass((func())(nil), should.BeZeroValue)
	assert.Pass([0]string{}, should.BeZeroValue)
	assert.Pass([3]int{}, should.BeZeroValue)

	assert.Fail(1, should.BeZeroValue)
	assert.Fail("hi", should.BeZeroValue)
	assert.Fail(true, should.BeZeroValue)
	assert.Fail(&struct{}{}, should.BeZeroValue)
	assert.Fail([]string{}, should.BeZeroValue)
	assert.Fail(map[string]string{}, should.BeZeroValue)
	assert.Fail(struct{ X int }{X: 1}, should.BeZeroValue)
}

func TestShouldNotBeZeroValue(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.NOT.BeZeroValue, "EXTRA")

	assert.Fail(nil, should.NOT.BeZeroValue)
	assert.Fail(0, should.NOT.BeZeroValue)
	assert.Fail("", should.NOT.BeZeroValue)
	assert.Fail(false, should.NOT.BeZeroValue)
	assert.Fail((*string)(nil), should.NOT.BeZeroValue)
	assert.Fail([]string(nil), should.NOT.BeZeroValue)

	assert.Pass(1, should.NOT.BeZeroValue)
	assert.Pass("hi", should.NOT.BeZeroValue)
	assert.Pass(true, should.NOT.BeZeroValue)
	assert.Pass(&struct{}{}, should.NOT.BeZeroValue)
	assert.Pass([]string{}, should.NOT.BeZeroValue)
	assert.Pass(map[string]string{}, should.NOT.BeZeroValue)
}
