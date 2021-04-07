package fizz

// SmartFizzer follows a set of steps to make possible the creation
// type SmartFizzer interface {
// 	Generate([]string) error
// 	Fizz() string
// 	UnFizz() string
// }

// func New(name string, args []string) (SmartFizzer, error) {
// 	var sf SmartFizzer

// 	switch {
// 	case strings.HasPrefix(name, "create_table"):
// 		sf = &createTable{name: name}
// 	case strings.HasPrefix(name, "drop_table"):
// 		sf = &dropTable{name: name}
// 	case strings.HasPrefix(name, "rename"):
// 		sf = &rename{name: name}
// 	default:
// 		sf = &createTable{name: name}
// 	}

// 	if err := sf.Generate(args); err != nil {
// 		return sf, err
// 	}

// 	return sf, nil
// }
