package model

import (
	"fmt"

	tm "github.com/nenecchuu/lizbeth-be-core/internal/app/token/model"
	um "github.com/nenecchuu/lizbeth-be-core/internal/app/user/model"
)

func BuildHandleLinkageCallbackReply(u *um.UserNoSqlSchema, t *tm.TokenNoSqlSchema) string {
	return fmt.Sprintf("Connected with username: %s and token: %s", u.Name, t.AccessToken)
}
