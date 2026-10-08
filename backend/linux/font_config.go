package linux

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"swing-go/fonts"
)

type fontConfigResolver struct{}

func NewFontConfigResolver() fontConfigResolver {
	return fontConfigResolver{}
}

func (fontConfigResolver) Match(query fonts.Query) (fonts.Match, error) {
	search := fmt.Sprintf(
		"%s:slant=%d:weight=%d",
		query.Family,
		parseSlant(query.Slant),
		parseWeight(query.Weight),
	)

	cmd := exec.Command("fc-match", search, "-f", "%{file}|%{index}")

	data, err := cmd.Output()
	if err != nil {
		return fonts.Match{}, err
	}

	metadata := strings.SplitN(string(data), "|", 2)

	if len(metadata) != 2 {
		return fonts.Match{}, fmt.Errorf("font found not has expected format")
	}

	num, err := strconv.Atoi(metadata[1])
	if err != nil {
		return fonts.Match{}, err
	}

	return fonts.Match{
		File:  metadata[0],
		Index: num,
	}, nil
}

func parseSlant(slant fonts.Slant) slant {
	switch slant {
	case fonts.SlantNormal:
		return fc_slant_roman

	case fonts.SlantItalic:
		return fc_slant_italic

	case fonts.SlantOblique:
		return fc_slant_oblique

	default:
		return fc_slant_roman
	}
}

func parseWeight(weight fonts.Weight) weight {
	switch weight {
	case fonts.WeightThin:
		return fc_weight_thin
	case fonts.WeightExtraLight:
		return fc_weight_extralight
	case fonts.WeightLight:
		return fc_weight_light
	case fonts.WeightNormal:
		return fc_weight_regular
	case fonts.WeightMedium:
		return fc_weight_medium
	case fonts.WeightSemiBold:
		return fc_weight_demibold
	case fonts.WeightBold:
		return fc_weight_bold
	case fonts.WeightExtraBold:
		return fc_weight_extrabold
	case fonts.WeightBlack:
		return fc_weight_black
	default:
		return fc_weight_regular
	}
}
