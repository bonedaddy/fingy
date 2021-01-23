package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Fingerprint holds information about a colelcted fingerpritn
type Fingerprint struct {
	gorm.Model
	IPAddress  string
	MurmurHash string
	UserAgent  string
	Language   string
	Platform   string
}

func main() {
	app := cli.NewApp()
	app.Name = "fingy"
	app.Usage = "fingy provides an api and database backend for collecting browser fingerprints using fingerprintjs2"
	app.Description = "fingy collects user agent, murmurFingerprint, language and platform information and is designed to be used with the companion index.html"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "listen.addr",
			Usage: "the address to listen on",
			Value: "0.0.0.0:6969",
		},
		&cli.StringFlag{
			Name:  "db.path",
			Usage: "path for the sqlite database",
			Value: "fingy.db",
		},
	}
	app.Action = func(c *cli.Context) error {
		db, err := gorm.Open(sqlite.Open(c.String("db.path")), &gorm.Config{})
		if err != nil {
			return err
		}
		defer func() {
			sdb, err := db.DB()
			if err != nil {
				return
			}
			sdb.Close()
		}()
		if err := db.AutoMigrate(&Fingerprint{}); err != nil {
			return err
		}
		r := chi.NewRouter()
		r.Use(cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
			AllowedHeaders:   []string{"*"},
			AllowCredentials: true,
		}).Handler)

		r.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
			data, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "failed to read body", http.StatusInternalServerError)
				return
			}
			// [["murmurFingerprint","7336c38f3c981f06e2c592d6ce4e7201"],["userAgent","Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0"],["language","en-CA"],["platform","Linux x86_64"]]
			var info [][2]string
			if err := json.Unmarshal(data, &info); err != nil {
				return
			}
			var entry Fingerprint
			for _, i := range info {
				switch i[0] {
				case "murmurFingerprint":
					entry.MurmurHash = i[1]
				case "userAgent":
					entry.UserAgent = i[1]
				case "language":
					entry.Language = i[1]
				case "platform":
					entry.Platform = i[1]
				}
			}
			entry.IPAddress = r.RemoteAddr
			log.Printf(
				"new fingerprint collected: hash %s, userAgent %s, language %s, platform %s, ip %s",
				entry.MurmurHash, entry.UserAgent, entry.Language, entry.Platform, entry.IPAddress,
			)
			if err := db.Create(&entry).Error; err != nil {
				log.Println("failed to store fingerprint in database: ", err)
			}
		})

		srv := &http.Server{
			Addr:    c.String("listen.addr"),
			Handler: r,
		}
		return srv.ListenAndServe()
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
