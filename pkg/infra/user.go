package infra

import (
	"context"

	"github.com/wolfmagnate/ai-friendly-example/pkg/domain/entity"
)

type UserRepository interface {
	// userUIDも型では守れないのでstringである
	// 対応するuserが存在しないときはnil, nilを返す
	GetUserByUserUID(ctx context.Context, tx Tx, userUID string) (*entity.User, error)
	CreateUser(ctx context.Context, tx Tx, userUID string) (*entity.User, error)
	DeleteUser(ctx context.Context, tx Tx, userID entity.UserID) error
}

type UserRepositoryImpl struct {
	Cache Cache
}

var _ UserRepository = UserRepositoryImpl{}
