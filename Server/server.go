package Server

import (
	"challenge/Storage/database"
	"challenge/models"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"log"
	"net/http"
)

var gPlayers = make([]models.Player, 0)
// Router chi
func NewServer () http.Handler {
	r := chi.NewRouter()
	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	Cors := cors.New(cors.Options{

		AllowedOrigins:   []string{"http://localhost:8081"}, // Use this to allow specific origin hosts
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "Access-Control-Allow-Headers", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	r.Use(Cors.Handler)
	r.Route("/", func(r chi.Router) {
		r.Get("/players", getAllPlayers)                 //GET all players
		r.Post("/player", addPlayer)                     //POST player
		r.Put("/player/{id}", updatePlayer)          	//PUT player
	})
	return r
}
// fetch all players
func getAllPlayers(w http.ResponseWriter, r *http.Request) {
	db := database.NewRoachClient()
	rows, err := db.Query("SELECT * FROM players_tb ORDER BY score DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var player models.Player
	players := make([]models.Player, 0)
	for rows.Next() {
		if err := rows.Scan(&player.Id, &player.Name, &player.Score); err != nil {
			log.Fatal(err)
		}
		players = append(players, player)
		gPlayers = players
		fmt.Println(players)
	}
	_ = json.NewEncoder(w).Encode(players)
}
// post a player into de data base
func addPlayer(w http.ResponseWriter, r *http.Request) {
	request := map[string]string{}
	_ = json.NewDecoder(r.Body).Decode(&request)
	db := database.NewRoachClient()
	var player models.Player
	player.Name = request["name"]
	player.Score = request["score"]
	_, err := db.Exec("INSERT INTO players_tb (name, score) VALUES ('"+player.Name+"','"+player.Score+"')"); if err != nil {
		fmt.Println("cannot insert the data", err)
	}
}
// update the score of a player
func updatePlayer(w http.ResponseWriter, r *http.Request) {
	request := map[string]string{}
	_ = json.NewDecoder(r.Body).Decode(&request)
	db := database.NewRoachClient()
	for _, player:= range gPlayers{
		playerId := chi.URLParam(r,"id")
		if player.Id == playerId{
			player.Score = request["score"]
			_, err := db.Exec("UPDATE players_tb SET score = '"+player.Score+"' WHERE id = '"+player.Id+"'"); if err != nil {
				fmt.Println("cannot update the data", err)
			}
		}
	}
}
