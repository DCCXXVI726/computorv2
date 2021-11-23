package main

func checkVar(str string) (string, int){
	i := 0
	for i = 0; i< len(str); i++ {
		if str[i] == ' ' || str[i] == '\t' || str[i] == '\v'{
			i++
		}
		break
	}
	result :=  make([]rune,0)
	k := 0
	for k = i; k < len(str); k++ {
		if str[k] >= 'a' && str[k] <= 'z' {
			i++
			result = append(result,rune(str[k]))
		} else if str[k] >= 'A' && str[k] <= 'Z'{
			result = append(result,rune(str[k] - 'A' + 'a'))
		} else {
			break
		}
	}
	if len(result) == 1 && result[0] == 'i' {
		return "", 0
	}
	return string(result), k-1
}
