package main

func Map[T any](data []T, action func(T) T) []T {
	if data == nil {
		return nil
	}

	if len(data) == 0 {
		return []T{}
	}

	res := make([]T, len(data))
	for i := range data {
		res[i] = action(data[i])
	}

	return res
}

func Filter[T any](data []T, action func(T) bool) []T {
	if data == nil {
		return nil
	}

	if len(data) == 0 {
		return []T{}
	}

	res := make([]T, 0)
	for i := range data {
		if action(data[i]) {
			res = append(res, data[i])
		}
	}

	return res
}

func Reduce[T any](data []T, initial T, action func(T, T) T) T {
	if data == nil {
		return initial
	}

	for i := range data {
		initial = action(data[i], initial)
	}

	return initial
}
