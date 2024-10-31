package main

import (
	"golang.org/x/oauth2"
)

var DouyinEP = oauth2.Endpoint{
	AuthURL:  "https://open.douyin.com/oauth/authorize/v2", // get json scope: login_id, user_info(nickname, avatar), mobile_alert(encrypt_mobile)
	TokenURL: "https://open.douyin.com/oauth/access_token", // post x-form
	// userInfo: https://open.douyin.com/oauth/userinfo post json access_token, open_id
}

var XiaoHongShuEP = oauth2.Endpoint{
	AuthURL:  "https://ark.xiaohongshu.com/ark/authorization",
	TokenURL: "https://ark.xiaohongshu.com/ark/open_api/v3/common_controller",
}

var JdEP = oauth2.Endpoint{
	AuthURL:  "https://open-oauth.jd.com/oauth2/to_login",
	TokenURL: "https://open-oauth.jd.com/oauth2/access_token",
}

var AlipayEP = oauth2.Endpoint{
	AuthURL:  "https://openauth.alipay.com/oauth2/publicAppAuthorize.htm",
	TokenURL: "https://openauth.alipay.com/gateway.do",
}

var DingtalkEP = oauth2.Endpoint{
	AuthURL:  "https://login.dingtalk.com/oauth2/auth",
	TokenURL: "https://api.dingtalk.com/v1.0/oauth2/userAccessToken",
}

var QqEP = oauth2.Endpoint{
	AuthURL:  "https://graph.qq.com/oauth2.0/authorize",
	TokenURL: "https://graph.qq.com/oauth2.0/token",
	// get openid:  "https://graph.qq.com/oauth2.0/me",
	// userInfo: "https://graph.qq.com/user/get_user_info",
}

var WechatEP = oauth2.Endpoint{
	AuthURL:  "https://open.weixin.qq.com/connect/qrconnect",
	TokenURL: "https://api.weixin.qq.com/sns/oauth2/access_token",
}

var WeiboEP = oauth2.Endpoint{
	AuthURL:  "https://api.weibo.com/oauth2/authorize",
	TokenURL: "https://api.weibo.com/oauth2/access_token",
}
