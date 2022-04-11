package server

type RenderValue struct {
	StoreKeys []string
	Path      string
}

type RenderDto struct {
	Values []RenderValue
}
