package main

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	d := make(map[string]string, len(s))
	d1 := make(map[string]string, len(s))
	for index, ch := range s {
		if v, ok := d[string(ch)]; ok && v != string(t[index]) {
			return false
		}

		if v, ok := d1[string(t[index])]; ok && v != string(ch) {
			return false
		}
		d[string(ch)] = string(t[index])
		d1[string(t[index])] = string(ch)
	}
	return true
}
