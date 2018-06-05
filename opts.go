package opts

//Opts
type Opts struct {
	Threshold              float64  `json:"threshold"`
	MinDist                float64  `json:"mindist"`
	RelaxDist              float64  `json:"relaxdist"`
	PlanarSelf             bool     `json:"planarself"`
	NonPlanarSelf          bool     `json:"nonplanarself"`
	AvoidNewSelfIntersects bool     `json:"avoidself"`
	GeomRelation           bool     `json:"geomrelate"`
	DistRelation           bool     `json:"distrelate"`
	DirRelation            bool     `json:"dirrelate"`
}
