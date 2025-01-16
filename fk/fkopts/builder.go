package fkopts

import "reflect"

type Builder[T any] interface {
	List() []func(*T) error
}

func Build[T any](opts ...Builder[T]) (*T, error) {
	args := new(T)
	if base, ok := any(args).(interface{ Default() *T }); ok {
		base.Default()
	}
	for _, opt := range opts {
		if opt == nil || reflect.ValueOf(opt).IsNil() {
			// Do nothing if the option is nil or if opt is nil but implicitly cast as
			// an Options interface by the NewArgsFromOptions function. The latter
			// case would look something like this:
			continue
		}

		for _, setArgs := range opt.List() {
			if setArgs == nil {
				continue
			}

			if err := setArgs(args); err != nil {
				return nil, err
			}
		}
	}
	return args, nil
}
