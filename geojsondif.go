package geojsondif

import (
	"errors"
	"github.com/paulmach/go.geojson"
	"math"
)

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func RoundPoint(point []float64) []float64 {
	x, y := point[0], point[1]
	return []float64{Round(x, .5, 6), Round(y, .5, 6)}
}

func CheckPoint(point1, point2 []float64) error {
	point1, point2 = RoundPoint(point1), RoundPoint(point2)
	if !(point1[0] == point2[0] && point1[1] == point2[1]) {
		return errors.New("Points Don't Match")
	}
	return nil
}

func CheckLine(line1, line2 [][]float64) error {
	if len(line1) != len(line2) {
		return errors.New("Line Sizes Don't Match.")
	}
	for i := range line1 {
		err := CheckPoint(line1[i], line2[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func CheckLines(lines1, lines2 [][][]float64) error {
	if len(lines1) != len(lines2) {
		return errors.New("Number of Rings Don't Match.")
	}
	for i := range lines1 {
		err := CheckLine(lines1[i], lines2[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func CheckPolygons(lines1, lines2 [][][][]float64) error {
	if len(lines1) != len(lines2) {
		return errors.New("Number of Polygons Don't Match.")
	}
	for i := range lines1 {
		err := CheckLines(lines1[i], lines2[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func CheckGeom(geom1, geom2 *geojson.Geometry) error {
	if geom1.Type != geom2.Type {
		return errors.New("Different Geometry Types.")
	}
	switch geom1.Type {
	case "Point":
		return CheckPoint(geom1.Point, geom2.Point)
	case "MultiPoint":
		return CheckLine(geom1.MultiPoint, geom2.MultiPoint)
	case "LineString":
		return CheckLine(geom1.LineString, geom2.LineString)
	case "MultiLineString":
		return CheckLines(geom1.MultiLineString, geom2.MultiLineString)
	case "Polygon":
		return CheckLines(geom1.Polygon, geom2.Polygon)
	case "MultiPolygon":
		return CheckPolygons(geom1.MultiPolygon, geom2.MultiPolygon)
	}
	return nil
}

func CheckProperties(p1, p2 map[string]interface{}) error {
	if len(p1) != len(p2) {
		return errors.New("Different Property Sizes.")
	}
	for k := range p1 {
		val1, boolval1 := p1[k]
		val2, boolval2 := p2[k]
		if boolval1 && boolval2 {
			if val1 != val2 {
				errors.New("Val1 not equal to val2.")
			}
		} else {
			return errors.New("Property k isn't in both maps.")
		}
	}
	return nil
}

// Checks two desired identical features against one another
// Current precision for decimal points is currently 6
func CheckFeatures(feat1, feat2 *geojson.Feature) error {
	err := CheckGeom(feat1.Geometry, feat2.Geometry)
	if err != nil {
		return err
	}

	err = CheckProperties(feat1.Properties, feat2.Properties)
	if err != nil {
		return err
	}

	return nil
}
