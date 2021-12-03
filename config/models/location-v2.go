package models

type LocationV2 struct {
	// Root path (all relative path will resolve base on this config)
	Paths []string
}

func DefaultLocationV2(paths []string) *LocationV2 {
	return &LocationV2{
		Paths: paths,
	}
}
