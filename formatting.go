package main

import (
	"fmt"
)

func (s Stack) String() string {
	str := ""
	
	
	for cursor := &s; cursor != nil; cursor = cursor.Next {
		if cursor.Kind == HEAD {
			continue
		}

		fmt.Println(cursor.Word)
	}
	
	return str
}

func (w Word) String() string {
	return fmt.Sprintf(`{"%v" → "%v"}`, w.Value, w.Kind)
}

func (t Token) String() string {
	return fmt.Sprintf(`{"%v" : %v}`, t.body, t.Kind)
}

func (k Kind) String() (s string) {
	switch k {
	case Integral:
		return "Integral"
	case Real:
		return "Real"
	case Procedure:
		return "Procedure"
	case String:
		return "String"
	case Variable:
		return "Variable"
	case NIL:
		return "NIL"
	default:
		return "UNKNOWN"
	}
	
	return
}
