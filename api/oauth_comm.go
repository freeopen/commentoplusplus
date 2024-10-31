package main

import (
	"os"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

// 兼容 code 授权 和 oidc 授权
var commConfig *oauth2.Config

func commOauthConfigure(brand string) error {
	commConfig = nil
	key := strings.ToUpper(brand) + "_KEY"
	secret := strings.ToUpper(brand) + "_SECRET"
	if os.Getenv(key) == "" && os.Getenv(secret) == "" {
		return nil
	}

	if os.Getenv(key) == "" {
		logger.Errorf("COMMENTO_%s not configured, but COMMENTO_%s is set", key, secret)
		return errorOauthMisconfigured
	}

	if os.Getenv(secret) == "" {
		logger.Errorf("COMMENTO_%s not configured, but COMMENTO_%s is set", secret, key)
		return errorOauthMisconfigured
	}
	logger.Infof("loading %s OAuth config", brand)

	switch brand {
	case "google":
		commConfig = &oauth2.Config{
			RedirectURL:  os.Getenv("ORIGIN") + "/api/oauth/" + brand + "/callback",
			ClientID:     os.Getenv(key),
			ClientSecret: os.Getenv(secret),
			Scopes:       []string{"openid", "profile", "email"},
			Endpoint:     google.Endpoint,
		}
		googleConfigured = true

	case "github":
		commConfig = &oauth2.Config{
			RedirectURL:  os.Getenv("ORIGIN") + "/api/oauth/" + brand + "/callback",
			ClientID:     os.Getenv(key),
			ClientSecret: os.Getenv(secret),
			Scopes: []string{
				"read:user",
				"user:email",
			},
			Endpoint: github.Endpoint,
		}

		githubConfigured = true

	case "douyin":
		commConfig = &oauth2.Config{
			RedirectURL:  os.Getenv("ORIGIN") + "/api/oauth/" + brand + "/callback",
			ClientID:     os.Getenv(key),
			ClientSecret: os.Getenv(secret),
			Scopes: []string{
				"login_id",
			},
			Endpoint: DouyinEP,
		}

		douyinConfigured = true

	case "wechat":
		commConfig = &oauth2.Config{
			RedirectURL:  os.Getenv("ORIGIN") + "/api/oauth/" + brand + "/callback",
			ClientID:     os.Getenv(key),
			ClientSecret: os.Getenv(secret),
			Scopes: []string{
				"snsapi_login",
			},
			Endpoint: WechatEP,
		}

		wechatConfigured = true

	case "qq":
		commConfig = &oauth2.Config{
			RedirectURL:  os.Getenv("ORIGIN") + "/api/oauth/" + brand + "/callback",
			ClientID:     os.Getenv(key),
			ClientSecret: os.Getenv(secret),
			Scopes: []string{
				"snsapi_login",
			},
			Endpoint: QqEP,
		}

		qqConfigured = true

	case "weibo":
		commConfig = &oauth2.Config{
			RedirectURL:  os.Getenv("ORIGIN") + "/api/oauth/" + brand + "/callback",
			ClientID:     os.Getenv(key),
			ClientSecret: os.Getenv(secret),
			Scopes: []string{
				"email",
			},
			Endpoint: WeiboEP,
		}

		weiboConfigured = true
	case "dingtalk":
		commConfig = &oauth2.Config{
			RedirectURL:  os.Getenv("ORIGIN") + "/api/oauth/" + brand + "/callback",
			ClientID:     os.Getenv(key),
			ClientSecret: os.Getenv(secret),
			Scopes: []string{
				"snsapi_login",
			},
			Endpoint: DingtalkEP,
		}

		dingtalkConfigured = true
	default:

	}

	return nil
}
