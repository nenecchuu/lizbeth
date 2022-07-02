package model

import (
	"strings"

	sam "github.com/nenecchuu/lizbeth-be-core/internal/integration/spotify_api/model"
	"github.com/nenecchuu/lizbeth-be-core/internal/util"
)

type LinkageCallback struct {
	State string
	Code  string
}

func (x *LinkageCallback) ValidateLinkageCallback() error {
	if x.State == "" {
		return util.NewRequiredFieldErr("state")
	}

	return nil
}

func (x *LinkageCallback) ToSpotifyAuthorizeData() *sam.SpotifyAuthorizeData {
	return &sam.SpotifyAuthorizeData{
		Code: x.Code,
	}
}

func (x *LinkageCallback) ParseFromChatMessage(msg string) *LinkageCallback {
	splmsg := strings.Split(msg, ":")
	return &LinkageCallback{
		State: splmsg[0],
		Code:  splmsg[1],
	}
}
