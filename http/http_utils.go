package http_utils

type QueryParam struct {
	Key    string
	Values []string
}

func BuildQueryParam(params ...QueryParam) string {
	var rawQuery = ""
	if len(params) > 0 {
		rawQuery = "?"
	}
	for i := range params {
		for j := range params[i].Values {
			rawQuery += (params[i].Key + "=" + params[i].Values[j])
			if (i == (len(params) - 1) && j == (len(params[i].Values) - 1)) == false {
				rawQuery += "&"
			}
		}
	}
	return rawQuery
}