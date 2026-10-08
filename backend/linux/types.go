package linux

type weight int

const (
	fc_weight_thin        weight = 0
	fc_weight_extralight  weight = 40
	fc_weight_light       weight = 50
	fc_weight_book        weight = 75
	fc_weight_regular     weight = 80
	fc_weight_medium      weight = 100
	fc_weight_demibold    weight = 180
	fc_weight_bold        weight = 200
	fc_weight_extrabold   weight = 205
	fc_weight_black       weight = 210
	fc_weight_extra_black weight = 215
	fc_weight_nord        weight = 205
)

type slant int

const (
	fc_slant_roman   slant = 0
	fc_slant_italic  slant = 100
	fc_slant_oblique slant = 110
)

type width int

const (
	fc_width_unknown        width = 0
	fc_width_ultracondensed width = 50
	fc_width_extracondensed width = 63
	fc_width_condensed      width = 75
	fc_width_semicondensed  width = 87
	fc_width_normal         width = 100
	fc_width_semiexpanded   width = 113
	fc_width_expanded       width = 125
	fc_width_extraexpanded  width = 150
	fc_width_ultraexpanded  width = 200
)
