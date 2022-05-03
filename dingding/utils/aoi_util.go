package utils

import (
	"net/url"
	"sort"
	"strings"
)

// 组合query到url里
func PackUrl(requestUrl string, queries ...map[string]string) string {
	var suffixArray = make([]string, 0)
	var suffix = ""
	for _, query := range queries {
		if query == nil {
			continue
		}
		var keys = make([]string, 0)
		for k := range query {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			suffixArray = append(suffixArray, k+"="+query[k])
		}
	}
	suffix = strings.Join(suffixArray, "&")
	if suffix != "" {
		suffix = "?" + suffix
	}

	return requestUrl + suffix
}

// 组合登录请求的post body
func PackPostBody(params map[string]string) url.Values {
	var postBody = make(url.Values)
	if params != nil {
		for k, v := range params {
			postBody.Set(k, v)
		}
	}
	return postBody
}
