package utils

import "net/url"

func Map2URLValues(from map[string]string) url.Values {
	res := make(url.Values, len(from))
	for k, v := range from {
		res.Add(k, v)
	}
	return res
}
