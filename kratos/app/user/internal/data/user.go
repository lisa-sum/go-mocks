package data

import (
	"context"
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
	"user/internal/biz"
)

type User struct {
	ID          int64  `gorm:"primaryKey;not null"`
	Mobile      string `gorm:"index:idx_mobile;unique;type:varchar(11) comment '手机号码, 用户唯一标识';not null"`
	Password    string `gorm:"type:varchar(25);not null"`
	NickName    string `gorm:"type:varchar(25)"`
	Birthday    int64  `gorm:"type:datetime comment '出生日期'"`
	Gender      string `gorm:"column:gender;default:male,type:varchar(16) comment 'male: 男, female: 女'"`
	Role        int    `gorm:"column:role;default:1;type:int comment '1:普通用户, 2:管理员'"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	IsDeletedAt bool
}

type DB interface {
	FetchMessage(lang string) (string, error)
	FetchDefaultMessage() (string, error)
}

type UserRepo struct {
	repo *Data
	log  *log.Helper
}

func (ur UserRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	var user User
	result := ur.repo.db.
		Where(&biz.User{Mobile: u.Mobile}).
		First(&user)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}

	user.Mobile = u.Mobile
	user.NickName = u.NickName
	user.Password = encrypt(u.Password)

	res := ur.repo.db.Create(&user)
	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}

	return &biz.User{
		ID:       user.ID,
		Mobile:   user.Mobile,
		Password: user.Password,
		NickName: user.NickName,
		Birthday: user.Birthday,
		Gender:   user.Gender,
		Role:     user.Role,
	}, nil
}

func NewUserRepo(repo *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

// Password encryption
func encrypt(psd string) string {
	options := &password.Options{SaltLen: 16, Iterations: 10000, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode(psd, options)
	return fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
}
