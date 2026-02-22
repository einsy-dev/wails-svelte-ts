package utils

func FilterUnique(arr []string) []string {
	u := make(map[string]bool)
	var r []string

	for _, item := range arr {
		if !u[item] {
			u[item] = true
			r = append(r, item)
		}
	}
	return r
}
