package main

// https://leetcode.com/problems/isomorphic-strings/description/
// O(n) - time, O(1) - space
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s2t, t2s := map[byte]byte{}, map[byte]byte{}

	for i := 0; i < len(s); i++ {

		sCh, tCh := s[i], t[i]

		if _, ok := s2t[sCh]; !ok {
			s2t[sCh] = tCh
		} else if s2t[sCh] != tCh {
			return false
		}

		if _, ok := t2s[tCh]; !ok {
			t2s[tCh] = sCh
		} else if t2s[tCh] != sCh {
			return false
		}
	}

	return true
}
