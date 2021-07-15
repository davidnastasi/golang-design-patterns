package example1

type ChristmasTreeDecorator interface {
	Do() string
}

type Decorator func(ChristmasTreeDecorator) ChristmasTreeDecorator

type ChristmasTreeDecoratorFunc func() string

func (cf ChristmasTreeDecoratorFunc) Do() string {
	return cf()
}

type ChristmasTree struct {
	ChristmasTreeDecorator
}

func Decorate(c ChristmasTreeDecorator, decorators ...Decorator) ChristmasTreeDecorator {
	decorated := c
	for _, decorator := range decorators {
		decorated = decorator(decorated)
	}
	return decorated
}


func (ct *ChristmasTree) Do() string {
	return "ChristmasTree:"
}

func  DecorateTreeTrooper() Decorator {
	return func(decorator ChristmasTreeDecorator) ChristmasTreeDecorator {
			return ChristmasTreeDecoratorFunc(func() string {
				return  decorator.Do() + " TreeTrooper"
			})
	}
}

func  DecorateTinsel() Decorator {
	return func(decorator ChristmasTreeDecorator) ChristmasTreeDecorator {
		return ChristmasTreeDecoratorFunc(func() string {
			return  decorator.Do() + " Tinsel"
		})
	}
}

func  DecorateGarland() Decorator {
	return func(decorator ChristmasTreeDecorator) ChristmasTreeDecorator {
		return ChristmasTreeDecoratorFunc(func() string {
			return  decorator.Do() + " Garland"
		})
	}
}


func  DecorateBubbleLights() Decorator {
	return func(decorator ChristmasTreeDecorator) ChristmasTreeDecorator {
		return ChristmasTreeDecoratorFunc(func() string {
			return  decorator.Do() + " BubbleLights"
		})
	}
}
