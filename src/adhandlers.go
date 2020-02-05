package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) handlepostadvertisement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Post ad has Been Called!")
		advertisement := PostAdvertisement{}
		err := json.NewDecoder(r.Body).Decode(&advertisement)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to add advertisement ")
			return
		}
		var advertisementposted bool
		var id, message string

		querystring := "SELECT * FROM public.addadvertisement('" + advertisement.UserID + "','" + advertisement.AdvertisementType + "','" + advertisement.EntityID + "','" + advertisement.Price + "','" + advertisement.Description + "')"
		err = s.dbAccess.QueryRow(querystring).Scan(&advertisementposted, &id, &message)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to add advertisement")
			fmt.Println(err.Error())
			fmt.Println("Error in communicating with database to add advertisement")
			return
		}

		postAdvertisementResult := PostAdvertisementResult{}
		postAdvertisementResult.AdvertisementPosted = advertisementposted
		postAdvertisementResult.ID = id
		postAdvertisementResult.Message = message

		js, jserr := json.Marshal(postAdvertisementResult)

		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to post advert")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handlegetadvertisement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Advert Has Been Called...")
		getadvertisementid := r.URL.Query().Get("id")
		advertisementid := AdvertisementID{}
		advertisementid.AdvertisementID = getadvertisementid

		var id, userid, advertisementtype, entityid, price, description string

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

		js, jserr := json.Marshal(advertisement)
		fmt.Println(js)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to get advertisement")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleupdateadvertisement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Update Advertisement Has Been Called...")
		advertisement := UpdateAdvertisement{}
		err := json.NewDecoder(r.Body).Decode(&advertisement)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to update advertisement")
			return
		}

		var advertisementupdated bool
		var msg string
		querystring := "SELECT * FROM public.updateadvertisement('" + advertisement.AdvertisementID + "','" + advertisement.UserID + "','" + advertisement.AdvertisementType + "','" + advertisement.EntityID + "','" + advertisement.Price + "','" + advertisement.Description + "')"
		err = s.dbAccess.QueryRow(querystring).Scan(&advertisementupdated, &msg)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to update advertisement")
			fmt.Println(err.Error())
			fmt.Println("Error in communicating with database to update advertisement")
			return
		}

		updateAdvertisementResult := UpdateAdvertisementResult{}
		updateAdvertisementResult.AdvertisementUpdated = advertisementupdated
		updateAdvertisementResult.Message = msg

		js, jserr := json.Marshal(updateAdvertisementResult)

		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to update advertisement")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleremoveadvertisement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Delete Advert Has Been Called..")
		getadvertisementid := r.URL.Query().Get("id")
		advertisementid := AdvertisementID{}
		advertisementid.AdvertisementID = getadvertisementid

		var advertisementDeleted bool
		querystring := "SELECT * FROM public.deleteadvertisement('" + advertisementid.AdvertisementID + "')"
		err := s.dbAccess.QueryRow(querystring).Scan(&advertisementDeleted)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to delete advertisement")
			fmt.Println(err.Error())
			fmt.Println("Error in communicating with database to delete advertisement")
			return
		}

		deleteAdvertisementResult := DeleteAdvertisementResult{}
		deleteAdvertisementResult.AdvertisementDeleted = advertisementDeleted
		deleteAdvertisementResult.AdvertisementID = getadvertisementid

		if advertisementDeleted {
			deleteAdvertisementResult.Message = "Advert Successfully Deleted!"
		} else {
			deleteAdvertisementResult.Message = "Unable to Delete Advert!"
		}

		js, jserr := json.Marshal(deleteAdvertisementResult)

		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to delete advert")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)

	}
}
