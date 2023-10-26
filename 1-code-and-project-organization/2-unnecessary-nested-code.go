package codeandprojectorganization

import "errors"

var s1 string
var s2 string
var max int

/*
é importante ficar atento ao quanto nosso código está se afastando da borda da esquerda
quanto mais longe estiver, normalmente mais esforço cognitivo temos que fazer para entender o código
*/
func BadExampleNestedCode() (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	} else {
		if s2 == "" {
			return "", errors.New("s2 is empty")
		} else {
			concat, err := concatenate(s1, s2)
			if err != nil {
				return "", err
			} else {
				if len(concat) > max {
					return concat[:max], nil
				} else {
					return concat, nil
				}
			}
		}
	}
}

func GoodExampleWithoutNestedCode() (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	}
	if s2 == "" {
		return "", errors.New("s2 is empty")
	}
	concat, err := concatenate(s1, s2)
	if err != nil {
		return "", err
	}
	if len(concat) > max {
		return concat[:max], nil
	}
	return concat, nil
}

func concatenate(str1, str2 string) (string, error) { return "", nil }
