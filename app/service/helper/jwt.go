package helper

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang-module/carbon/v2"
	"kwd/app/helper/crypt"
	"kwd/app/helper/str"
	"kwd/kernel/app"
	"time"
)

func JwtToken(id any) string {
	now := carbon.Now()
	return crypt.Md5(fmt.Sprintf("%s%v%d%s", app.Cfg.Server.Name, id, now.Timestamp(), str.Random(8)))
}

func CheckJwt(ctx context.Context, channel string, claims jwt.RegisteredClaims) bool {

	result, _ := app.Redis.Exists(ctx, Blacklist(channel, "login", claims.ID)).Result()

	return result != 1
}

func BlackJwt(ctx context.Context, channel string, claims jwt.RegisteredClaims) bool {

	now := carbon.Now()

	expires := time.Duration(claims.ExpiresAt.Unix()+12*60*60-now.Timestamp()) * time.Second

	_, err := app.Redis.Set(ctx, Blacklist(channel, "login", claims.ID), now.ToDateTimeString(), expires).Result()

	if err != nil {
		return false
	}

	return true
}

func Blacklist(channel string, method string, str string) string {
	return app.Cfg.Server.Name + ":" + channel + ":blacklist:" + method + ":" + str
}
