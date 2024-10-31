package main

var (
	googleConfigured      bool
	twitterConfigured     bool
	githubConfigured      bool
	gitlabConfigured      bool
	qqConfigured          bool
	wechatConfigured      bool
	douyinConfigured      bool
	dingtalkConfigured    bool
	xiaohongshuConfigured bool
	weiboConfigured       bool
)

func oauthConfigure() error {
	if err := googleOauthConfigure(); err != nil {
		return err
	}

	if err := twitterOauthConfigure(); err != nil {
		return err
	}

	if err := githubOauthConfigure(); err != nil {
		return err
	}

	if err := gitlabOauthConfigure(); err != nil {
		return err
	}

	return nil
}
