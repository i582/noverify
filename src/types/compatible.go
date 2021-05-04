package types

import (
	"fmt"
	"strings"
)

// CompatibleTypes T1 is Wanted, T2 is Gotten
func CompatibleTypes(T1, T2 Map) (ok bool, desc string) {
	if T1.Empty() || T2.Empty() {
		return true, ""
	}

	ok, desc = compatibleOneWithMany(T1, T2)
	if !ok {
		return false, desc
	}

	ok, desc = compatibleOneWithMany(T2, T1)
	if !ok {
		return false, desc
	}

	ok, desc = compatibleManyWithMany(T1, T2)
	if !ok {
		return false, desc
	}

	return true, ""
}

func compatibleManyWithMany(T1 Map, T2 Map) (ok bool, desc string) {
	if T1.Len() == 1 || T2.Len() == 1 {
		return true, ""
	}

	if T1.Contains("null") && !T2.Contains("null") {
		return false, fmt.Sprintf("cannot use nullable %s as %s", T1, T2)
	}
	if !T1.Contains("null") && T2.Contains("null") {
		return false, fmt.Sprintf("cannot use %s as nullable %s", T1, T2)
	}

	var compatibleWithOne bool

	T1.Iterate(func(T1Typ string) {
		T2.Iterate(func(T2Typ string) {
			ok, _ = CompatibleType(T1Typ, T2Typ)
			if ok {
				compatibleWithOne = true
			}
		})
	})

	return compatibleWithOne, ""
}

func compatibleOneWithMany(T1 Map, T2 Map) (ok bool, desc string) {
	if T1.Len() == 1 {
		if T2.Len() == 1 {
			return CompatibleType(T1.String(), T2.String())
		}

		T1S := T1.String()

		if strings.Contains(T1S, "mixed") {
			return true, ""
		}

		T2IsNullable := T2.Find(func(typ string) bool {
			return typ == "null"
		})
		if T2IsNullable {
			return false, fmt.Sprintf("cannot use type %s as nullable type %s", T1S, T2.String())
		}

		var compatibleWithOne bool

		T2.Iterate(func(typ string) {
			ok, desc = CompatibleType(T1S, typ)
			if ok {
				compatibleWithOne = true
			}
		})

		if !compatibleWithOne {
			return false, fmt.Sprintf("none of the possible types (%s) are compatible with %s", T2.String(), T1S)
		}
	}

	return true, ""
}

func CompatibleType(T1, T2 string) (ok bool, desc string) {
	if T1 == "mixed" || T2 == "mixed" {
		return true, ""
	}

	if IsPOD(T1) && IsPOD(T2) {
		if compatibleBoolean(T1, T2) {
			return true, ""
		}

		if T1 != T2 {
			return false, fmt.Sprintf("cannot use %s as %s", T1, T2)
		}
	}

	T1A, ok := Alias(T1)
	if ok {
		T1 = T1A
	}
	T2A, ok := Alias(T1)
	if ok {
		T2 = T2A
	}

	ok, desc = compatibleClass(T1, T2)
	if !ok {
		return false, desc
	}

	ok, desc = compatibleArray(T1, T2)
	if !ok {
		return false, desc
	}

	ok, desc = compatibleCallable(T1, T2)
	if !ok {
		return false, desc
	}

	return true, ""
}

func compatibleClass(T1 string, T2 string) (ok bool, desc string) {
	if IsClass(T1) && IsClass(T2) {
		if T1 != T2 {
			return false, fmt.Sprintf("cannot use class %s as class %s", T1, T2)
		}

		return true, ""
	}

	if IsClass(T1) {
		if T2 == "object" {
			return true, ""
		}

		return false, fmt.Sprintf("cannot use class %s as %s", T1, T2)
	}

	if IsClass(T2) {
		if T1 == "object" {
			return true, ""
		}

		return false, fmt.Sprintf("cannot use %s as class %s", T1, T2)
	}

	return true, ""
}

func compatibleArray(T1 string, T2 string) (ok bool, desc string) {
	if IsArray(T1) && IsArray(T2) {
		T1El := ArrayElementType(T1)
		T2El := ArrayElementType(T2)

		comp, des := CompatibleType(T1El, T2El)
		if !comp {
			return false, fmt.Sprintf("cannot use array of %s as array of %s: %s", T1El, T2El, des)
		}

		return true, ""
	}

	if IsArray(T1) {
		if T2 == "iterable" {
			return true, ""
		}

		return false, fmt.Sprintf("cannot use array of %s as %s", T1, T2)
	}

	if IsArray(T2) {
		if T1 == "iterable" {
			return true, ""
		}

		return false, fmt.Sprintf("cannot use %s as array of %s", T1, T2)
	}

	return true, ""
}

func compatibleCallable(T1, T2 string) (ok bool, desc string) {
	if T1 == "callable" {
		if T2 == "callable" || IsClosure(T2) || IsClass(T2) {
			return true, ""
		}

		return false, fmt.Sprintf("cannot use %s as %s", T1, T2)
	}
	if T2 == "callable" {
		if T1 == "callable" || IsClosure(T1) || IsClass(T1) {
			return true, ""
		}

		return false, fmt.Sprintf("cannot use %s as %s", T1, T2)
	}
	return true, ""
}

func compatibleBoolean(T1, T2 string) bool {
	if T1 == "bool" {
		return T2 == "bool" || T2 == "true" || T2 == "false"
	}
	if T2 == "bool" {
		return T1 == "bool" || T1 == "true" || T1 == "false"
	}
	return false
}
