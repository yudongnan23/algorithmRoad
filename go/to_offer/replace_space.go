package to_offer

const (
	_spaceReplaceStr = "%20"
)

func replaceSpace(s string) string {
	var resultStr string
	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			resultStr = resultStr + _spaceReplaceStr
			continue
		}
		resultStr = resultStr + string(s[i])
	}

	return resultStr
}
