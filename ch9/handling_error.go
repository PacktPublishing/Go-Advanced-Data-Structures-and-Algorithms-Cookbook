func MyFunc1(v interface{}) (interface{}, error) {
	var ok bool

	if !ok {
		return nil, errors.New("not ok error")
	}
	return v, nil
}

func MyFunc2() {
	defer func() {
		if err := recover(); err != nil {
			// recover from panic
			fmt.Println("recover: ", err)
		}
	}()
	v := struct{}{}
	if _, err := MyFunc1(v); err != nil {
		panic(err) // panic
	}

	fmt.Println("never happen")
}

func main() {
	MyFunc2()
	fmt.Println("main finish")
}
