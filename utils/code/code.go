package code

import "container/list"

var baseStrLists string = "0A1B2C3D4E5F6G7H8J9KLMNPQRSTUVWXYZ"
var base []byte = []byte(baseStrLists)

func InviteCode(n uint32) string {
	quotient := n
	mod := uint32(0)
	l := list.New()
	for quotient != 0 {
		mod = quotient % 34
		quotient = quotient / 34
		l.PushFront(base[int(mod)])
	}
	listLen := l.Len()
	if listLen >= 6 {
		res := make([]byte, 0, listLen)
		for i := l.Front(); i != nil; i = i.Next() {
			res = append(res, i.Value.(byte))
		}
		return string(res)
	} else {
		res := make([]byte, 0, 6)
		for i := 0; i < 6; i++ {
			if i < 6-listLen {
				res = append(res, base[0])
			} else {
				res = append(res, l.Front().Value.(byte))
				l.Remove(l.Front())
			}

		}
		return string(res)
	}
}
