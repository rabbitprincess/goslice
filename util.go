package goslice

// prevent index out of range ( panic )
func preprocIndexException(i, j *int, len int) {
	if *i < 0 {
		*i = 0
	} else if *i > len {
		*i = len
	}

	if j != nil {
		if *j > len {
			*j = len
		}
		if *j < 0 {
			*j = 0
		}
		if *i > *j {
			*i, *j = *j, *i
		}
	}
}
