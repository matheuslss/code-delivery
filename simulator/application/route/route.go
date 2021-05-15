package route

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type Route struct {
	ID        string
	ClientID  string
	Positions []Position
}

type Position struct {
	Lat  float64
	Long float64
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
