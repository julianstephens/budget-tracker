package router

type urls struct {
	HOME_PATH string
	USER_PATH string
}

func GetURLs() urls {
	var urlPatterns urls
	urlPatterns.HOME_PATH = "/"
	urlPatterns.USER_PATH = "/users/"
	return urlPatterns
}
