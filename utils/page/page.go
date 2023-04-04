package utils

// TODO move ohdat/app/utils to here
func Page(page *int) {
	if page == nil || *page < 1 {
		*page = 1
	}
}

func Limit(limit *int) {
	if limit == nil || *limit < 1 {
		*limit = 10
	}
}
