package graphdb

// Scramble a key in a way that may encourage fast access in a blob store.
func ScrambleKey(key string) string {
	// Revese the key because prefixes are common.
  runeskey := []rune(key)
  for i, j := 0, len(runeskey)-1; i < j; i, j = i+1, j-1 {
      runeskey[i], runeskey[j] = runeskey[j], runeskey[i]
  }
	key = string(runeskey)

	// Keys longer than 5, split up into directories.
	if len(key) > 5 {
		// Split the string into directories to improve searches.
		key = key[0:2] + "/" + key[2:4] + "/" + key[4:]
	}

	return key
}
