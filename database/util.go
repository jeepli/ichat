package database

func ArrayToSequence(strs []string) string {
	seq := ""
	if len(strs) == 0 {
		return seq
	}

	seq += strs[0]
	for i := 1; i < len(strs); i++ {
		seq += ","
		seq += strs[i]
	}
	return seq
}
