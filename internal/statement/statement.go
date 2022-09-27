package statement

type IBlock interface {
	GetDirectives() []IDirective
	FindDirectives(directiveName string) []IDirective
}

type IDirective interface {
	GetName() string
	GetParameters() []string
	GetBlock() IBlock
}

type FileDirective interface {
	isFileDirective()
}

type IncludeDirective interface {
	FileDirective
}
