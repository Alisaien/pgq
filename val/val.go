package val

func Int64(v *int64) int64 {
	if v == nil {
		return 0
	}

	return *v
}
