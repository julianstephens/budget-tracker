package router

type urls struct {
	HOME_PATH string
	USER_PATH string
}

func GetURLs() urls {
	var url_patterns urls
	url_patterns.HOME_PATH = "/"
	url_patterns.USER_PATH = "/users/"
	return url_patterns
}
