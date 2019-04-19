package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/context"
	"github.com/go-chi/chi"
	"github.com/dgrijalva/jwt-go"
	. "github.com/user/go-rest-api-mongodb/config"
	. "github.com/user/go-rest-api-mongodb/crypto"
	. "github.com/user/go-rest-api-mongodb/dao"
	. "github.com/user/go-rest-api-mongodb/helper"
	. "github.com/user/go-rest-api-mongodb/models"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type Controller struct{}

type JwtToken struct {
    Token string `json:"token"`
    Success bool `json:"success"`
    Name string  `json:"name"`
}

var config = Config{}
var dao = Dao{}
var helper = Helper{}
var crypto = Hash{}

func (c *Controller) GetAllMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println(context.Get(r, "token"))
	authenticated := context.Get(r, "token")
	if authenticated != false {
		movies, err := dao.FindAll()
		if err != nil {
			helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		helper.RespondWithJson(w, http.StatusOK, movies)
	} else {
		helper.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
	}

}

func (c *Controller) GetMovieById(w http.ResponseWriter, r *http.Request) {

	authenticated := context.Get(r, "token")
	if authenticated != false {
		params := chi.URLParam(r, "id")
		movie, err := dao.FindById(params)
		if err != nil {
			helper.RespondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
			return
		}
		helper.RespondWithJson(w, http.StatusOK, movie)
	} else {
		helper.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
	}

}

func (c *Controller) AddMovie(w http.ResponseWriter, r *http.Request) {
	authenticated := context.Get(r, "token")
	if authenticated != false {
		defer r.Body.Close()
		var movie Movie
		if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
			helper.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		movie.ID = bson.NewObjectId()
		if err := dao.Insert(movie); err != nil {
			helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		helper.RespondWithJson(w, http.StatusCreated, movie)
	} else {
		helper.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
	}

}

func (c *Controller) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	authenticated := context.Get(r, "token")
	if authenticated != false {
		defer r.Body.Close()
		var movie Movie
		if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
			helper.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		if err := dao.Update(movie); err != nil {
			helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		helper.RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
	} else {
		helper.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
	}
}

func (c *Controller) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	authenticated := context.Get(r, "token")
	if authenticated != false {
		params := chi.URLParam(r, "id")
		movie, err := dao.FindById(params)
		if err != nil {
			helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		errFromDelete := dao.Delete(movie)
		if errFromDelete != nil {
			helper.RespondWithError(w, http.StatusInternalServerError, errFromDelete.Error())
			return
		}
		helper.RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
	} else {
		helper.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
	}
}

func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user.ID = bson.NewObjectId()
	user.Password, _ = crypto.Generate(user.Password)
	if err := dao.CreateUser(user); err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	user.Password = ""
	helper.RespondWithJson(w, http.StatusCreated, user)
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	fmt.Println(user)

	// find user by name
	newUser, err := dao.FindUserByName(user.Name)
	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// compare password
	error := crypto.Compare(newUser.Password, user.Password)
	
	if error != nil{
		helper.RespondWithError(w, http.StatusInternalServerError, "Incorrect Password")
		return
	}
	newUser.Password = ""

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "name": newUser.Name,
        "profile_image": newUser.ProfileImage,
        "phone_number": newUser.PhoneNumber,
    })
    
    tokenString, error := token.SignedString([]byte("secret"))
	helper.RespondWithJson(w, http.StatusOK, JwtToken{Token: tokenString, Success: true, Name: newUser.Name })
}

// func (c *Controller) InsertMessage(message string) {
// 	if err := json.NewDecoder(message).Decode(&message); err != nil {
// 		helper.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}
// 	message.ID = bson.NewObjectId()
// 	if err := dao.InsertMessage(message); err != nil {
// 		// helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	return
// }


func (c *Controller) Init() {
	config.Read()
	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}
