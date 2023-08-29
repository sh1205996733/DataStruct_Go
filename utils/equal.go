package utils

func Equal(v1, v2 any) bool {
	if v1 == nil {
		return v2 == nil
	}
	//return v1.equals(v2);
	return v1 == v2
}
