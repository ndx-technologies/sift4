package sift4

import "slices"

// Buffer stores temporary structures and can be reused in multiple runs.
// Zero value is safe to use.
type Buffer struct {
	offset []offset
}

type offset struct {
	c1    int
	c2    int
	trans bool
}

// Distance is fast O(N) string distance algorithm.
// If pointer to buffer is provided, then it will be reused for storing temporary structures.
//   - maxOffset is the number of characters to search for matching letters
//   - maxDistance is the distance at which the algorithm should stop computing the value and just exit (the strings are too different anyway)
//
// https://siderite.dev/blog/super-fast-and-accurate-string-distance.html
func Distance(s1, s2 string, maxOffset, maxDistance int, buffer *Buffer) int {
	if len(s1) == 0 {
		return len(s2)
	}
	if len(s2) == 0 {
		return len(s1)
	}
	if s1 == s2 {
		return 0
	}

	var (
		c1, c2  int // cursor for string 1 and 2
		lcss    int // largest common subsequence
		localCS int // local common substring
		trans   int // number of transpositions
	)

	if buffer == nil {
		buffer = &Buffer{}
	}
	if buffer.offset == nil {
		buffer.offset = make([]offset, 0, min(maxOffset, len(s1), len(s2)))
	}
	buffer.offset = buffer.offset[:0]

	for (c1 < len(s1)) && (c2 < len(s2)) {
		if s1[c1] == s2[c2] {
			isTrans := false

			localCS++

			for i := 0; i < len(buffer.offset); {
				ofs := buffer.offset[i]
				if c1 <= ofs.c1 || c2 <= ofs.c2 {
					isTrans = abs(c2-c1) >= abs(ofs.c2-ofs.c1)
					if isTrans {
						trans++
					} else if !ofs.trans {
						ofs.trans = true
						trans++
					}
					break
				} else {
					if c1 > ofs.c2 && c2 > ofs.c1 {
						buffer.offset = slices.Delete(buffer.offset, i, i+1)
					} else {
						i++
					}
				}
			}

			buffer.offset = append(buffer.offset, offset{c1, c2, isTrans})
		} else {
			lcss += localCS
			localCS = 0

			if c1 != c2 {
				c1 = min(c1, c2) // using min allows the computation of transpositions
				c2 = c1
			}

			// if matching characters are found, remove 1 from both cursors (they get incremented at the end of the loop)
			// so that we can have only one code block handling matches
			for i := 0; i < maxOffset && (c1+i < len(s1) || c2+i < len(s2)); i++ {
				if c1+i < len(s1) && s1[c1+i] == s2[c2] {
					c1 += i - 1
					c2--
					break
				}
				if c2+i < len(s2) && s1[c1] == s2[c2+i] {
					c1--
					c2 += i - 1
					break
				}
			}
		}

		c1++
		c2++

		if maxDistance > 0 {
			if d := max(c1, c2) - lcss + trans; d > maxDistance {
				return d
			}
		}

		// this covers the case where the last match is on the last token in list, so that it can compute transpositions correctly
		if c1 >= len(s1) || c2 >= len(s2) {
			lcss += localCS
			localCS = 0
			c1 = min(c1, c2)
			c2 = c1
		}
	}

	lcss += localCS
	return max(len(s1), len(s2)) - lcss + trans // add the cost of transpositions to the final result
}

func abs[T int](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
