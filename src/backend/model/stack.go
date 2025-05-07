package model

type Stack struct {
	Val  *Tree
	Next *Stack
}

func Push(s *Stack, t *Tree) {
	var temp *Stack = s
	s = &Stack{
		Val:  t,
		Next: temp,
	}
}
func Pop(s **Stack) {
	if *s != nil {
		*s = (*s).Next
	}
}

func StackLength(s *Stack) int64 {
	var count int64 = 0
	for s != nil {
		count++
		s = s.Next
	}
	return count
}
