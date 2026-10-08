package fonts

type Weight uint8

const (
	WeightThin Weight = iota
	WeightExtraLight
	WeightLight
	WeightNormal
	WeightMedium
	WeightSemiBold
	WeightBold
	WeightExtraBold
	WeightBlack
)

type Slant uint8

const (
	SlantNormal Slant = iota
	SlantItalic
	SlantOblique
)
