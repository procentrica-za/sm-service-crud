package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) handlepostadvertisement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get JSON payload
		advertisement := PostAdvertisement{}
		err := json.NewDecoder(r.Body).Decode(&advertisement)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to add advertisement ")
			return
		}
		//set response variables
		var advertisementposted bool
		var id, message string

		//communcate with the database
		querystring := "SELECT * FROM public.addadvertisement('" + advertisement.UserID + "','" + advertisement.AdvertisementType + "','" + advertisement.EntityID + "','" + advertisement.Price + "','" + advertisement.Description + "')"

		//retrieve message from database tt set to JSON object
		err = s.dbAccess.QueryRow(querystring).Scan(&advertisementposted, &id, &message)

		//check for response error of 500
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to add advertisement")
			fmt.Println(err.Error())
			fmt.Println("Error in communicating with database to add advertisement")
			return
		}

		//set JSON object variables for respinse
		postAdvertisementResult := PostAdvertisementResult{}
		postAdvertisementResult.AdvertisementPosted = advertisementposted
		postAdvertisementResult.ID = id
		postAdvertisementResult.Message = message

		//convert struct back to JSON
		js, jserr := json.Marshal(postAdvertisementResult)

		//error occured when trying to convert struct to a JSON object
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to post advert")
			return
		}

		//return back to advert service
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handlegetadvertisement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//retrieve URL from ad service
		getadvertisementid := r.URL.Query().Get("id")
		advertisementid := AdvertisementID{}
		advertisementid.AdvertisementID = getadvertisementid

		//set response variables
		var id, userid, advertisementtype, entityid, price, description string

		//communcate with the database
		querystring := "SELECT * FROM public.getadvertisement('" + advertisementid.AdvertisementID + "')"
		err := s.dbAccess.QueryRow(querystring).Scan(&id, &userid, &advertisementtype, &entityid, &price, &description)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to get advertisement")
			fmt.Println(err.Error())
			fmt.Println("Error in communicating with database to get advertisement")
			return
		}
		//fmt.Println("This is Advertisement!: " + id)
		advertisement := getAdvertisement{}
		advertisement.AdvertisementID = id
		advertisement.UserID = userid
		advertisement.AdvertisementType = advertisementtype
		advertisement.EntityID = entityid
		advertisement.Price = price
		advertisement.Description = description

		//convert struct back to JSON
		js, jserr := json.Marshal(advertisement)
		//error occured when trying to convert struct to a JSON object
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to get advertisement")
			return
		}

		//return back to advert service
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleupdateadvertisement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get JSON payload
		advertisement := UpdateAdvertisement{}
		err := json.NewDecoder(r.Body).Decode(&advertisement)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to update advertisement")
			return
		}

		//set response variables
		var advertisementupdated bool
		var msg string

		//communcate with the database
		querystring := "SELECT * FROM public.updateadvertisement('" + advertisement.AdvertisementID + "','" + advertisement.UserID + "','" + advertisement.AdvertisementType + "','" + advertisement.EntityID + "','" + advertisement.Price + "','" + advertisement.Description + "')"
		err = s.dbAccess.QueryRow(querystring).Scan(&advertisementupdated, &msg)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to update advertisement")
			fmt.Println("Error in communicating with database to update advertisement")
			return
		}

		updateAdvertisementResult := UpdateAdvertisementResult{}
		updateAdvertisementResult.AdvertisementUpdated = advertisementupdated
		updateAdvertisementResult.Message = msg

		//convert struct back to JSON
		js, jserr := json.Marshal(updateAdvertisementResult)

		//error occured when trying to convert struct to a JSON object
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to update advertisement")
			return
		}

		//return back to advert service
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleremoveadvertisement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//retrieve ID from advert service
		getadvertisementid := r.URL.Query().Get("id")

		advertisementid := AdvertisementID{}
		advertisementid.AdvertisementID = getadvertisementid

		var advertisementDeleted bool
		querystring := "SELECT * FROM public.deleteadvertisement('" + advertisementid.AdvertisementID + "')"
		err := s.dbAccess.QueryRow(querystring).Scan(&advertisementDeleted)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to delete advertisement")
			fmt.Println("Error in communicating with database to delete advertisement")
			return
		}

		//set response variables

		deleteAdvertisementResult := DeleteAdvertisementResult{}
		deleteAdvertisementResult.AdvertisementDeleted = advertisementDeleted
		deleteAdvertisementResult.AdvertisementID = getadvertisementid

		if advertisementDeleted {
			deleteAdvertisementResult.Message = "Advert Successfully Deleted!"
		} else {
			deleteAdvertisementResult.Message = "Unable to Delete Advert!"
		}

		//convert struct back to JSON
		js, jserr := json.Marshal(deleteAdvertisementResult)

		//error occured when trying to convert struct to a JSON object
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to delete advert")
			return
		}

		//return back to advert service
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)

	}
}

func (s *Server) handlegetuserdetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//retrieve URL from ad service
		getadvertisementid := r.URL.Query().Get("id")
		advertisementid := AdvertisementID{}
		advertisementid.AdvertisementID = getadvertisementid

		//set response variables
		var id, userid, advertisementtype, entityid, price, description string

		//communcate with the database
		querystring := "SELECT * FROM public.getadvertisement('" + advertisementid.AdvertisementID + "')"
		err := s.dbAccess.QueryRow(querystring).Scan(&id, &userid, &advertisementtype, &entityid, &price, &description)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to get advertisement")
			fmt.Println(err.Error())
			fmt.Println("Error in communicating with database to get advertisement")
			return
		}
		//fmt.Println("This is Advertisement!: " + id)
		advertisement := getAdvertisement{}
		advertisement.AdvertisementID = id
		advertisement.UserID = userid
		advertisement.AdvertisementType = advertisementtype
		advertisement.EntityID = entityid
		advertisement.Price = price
		advertisement.Description = description

		//convert struct back to JSON
		js, jserr := json.Marshal(advertisement)
		fmt.Println(js)
		//error occured when trying to convert struct to a JSON object
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to get advertisement")
			return
		}

		//return back to advert service
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handledeleteuseradvertisements() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//retrieve ID from advert service
		getuserid := r.URL.Query().Get("id")

		userid := UserID{}
		userid.UserID = getuserid

		var advertisementsDeleted bool
		querystring := "SELECT * FROM public.deleteuseradvertisements('" + userid.UserID + "')"
		err := s.dbAccess.QueryRow(querystring).Scan(&advertisementsDeleted)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to delete advertisements")
			fmt.Println(err.Error())
			fmt.Println("Error in communicating with database to delete advertisements")
			return
		}

		//set response variables

		deleteAdvertisementsResult := DeleteAdvertisementsResult{}
		deleteAdvertisementsResult.AdvertisementsDeleted = advertisementsDeleted

		if advertisementsDeleted {
			deleteAdvertisementsResult.Message = "Adverts Successfully Deleted!"
		} else {
			deleteAdvertisementsResult.Message = "Unable to Delete Advert!"
		}

		//convert struct back to JSON
		js, jserr := json.Marshal(deleteAdvertisementsResult)

		//error occured when trying to convert struct to a JSON object
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to delete adverts")
			return
		}

		//return back to advert service
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)

	}
}








func (s *Server) handlegetuseradvertisements() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userid := r.URL.Query().Get("id")

		rows, err := s.dbAccess.Query("SELECT * FROM getadvertisementbyuserid('" + userid + "')")
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		userAdvertList := UserAdvertisementList{}
		userAdvertList.UserAdvertisements = []GetUserAdvertisementResult{}

		var advertid string
		var advertisementtype string
		var entityid string
		var price string
		var description string

		for rows.Next() {
			err = rows.Scan(&advertid, &advertisementtype, &entityid, &price, &description)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from User Advertisement List...")
				fmt.Println(err.Error())
				return
			}
			userAdvertList.UserAdvertisements = append(userAdvertList.UserAdvertisements, GetUserAdvertisementResult{advertid, advertisementtype, entityid, price, description})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from Advertisement List...")
			return
		}

		js, jserr := json.Marshal(userAdvertList)

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







func (s *Server) handlegetalladvertisements() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := s.dbAccess.Query("SELECT * FROM public.getalladvertisements()")
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		advertList := AdvertisementList{}
		advertList.Advertisements = []getAdvertisement{}

		var advertid string
		var userid string
		var advertisementtype string
		var entityid string
		var price string
		var description string

		for rows.Next() {
			err = rows.Scan(&advertid, &userid, &advertisementtype, &entityid, &price, &description)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from Advertisement List...")
				fmt.Println(err.Error())
				return
			}
			advertList.Advertisements = append(advertList.Advertisements, getAdvertisement{advertid, userid, advertisementtype, entityid, price, description})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from Advertisement List...")
			return
		}

		js, jserr := json.Marshal(advertList)

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







func (s *Server) handlegetadvertisementbytype() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		advertisementType := r.URL.Query().Get("adverttype")

		rows, err := s.dbAccess.Query("SELECT * FROM getadvertisementbytype('" + advertisementType + "')")
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		advertList := TypeAdvertisementList{}
		advertList.TypeAdvertisements = []getAdvertisement{}

		var advertid string
		var userid string
		var advertisementtype string
		var entityid string
		var price string
		var description string

		for rows.Next() {
			err = rows.Scan(&advertid, &userid, &advertisementtype, &entityid, &price, &description)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from Advertisement List...")
				fmt.Println(err.Error())
				return
			}
			advertList.TypeAdvertisements = append(advertList.TypeAdvertisements, getAdvertisement{advertid, userid, advertisementtype, entityid, price, description})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from Advertisement List...")
			return
		}

		js, jserr := json.Marshal(advertList)

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
