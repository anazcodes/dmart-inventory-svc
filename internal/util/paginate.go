package util

func Paginate(page, count int64) (skip, limit int64) {
	if page <= 0 {
		page = 1
	}
	if count < 10 {
		count = 6
	}

	skip = (page - 1) * count
	limit = skip + count
	return
}
