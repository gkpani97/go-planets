package main

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Planet struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	DisatnceFromEarth int64  `json:"dist_from_earth"`
	Radius            int64  `json:"radius"`
	Mass              int64  `json:"mass"`
	IsTerrestrial     bool   `json:"is_terrestrial"`
}

type FuelEstimate struct {
	PlanetID     int `json:"id"`
	CrewCapacity int `json:"crew_cap"`
}

var planets = []Planet{
	{
		ID:                0,
		Name:              "Pluto",
		Description:       "Farthest Planet",
		DisatnceFromEarth: 299792458,
		Radius:            300,
		Mass:              300000,
		IsTerrestrial:     true,
	},
	{
		ID:                1,
		Name:              "Jupiter",
		Description:       "Largest Planet",
		DisatnceFromEarth: 29299792458,
		Radius:            300000,
		Mass:              300000000,
		IsTerrestrial:     true,
	},
}

func main() {
	r := gin.Default()

	r.GET("/planets", GetPlanets)
	r.GET("/planets/:id", GetPlanetByID)
	r.POST("/planet", AddPlanet)
	r.PUT("/planet/:id", UpdatePlanet)
	r.DELETE("/planet/:id", DeletePlanet)
	r.POST("/fuelestimate", EstimateFuelConsumption)

	r.Run()
}

func GetPlanets(c *gin.Context) {
	c.JSON(http.StatusOK, planets)
}

func GetPlanetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, planet := range planets {
		if planet.ID == id {
			c.JSON(http.StatusOK, planet)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Planet not found"})
}

func AddPlanet(c *gin.Context) {
	var planet Planet
	if err := c.BindJSON(&planet); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	planet.ID = planets[len(planets) - 1].ID + 1
	planets = append(planets, planet)
	c.JSON(http.StatusCreated, planet)
}

func UpdatePlanet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, planet := range planets {
		if planet.ID == id {
			var updatedPlanet Planet
			if err := c.BindJSON(&updatedPlanet); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			updatedPlanet.ID = id
			planets[i] = updatedPlanet
			c.JSON(http.StatusOK, updatedPlanet)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Planet not found"})
}

func DeletePlanet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, planet := range planets {
		if planet.ID == id {
			planets = append(planets[:i], planets[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Planet deleted successfully"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Planet not found"})
}

func EstimateFuelConsumption(c *gin.Context) {
	var target FuelEstimate
	if err := c.BindJSON(&target); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, planet := range planets {
		if planet.ID == target.PlanetID {
			targetPlanet := planet
			m, r, d, cc := float64(targetPlanet.Mass), float64(targetPlanet.Radius), float64(targetPlanet.DisatnceFromEarth), float64(target.CrewCapacity)
			var g, f float64
			if targetPlanet.IsTerrestrial {
				g = m / (r * r)
			} else {
				r := float64(targetPlanet.Radius)
				g = 0.5 / (r * r)
			}
			f = d / (g * g * cc)
			fuelCostEstimate := math.Round(f)
			c.JSON(http.StatusOK, gin.H{"Fuel Cost Estimation": fuelCostEstimate})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Target Planet not found"})
}
