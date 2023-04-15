package model

import (
	sm "github.com/nenecchuu/lizbeth-be-core/internal/app/session/model"
	tm "github.com/nenecchuu/lizbeth-be-core/internal/app/token/model"
	um "github.com/nenecchuu/lizbeth-be-core/internal/app/user/model"
)

type CommandMetadata struct {
	User      um.UserNoSqlSchema
	Session   sm.SessionNoSqlSchema
	HostToken tm.TokenNoSqlSchema
}
