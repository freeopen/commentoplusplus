package main

import (
	"golang.org/x/oauth2"
)

var DouyinEP = oauth2.Endpoint{
	AuthURL:  "https://open.douyin.com/platform/oauth/connect", // scope: login_id, user_info(nickname, avatar), mobile_alert(encrypt_mobile)
	TokenURL: "https://open.douyin.com/oauth/access_token",     // post x-form: client_key
	// userInfo: https://open.douyin.com/oauth/userinfo post json access_token, open_id
}

var XiaoHongShuEP = oauth2.Endpoint{
	AuthURL:  "https://ark.xiaohongshu.com/ark/authorization",
	TokenURL: "https://ark.xiaohongshu.com/ark/open_api/v3/common_controller",
}

var JdEP = oauth2.Endpoint{
	AuthURL:  "https://open-oauth.jd.com/oauth2/to_login",     // ?app_key=XXXXX&response_type=code&redirect_uri=XXXXX&state=20180416&scope=snsapi_base
	TokenURL: "https://open-oauth.jd.com/oauth2/access_token", // ?app_key=XXXXX&app_secret=XXXXX&grant_type=authorization_code&code=XXXXX
}

// JD
// {
//
//    "access_token": "b6787895ee524cf0b07c81f2566aafa1gm4o",
//    "expires_in": 31536000,
//    "refresh_token": "097r1b3d3f4e4c55a4e45d08d29aafe0otfi",
//    "scope": "snsapi_base",
//    "open_id": "KV-4zX_bLFUD6617rH7y0ipkr5n2-3uU6WCAnL4999U",
//    "uid": "5618242670",
//    "time": 1695200156425,
//    "token_type": "bearer",
//    "code": 0,
//    "xid": "o*AAQu5M9OMBphkmnhfTKWVl_IZTBkNsDX107fohhH6_DTli9rrUY"
// }

var AlipayEP = oauth2.Endpoint{
	AuthURL:  "https://openauth.alipay.com/oauth2/publicAppAuthorize.htm",
	TokenURL: "https://openauth.alipay.com/gateway.do",
}

var DingtalkEP = oauth2.Endpoint{
	AuthURL:  "https://login.dingtalk.com/oauth2/auth", // openid
	TokenURL: "https://api.dingtalk.com/v1.0/oauth2/userAccessToken",
}

// 需要企业资质或个人身份证
var QqEP = oauth2.Endpoint{
	AuthURL:  "https://graph.qq.com/oauth2.0/authorize",
	TokenURL: "https://graph.qq.com/oauth2.0/token",
	// get openid:  "https://graph.qq.com/oauth2.0/me",
	// userInfo: "https://graph.qq.com/user/get_user_info",
}

// 需要企业资质
var WechatEP = oauth2.Endpoint{
	AuthURL:  "https://open.weixin.qq.com/connect/qrconnect",
	TokenURL: "https://api.weixin.qq.com/sns/oauth2/access_token",
}

// sina weibo
var WeiboEP = oauth2.Endpoint{
	AuthURL:  "https://api.weibo.com/oauth2/authorize", // scope: email
	TokenURL: "https://api.weibo.com/oauth2/access_token",
	// userInfo https://api.weibo.com/2/users/show.json
}
