package urls

import "net/url"

func MustJoinPath(baseUrl string, elem ...string) string {
	res, err := url.JoinPath(baseUrl, elem...)
	if err != nil {
		panic(err)
	}

	return res
}
