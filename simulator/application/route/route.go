package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type Route struct {
	ID        string     `json:"routeID"`
	ClientID  string     `json:"clientID"`
	Positions []Position `json:"positions"`
}

type Position struct {
	Lat  float64 `json:"latitude"`
	Long float64 `json:"longitude"`
}

type PartialRoutePosition struct {
	RouteID   string    `json:"routeID"`
	ClientID  string    `json:"clientID"`
	Positions []float64 `json:"positions"`
	Finished  bool      `json:"finished"`
}

func NewRoute() *Route {
	return &Route{}
}

func (r *Route) LoadPositions() error {
	if r.ID == "" {
		logrus.Errorln(errors.New("error: invalid routeID"))
		return errors.New("error: invalid routeID")
	}

	file, err := os.Open("destinations/" + r.ID + ".txt")
	if err != nil {
		logrus.Error(err)
		return errors.New("error: the file could not be opened")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			logrus.Error(err)
			return errors.New("error: could convert latitude")
		}
		long, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			logrus.Error(err)
			return errors.New("error: could not convert longitude")
		}
		r.Positions = append(r.Positions, Position{
			Lat:  lat,
			Long: long,
		})
	}

	return nil
}

func (r *Route) ExportJsonPositions() ([]string, error) {
	var route PartialRoutePosition
	var result []string

	for i, position := range r.Positions {
		route.RouteID = r.ID
		route.ClientID = r.ClientID
		route.Positions = []float64{position.Lat, position.Long}
		route.Finished = false

		if len(r.Positions)-1 == i {
			route.Finished = true
		}

		jsonRoute, err := json.Marshal(route)
		if err != nil {
			logrus.Error(err)
			return []string{}, errors.New("error: unable to export route")
		}

		result = append(result, string(jsonRoute))
	}
	return result, nil
}
