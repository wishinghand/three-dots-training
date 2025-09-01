package main

func Merge(a []string, b []string) []string {
	return append(a, b...)
}

func Remove(a []string, index int) []string {
	if(index == 0) {
		return a[1:]
	}
	return append(a[:index], a[index+1:]...)
}

func RemoveLast(a []string) []string {
	return a[0:len(a) - 1]
}
