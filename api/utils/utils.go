package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/copier"
	"github.com/joho/godotenv"
)

// encoding pkg is to make tha json data form the data in the req
// io/ioutil pkg is to read data from req
// net/http pkg handles the req

// this fuction takes the r as request pointer and x as interface
// x is the Data Model stuct of the Specfic Database-Model(e.g... User,Book )
// x should be a pointer of the specific Model for which the req is made...
// the json data will be converted to the pointer passed as x
// in the end func will store the processed data to x
func PraseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			// fmt.Println(err)
			// fmt.Println("Error in data parsing")
			return
		}

	}
}

func JsonMarshal(x interface{}) ([]byte, error) {
	bodyRequest, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return bodyRequest, nil
}

// func to send response back to the client.....
func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v) // the .Encode converts the data to json dat type

}

func DataMapper(toValue interface{}, formValue interface{}, opt copier.Option) error {
	if err := copier.CopyWithOption(toValue, formValue, opt); err != nil {
		return err
	}
	return nil
}

// function to get load the api key data from the env
func EnvAMQPKEY() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
		log.Fatalln("Error While loading env")
	}
	return os.Getenv("RABBIT_AMQP_URI")
}
