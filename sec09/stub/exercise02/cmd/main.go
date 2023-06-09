package main

import (
	"fmt"

	"github.com/egbordzor/learn-glft/sec09/stub/exercise02/acura"
	"github.com/egbordzor/learn-glft/sec09/stub/exercise02/ford"
	"github.com/egbordzor/learn-glft/sec09/stub/exercise02/honda"
	"github.com/egbordzor/learn-glft/sec09/stub/exercise02/tesla"
	"github.com/egbordzor/learn-glft/sec09/stub/exercise02/vehicle"
)

func main() {
	myGarage := make([]vehicle.Vehicle, 4)

	myGarage[0], _ = ford.NewExplorer(1900)
	myGarage[1], _ = honda.NewPilot(2010)
	myGarage[2], _ = acura.NewTSX(2006)
	myGarage[3], _ = tesla.NewModel3(2018)

	for _, v := range myGarage {
		fmt.Println(v)
	}
}
