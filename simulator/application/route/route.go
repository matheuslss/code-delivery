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
	ID        string     `json:"id"`
	ClientID  string     `json:"clienteID"`
	Positions []Position `json:"positions"`
}

type Position struct {
	Lat  float64 `json:"latitude"`
	Long float64 `json:"longitude"`
}

type PartialRoutePosition struct {
	RouteID   string    `json:"routeID"`
	ClienteID string    `json:"clienteID"`
	Positions []float64 `json:"positions"`
	Finished  bool      `json:"finished"`
}

func NewRoute() *Route {
	return &Route{}
}

func (r *Route) LoadPositions() error {
	if r.ID == "" {
		return errors.New("Não foi possível obter o identificador da rota.")
	}

	file, err := os.Open("destinations/" + r.ID + ".txt")
	if err != nil {
		logrus.Error(err)
		return errors.New("Não foi possível obter as rotas.")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			logrus.Error(err)
			return errors.New("Não foi possível obter as rotas.")
		}
		long, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			logrus.Error(err)
			return errors.New("Não foi possível obter as rotas.")
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
		route.ClienteID = r.ClientID
		route.Positions = []float64{position.Lat, position.Long}
		route.Finished = false

		if len(r.Positions)-1 == i {
			route.Finished = true
		}

		jsonRoute, err := json.Marshal(route)
		if err != nil {
			logrus.Error(err)
			return []string{}, errors.New("Não foi possível exportar as rotas.")
		}

		result = append(result, string(jsonRoute))
	}
	return result, nil
}
