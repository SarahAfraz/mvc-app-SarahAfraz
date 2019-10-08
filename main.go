package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

// The next three functions are "controllers" or "handlers":
// you can mostly use those terms interchangably. A controller
// typically extracts some information about the HTTP request
// (like the path information below: try going to "/hi/cats"),
// grab any data that are needed from the models, and stuff
// those data into the view, returning that to the user.
//
func loveHandler(w http.ResponseWriter, r *http.Request) {
	statementOfLove := ""
	lovedThings, foundLovedThings := r.URL.Query()["things"]
	if foundLovedThings {
		statementOfLove = strings.Join(lovedThings, ", ") + "!"
	} else {
		statementOfLove = "nothing. I am a rock. I am an island."
	}
	loveTemplate.Execute(w, fmt.Sprintf("Hi there, I love %s", statementOfLove))
}

func nicknameHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "happy-pelican")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexTemplate.Execute(w, nil)
}

func stringMatches(x string, y string) bool {
	x2 := strings.ToLower(x)
	y2 := strings.ToLower(y)
	return strings.Contains(x2, y2)
}

func attendeesHandler(w http.ResponseWriter, r *http.Request) {
	// The data that we collect to pass into the view is
	// often called the "context data". When you visit Instagram
	// to see your feed, they have one template (view) and just
	// populate it with different data for different people
	// We're doing the same thing.  The `party` struct is
	// defined in `models.go`. It has just one field called
	// `Attendees`, which is an array of strings.
	peopleAttendingParty := people
	queries, hasQuery := r.URL.Query()["q"]
	if hasQuery {
		peopleAttendingParty = make([]string, 0)
		for _, personName := range people {
			if stringMatches(personName, queries[0]) {
				peopleAttendingParty = append(peopleAttendingParty, personName)
			}
		}
	}

	contextData := party{Attendees: peopleAttendingParty}
	attendeesTemplate.Execute(w, contextData)
}

func getEnv(key, fallback string) string {
	value, foundValue := os.LookupEnv(key)
	if foundValue {
		return value
	}
	return fallback
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/love", loveHandler)
	http.HandleFunc("/nickname", nicknameHandler)
	http.HandleFunc("/attendees", attendeesHandler)
	http.ListenAndServe(":"+getEnv("PORT", "8080"), nil)
}
