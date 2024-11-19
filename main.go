package main

import (
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Risk struct {
	UUID        string `json:"uuid"`
	State       string `json:"state"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var AllowedStates = map[string]bool{
	"open":          true,
	"closed":        true,
	"accepted":      true,
	"investigating": true,
}

var risks = []Risk{}

// getRisks responds with the list of all risks as JSON.
func getRisks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, risks)
}

//Checking if state is one of [open, closed, accepted, investigating]
func NewState(state string, c *gin.Context) error {
	if _, valid := AllowedStates[state]; !valid {
		return errors.New("invalid state: must be one of [open, closed, accepted, investigating]")
	}
	return nil
}

func postRisks(c *gin.Context) {
	var newRisk Risk

	// Call BindJSON to bind the received JSON to newRisk.
	if err := c.BindJSON(&newRisk); err != nil {
		return
	}

	//State must be one of [open, closed, accepted, investigating]
	err := NewState(newRisk.State, c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "invalid state: must be one of [open, closed, accepted, investigating]"})
		return
	}
	newuuid := uuid.New()
	newRisk.UUID = newuuid.String()

	// Add the new risk to Risks.
	risks = append(risks, newRisk)
	c.IndentedJSON(http.StatusCreated, newRisk)
}
func getRiskByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of risks, looking for
	// an risk whose UUID value matches the parameter.
	for _, risk := range risks {
		if risk.UUID == id {
			c.IndentedJSON(http.StatusOK, risk)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Risk not found"})
}
func main() {
	r := gin.Default()

	r.GET("/v1/risks", getRisks)
	r.POST("/v1/risks", postRisks)
	r.GET("/v1/risks/:id", getRiskByID)

	// Run the server
	r.Run(":8080")
}


