package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/questions", getAllQuestions)    // all questions
	r.GET("/question", getRandomQuestion)   // random question
	r.GET("/question/:id", getQuestionByID) // specific Q
	r.Run("0.0.0.0:8080")
}

type question struct {
	ID       string `json:"id"`
	Question string `json:"question"`
}

var questionsMap = map[string]question{
	"1152881b-221d-437d-8e8d-aae7bbb24abc": {"1152881b-221d-437d-8e8d-aae7bbb24abc", "What is your favorite color and why?"},
	"15240739-1ad9-43ea-bf5a-66dce200dced": {"15240739-1ad9-43ea-bf5a-66dce200dced", "What is your dream vacation like?"},
	"266ab3b0-cedf-4450-b11f-2ae81107636d": {"266ab3b0-cedf-4450-b11f-2ae81107636d", "Who is your biggest role model and why?"},
	"b0e5e096-76cc-475f-b523-8820b4aa5c00": {"b0e5e096-76cc-475f-b523-8820b4aa5c00", "Where did you see yourself in 10 years when you were 16?"},
	"b5266b38-1992-46a0-9bc9-defcd20ee7d5": {"b5266b38-1992-46a0-9bc9-defcd20ee7d5", "Why did you choose a career in tech?"},
}

// returns random key from map using slice
func randomizer(m map[string]question) string {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	randomNum := rand.Intn(len(m))
	return keys[randomNum]
}

func getRandomQuestion(c *gin.Context) {
	key := randomizer(questionsMap)
	question := questionsMap[key]
	c.JSON(http.StatusOK, question)
}

func getQuestionByID(c *gin.Context) {
	id := c.Param("id")
	question := questionsMap[id]
	c.JSON(http.StatusOK, question)
}

func getAllQuestions(c *gin.Context) {
	var questions []question
	for _, v := range questionsMap {
		questions = append(questions, v)
	}
	c.JSON(http.StatusOK, questions)
}
