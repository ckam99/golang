package helper

import (
	"context"
	"example/grpc/pkg/security"
	"os"
)

func Authorized(ctx context.Context) error {
	mtdt := GetMetaData(ctx)
	_, err := security.VerifyToken(mtdt.Authorization, os.Getenv("SECRET_KEY"))
	if err != nil {
		return err
	}
	return nil
}
