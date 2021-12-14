package misc

func IndexOf(sl []string, tf string) (index int, err error) {
	for index, value := range sl {
		if tf == value {
			return index, nil
		}
	}
	return -1, err
}
