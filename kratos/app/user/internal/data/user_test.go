package data

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
	v1 "user/api/user/v1"
	"user/internal/biz"
)

var userClient v1.UserClient
var conn *grpc.ClientConn

// Init 初始化 grpc 链接
func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("grpc link err" + err.Error())
	}
	userClient = v1.NewUserClient(conn)
}

// 测试创建用户接口
func TestCreateUser(t *testing.T) {
	Init()
	mockUser := biz.MockUserRepo{}
	mockUser.On("CreateUser", mock.Anything, &biz.User{
		Mobile:   fmt.Sprintf("1388888888%d", 1),
		NickName: fmt.Sprintf("YWWW%d", 1),
	}).Return(&biz.User{
		Mobile:   fmt.Sprintf("1388888888%d", 1),
		NickName: fmt.Sprintf("YWWW%d", 1),
	}, nil)

	rsp, err := mockUser.CreateUser(context.Background(), &biz.User{
		Mobile:   fmt.Sprintf("1388888888%d", 1),
		NickName: fmt.Sprintf("YWWW%d", 1),
	})

	// 断言期望的返回值
	assert.NoError(t, err)
	assert.Equal(t, "13888888881", rsp.Mobile)
	assert.Equal(t, "YWWW1", rsp.NickName)

	// 断言模拟对象的期望行为是否被满足
	mockUser.AssertExpectations(t)
}
