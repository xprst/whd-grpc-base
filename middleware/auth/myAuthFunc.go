package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/xprst/whd-grpc-base/util"
	"strings"
)

var headerOpenID = "openid"

func AuthFunction(ctx context.Context) (context.Context, error) {
	// 若请求头中包含openid，则不在解析token
	if val := util.NiceMD(ctx).Get(headerOpenID); val != "" {
		return ctx, nil
	}

	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	parseToken(token)
	return ctx, nil
}

type identity struct {
	openid   string // OpenId
	createOn int32  // 身份生成的时间戳(ms)
	expireIn int32  // 身份过期的时间间隔(s)
	scope    string // 授权范围： ,分隔的字符串
}

func parseToken(token string) {
	token, _ = fromWebSafeBase64(token)
	//if err != nil {
	//	return nil, err
	//}

	pkcs := util.NewPkcs(util.DEFAULT_RSA_PUBLIC_KEY, util.DEFAULT_RSA_PRIVATE_KEY)
	str, _ := pkcs.Decrypt([]byte(token))
	fmt.Println(str)
}

// 从query string 安全的base64字符串还原成常规base64
func fromWebSafeBase64(base64Str string) (string, error) {
	// 将URL 安全的 - 到 +, _ 到 /
	base64Str = strings.Replace(base64Str, "-", "+", -1)
	base64Str = strings.Replace(base64Str, "_", "/", -1)

	// 补齐 = 号
	var missingPaddingCharacters = 0
	switch len(base64Str) % 4 {
	case 3:
		missingPaddingCharacters = 1
		break
	case 2:
		missingPaddingCharacters = 2
		break
	case 0:
		missingPaddingCharacters = 0
		break
	default:
		return "", errors.New("invalid web safe base64 format")
	}
	for i:=0; i<missingPaddingCharacters; i++ {
		base64Str += "="
	}

	return base64Str, nil
}

