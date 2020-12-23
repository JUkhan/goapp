package util

type (
	T    interface{}
	List struct {
		Data []T
	}
)

func New(data []T) *List {
	return &List{data}
}

func (l *List) Map(fx func(T, int) T) *List {
	res := []T{}
	for index, item := range l.Data {
		res = append(res, fx(item, index))
	}
	l.Data = res
	return l
}

func (l *List) Filter(fx func(T, int) bool) *List {
	res := []T{}
	for index, item := range l.Data {
		if fx(item, index) {
			res = append(res, item)
		}
	}
	l.Data = res
	return l
}

func (l *List) Foreach(fx func(T)) *List {
	for _, item := range l.Data {
		fx(item)
	}
	return l
}

func (l *List) Reduce(init T, fx func(T, T, int) T) T {
	for index, item := range l.Data {
		init = fx(init, item, index)
	}
	return init
}

func (l *List) Scan(init T, fx func(T, T, int) T) *List {
	res := []T{}
	count := 0
	for index, item := range l.Data {
		count += fx(init, item, index).(int)
		res = append(res, count)
	}
	l.Data = res
	return l
}
