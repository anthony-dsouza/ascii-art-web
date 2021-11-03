package ascii

func Map(arr []string) map[rune][]string {
	bannerMap := make(map[rune][]string)

	key := ' '
	for i := range arr {
		if i%9 == 0 && i != 0 {
			for j := i - 8; j < i; j++ {
				bannerMap[key] = append(bannerMap[key], arr[j])
			}
			key++
		}
	}

	return bannerMap
}
