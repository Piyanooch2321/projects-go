package main

import (
	"hello/fizzbuzz"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// import "fmt"

func main() {
	router := gin.Default()

	router.POST("/credentials", CredintialHendler)

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.GET("/test/:number", func(c *gin.Context) {
		number := c.Param("number")
		n, err := strconv.Atoi(number) // shortcut for c.Request.URL.Query().Get("lastname")

		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		c.String(http.StatusOK, fizzbuzz.String(n))
	})

	router.Run(":8081")

}

type credential struct {
	EmailAddress string
	Password     string
}

func CredintialHendler(c *gin.Context) {
	var cred credential
	err := c.Bind(&cred)
	if err != nil {

	}

	mySigningKey := []byte("P@ssw0rd")

	type MyCustomClaims struct {
		EmailAddress string `json:"email"`
		Password     string `json:"password"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := MyCustomClaims{
		cred.EmailAddress,
		cred.Password,
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {

	}

	c.JSON(http.StatusOK, map[string]string{
		"token": ss,
	})
	// fmt.Printf("%v %v", ss, err)

}
