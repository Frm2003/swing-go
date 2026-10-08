package fonts

type resolver interface {
	Match(query Query) (Match, error)
}
