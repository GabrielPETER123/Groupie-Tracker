package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
)

type Artists struct {
    ID   int    `json:"id"`
	Image string `json:"image"`
    Name string `json:"name"`
	Members []string `json:"members"`
    CreationDate int `json:"creationDate"`
    FirstAlbum string `json:"firstAlbum"`
    Locations string `json:"locations"`
    ConcertDates string `json:"concertDates"`
    Relations string `json:"relations"`
}

type Dates struct {
	ID int `json:"id"`
	Dates []string `json:"dates"`
}
//** Structure de Dates */
type IndexDates struct {
	Index []Dates `json:"index"`
}

type Locations struct {
	ID int `json:"id"`
	Locations []string `json:"locations"`
	Dates string `json:"dates"`
}
//** Structure de Locations*/
type IndexLocations struct {
	Index []Locations `json:"index"`
}

type Relation struct {
	ID int `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`

}
//** Structure de Relation */
type IndexRelation struct {
	Index []Relation `json:"index"`
}

func main() {
	getArtists()
	getDates()
	getLocations()
	getRelation()
}

func getArtists() {
	//** GET l'API d'Artists*/
    respArt, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
    if err != nil {
        fmt.Println("Erreur lors de la requête:", err)
        os.Exit(1)
    }
    defer respArt.Body.Close()

    if respArt.StatusCode != http.StatusOK {
        fmt.Println("Erreur de réponse:", respArt.Status)
        os.Exit(1)
    }

    var artists []Artists

	//** Décodeur de la réponse JSON selon la structure Artist */
    if err := json.NewDecoder(respArt.Body).Decode(&artists); err != nil {
        fmt.Println("Erreur lors du décodage JSON:", err)
        os.Exit(1)
    }

	//** Supression du fichier artists.json */
    if err := os.Remove("artists.json"); err != nil {
        fmt.Println("Erreur lors de la suppression du fichier:", err)
    } else {
        fmt.Println("Le fichier dates.json a été supprimé")
    }

	//** Création du fichier artists.json */
	file, err := os.Create("artists.json")
    if err != nil {
        fmt.Println("Erreur lors de la création du fichier:", err)
        os.Exit(1)
    }
    defer file.Close()

	//** Écriture sur le fichier */
    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ") //Sert à mettre en forme pour une meilleure lisibilité
    if err := encoder.Encode(artists); err != nil {
        fmt.Println("Erreur lors de l'encodage JSON:", err)
        os.Exit(1)
    }

    fmt.Println("Les données des artistes ont été écrites dans artists.json")
}

func getDates() {
	//** GET l'API de Dates*/
    respDate, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
    if err != nil {
        fmt.Println("Erreur lors de la requête:", err)
        os.Exit(1)
    }
    defer respDate.Body.Close()

    if respDate.StatusCode != http.StatusOK {
        fmt.Println("Erreur de réponse:", respDate.Status)
        os.Exit(1)
    }

    var index IndexDates

	//** Décodeur de la réponse JSON selon la structure Dates */
    if err := json.NewDecoder(respDate.Body).Decode(&index); err != nil {
        fmt.Println("Erreur lors du décodage JSON:", err)
        os.Exit(1)
    }

	//** Supression du fichier dates.json */
    if err := os.Remove("dates.json"); err != nil {
        fmt.Println("Erreur lors de la suppression du fichier:", err)
    } else {
        fmt.Println("Le fichier dates.json a été supprimé")
    }

	//** Création du fichier dates.json */
	file, err := os.Create("dates.json")
    if err != nil {
        fmt.Println("Erreur lors de la création du fichier:", err)
        os.Exit(1)
    }
    defer file.Close()

	//** Écriture sur le fichier */
    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ") //Sert à mettre en forme pour une meilleure lisibilité
    if err := encoder.Encode(index); err != nil {
        fmt.Println("Erreur lors de l'encodage JSON:", err)
        os.Exit(1)
    }

    fmt.Println("Les données des artistes ont été écrites dans dates.json")
}

func getLocations() {
	//** GET l'API de Locations*/
	respLoc, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		fmt.Println("Erreur lors de la requête:", err)
		os.Exit(1)
	}
	defer respLoc.Body.Close()

	if respLoc.StatusCode != http.StatusOK {
		fmt.Println("Erreur de réponse:", respLoc.Status)
		os.Exit(1)
	}

	var index IndexLocations

	//** Décodeur de la réponse JSON selon la structure Locations */
	if err := json.NewDecoder(respLoc.Body).Decode(&index); err != nil {
		fmt.Println("Erreur lors du décodage JSON:", err)
		os.Exit(1)
	}

	//** Supression du fichier locations.json */
	if err := os.Remove("locations.json"); err != nil {
		fmt.Println("Erreur lors de la suppression du fichier:", err)
	} else {
		fmt.Println("Le fichier locations.json a été supprimé")
	}

	//** Création du fichier locations.json */
	file, err := os.Create("locations.json")
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier:", err)
		os.Exit(1)
	}
	defer file.Close()

	//** Écriture sur le fichier */
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") //Sert à mettre en forme pour une meilleure lisibilité
	if err := encoder.Encode(index); err != nil {
		fmt.Println("Erreur lors de l'encodage JSON:", err)
		os.Exit(1)
	}

	fmt.Println("Les données des artistes ont été écrites dans locations.json")
}

func getRelation() {
//** GET l'API de Relation*/
respRel, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
if err != nil {
	fmt.Println("Erreur lors de la requête:", err)
	os.Exit(1)
}
defer respRel.Body.Close()

if respRel.StatusCode != http.StatusOK {
	fmt.Println("Erreur de réponse:", respRel.Status)
	os.Exit(1)
}

var index IndexRelation

//** Décodeur de la réponse JSON selon la structure relation */
if err := json.NewDecoder(respRel.Body).Decode(&index); err != nil {
	fmt.Println("Erreur lors du décodage JSON:", err)
	os.Exit(1)
}

//** Supression du fichier relation.json */
if err := os.Remove("relation.json"); err != nil {
	fmt.Println("Erreur lors de la suppression du fichier:", err)
} else {
	fmt.Println("Le fichier relation.json a été supprimé")
}

//** Création du fichier relation.json */
file, err := os.Create("relation.json")
if err != nil {
	fmt.Println("Erreur lors de la création du fichier:", err)
	os.Exit(1)
}
defer file.Close()

//** Écriture sur le fichier */
encoder := json.NewEncoder(file)
encoder.SetIndent("", "  ") //Sert à mettre en forme pour une meilleure lisibilité
if err := encoder.Encode(index); err != nil {
	fmt.Println("Erreur lors de l'encodage JSON:", err)
	os.Exit(1)
}

fmt.Println("Les données des artistes ont été écrites dans relation.json")
}