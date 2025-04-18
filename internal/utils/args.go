package utils

func ParseArgs(args []string) (pattern string, file_list []string) {
	if index := indexOf(args, "--"); index != -1 {
		return args[0], args[index+1:]
	}

	return args[0], []string{}
}

func indexOf(arr []string, str string) (index int) {
	for i, v := range arr {
		if v == str {
			return i
		}
	}

	return -1
}
