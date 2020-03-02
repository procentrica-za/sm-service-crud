package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// This function returns the filename and path for the advertisement card
func (s *Server) handlegetcardimagepath() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle get Card Image Path has been called...")

		entityid := r.URL.Query().Get("entityid")

		// Creating new struct for user login
		cardImage := CardImage{}
		cardImage.EntityID = entityid

		// declaring variables to catch response from database.
		var filepath, filename string

		// build query string.
		querystring := "SELECT * FROM public.getcardmainimage('" + cardImage.EntityID + "')"

		// querying the database and scanning database results into variables.
		err := s.dbAccess.QueryRow(querystring).Scan(&filepath, &filename)

		// checking for any errors with reading db response into variables.
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, err.Error())
			fmt.Println("Error in scanning the return variables from sql function into specified variables.")
			return
		}

		cardImage.FileName = filename
		cardImage.FilePath = filepath

		// converting response struct to JSON payload to send to service that called this function.
		js, jserr := json.Marshal(cardImage)

		// check to see if any errors occured with coverting to JSON.
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to fetch Card Image")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handlegetcardimagepathbatch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle get Card Image Path Batch has been called...")

		//get JSON payload
		imagerequest := CardImageBatchRequest{}
		err := json.NewDecoder(r.Body).Decode(&imagerequest)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to get batch image paths from CRUD")
			return
		}

		//communcate with the database
		querystring := "SELECT * FROM public.getcardmainimagelist("

		//The below foreach loop builds the query string with each element that needs to be included in the call
		for _, image := range imagerequest.Cards {
			querystring += "'" + image.EntityID + "',"
		}

		//Remove trailing "," from the querystring
		querystring = strings.TrimRight(querystring, ",")
		querystring = querystring + ")"

		// TODO: Remove below
		//THE BELOW IS FOR TESTING PURPOSES
		fmt.Println("Query String --> " + querystring)

		//retrieve rows from db
		rows, err := s.dbAccess.Query(querystring)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		cardBatchResult := CardImageBatch{}
		cardBatchResult.Images = []CardImage{}

		var entityid string
		var filepath string
		var filename string

		//read each row that has been returned and populate the batch object
		for rows.Next() {
			err = rows.Scan(&entityid, &filepath, &filename)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from Image Batch List")
				return
			}
			// TODO : remove below line
			fmt.Println("Image: " + entityid + " - " + filepath + " - " + filename)
			cardBatchResult.Images = append(cardBatchResult.Images, CardImage{entityid, filepath, filename})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from Student List...")
			return
		}

		//check for response error of 500
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to add advertisement")
			fmt.Println(err.Error())
			fmt.Println("Error in communicating with database to add advertisement")
			return
		}

		js, jserr := json.Marshal(cardBatchResult)

		//If Queryrow returns error, provide error to caller and exit
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON from DB result...")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}
