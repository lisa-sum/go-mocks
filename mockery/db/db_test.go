package db

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// Mock无参方法
// 在DB interface上有一个不带参数的方法FetchDefaultMessage,
// 可以像下面这样创建一个模拟对象:
func TestMockMethodWithoutArgs(t *testing.T) {
	// create the mock
	theDBMock := MockDB{}
	// mock the expectation
	theDBMock.On("FetchDefaultMessage").Return("yes", nil)
	// 使用模拟数据库创建 问候 对象
	g := greeter{&theDBMock, "en"}

	// 正确地断言测试
	assert.Equal(t, "Message is: yes", g.GreetInDefaultMsg())

	// 错误地断言测试
	// assert.Equal(t, "Message is: no 这不符合断言的值yes", g.GreetInDefaultMsg())

	// 可以断言被模拟的方法将被调用多少次
	theDBMock.AssertNumberOfCalls(t, "FetchDefaultMessage", 1)
}

// Mock带参数的方法
// 在DB interface上有一个带参数的方法FetchMessage(lang string),
// 可以像下面这样创建一个模拟对象
func TestMockMethodWithArgs(t *testing.T) {
	theDBMock := MockDB{}
	// if FetchMessage("sg") is called, then return "lah"
	theDBMock.On("FetchMessage", "sg").Return("lah", nil)

	// 正确的断言
	g := greeter{&theDBMock, "sg"}

	// 错误地断言测试
	// g := greeter{&theDBMock, "ch"}

	assert.Equal(t, "Message is: lah", g.Greet())
	theDBMock.AssertExpectations(t)
}

// Mock带参数的方法, 但是参数具体内容非测试重点
// 有时我们想模拟一个方法，但我们不在乎传递的实际参数。为此，我们可以在On()方法参数后面的第二个参数中使用mock.Anything。
func TestMockMethodWithArgsIgnoreArgs(t *testing.T) {
	theDBMock := MockDB{}
	// if FetchMessage(...) is called with any argument, please also return lah
	theDBMock.On("FetchMessage", mock.Anything).Return("lah", nil)
	g := greeter{&theDBMock, "in"}
	assert.Equal(t, "Message is: lah", g.Greet())
	theDBMock.AssertCalled(t, "FetchMessage", "in")
	theDBMock.AssertNotCalled(t, "FetchMessage", "ch")
	theDBMock.AssertExpectations(t)
	mock.AssertExpectationsForObjects(t, &theDBMock)
}

// Mock带参数的方法, 并校验实际参数
// 如果需要模拟一个具有复杂参数的方法，但希望根据参数的某些属性或从中进行计算来匹配mock。
// 例如，我们想模仿FetchMessage方法，但前提是lang参数以字母i开头。
func TestMatchedBy(t *testing.T) {
	theDBMock := MockDB{}
	theDBMock.On("FetchMessage", mock.MatchedBy(func(lang string) bool { return lang[0] == 'i' })).Return("bzzzz", nil) // all of these call FetchMessage("iii"), FetchMessage("i"), FetchMessage("in") will match

	// 成功地断言测试
	g := greeter{&theDBMock, "izz"}

	// 失败地断言测试, 因为lang参数不是字母i开头。
	// g := greeter{&theDBMock, "zz"}

	msg := g.Greet()
	assert.Equal(t, "Message is: bzzzz", msg)
	theDBMock.AssertExpectations(t)
}
