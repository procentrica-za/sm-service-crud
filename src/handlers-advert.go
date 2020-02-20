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

		//set JSON object variables for response
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

func (s *Server) handleaddtextbook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		//get JSON payload
		textbook := Textbook{}
		err := json.NewDecoder(r.Body).Decode(&textbook)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to add textbook ")
			return
		}

		// set response variables.
		var textbookAdded bool
		var id string

		// write querystring
		querystring := "SELECT * FROM public.addtextbook('" + textbook.ModuleCode + "','" + textbook.Name + "','" + textbook.Edition + "','" + textbook.Quality + "','" + textbook.Author + "')"

		// query the database.
		err = s.dbAccess.QueryRow(querystring).Scan(&textbookAdded, &id)
		
		//check for response error of 500
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to add Textbook")
			fmt.Println(err.Error())
			fmt.Println("Error in communicating with database to add Textbook")
			return
		}

		// define response struct
		postTextbookResult := TextbookResult{}

		// populate response struct variables.
		if(textbookAdded){
			postTextbookResult.TextbookAdded = textbookAdded
			postTextbookResult.TextbookID = id
			postTextbookResult.Message = "Textbook Successfully Added!"
		}else{
			postTextbookResult.TextbookAdded = textbookAdded
			postTextbookResult.TextbookID = id
			postTextbookResult.Message = "Failed to add Textbook!"
		}

		//convert struct back to JSON
		js, jserr := json.Marshal(postTextbookResult)
		
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

func (s *Server) handleupdatetextbook() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
	//get JSON payload
	textbook := UpdateTextbook{}
	err := json.NewDecoder(r.Body).Decode(&textbook)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Bad JSON provided to update textbook")
		return
	}

	//set response variables
	var TextbookUpdated bool
	var msg string

	//communcate with the database
	querystring := "SELECT * FROM public.updatetextbook('" + textbook.TextbookID + "','" + textbook.ModuleCode + "','" + textbook.Name + "','" + textbook.Edition + "','" + textbook.Quality + "','" + textbook.Author + "')"
	err = s.dbAccess.QueryRow(querystring).Scan(&TextbookUpdated, &msg)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Unable to process DB Function to update textbook")
		fmt.Println("Error in communicating with database to update textbook")
		return
	}

	updateTextbookResult := UpdateTextbookResult{}
	updateTextbookResult.TextbookUpdated = TextbookUpdated
	updateTextbookResult.Message = msg

	//convert struct back to JSON
	js, jserr := json.Marshal(updateTextbookResult)

	//error occured when trying to convert struct to a JSON object
	if jserr != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Unable to create JSON object from DB result to update Textbook")
		return
	}

	//return back to advert service
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(js)
	}
}

func (s *Server) handlegettextbooksbyfilter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get JSON payload
		textbookfilter := TextbookFilter{}
		err := json.NewDecoder(r.Body).Decode(&textbookfilter)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to filter textbook ")
			return
		}
		//Build Query for Filtering by prepending and appending % to the filtering queries.
		querystring := "SELECT * FROM gettextbookbyfilter('%" + textbookfilter.ModuleCode + "%', '%" + textbookfilter.Name + "%' , '%" + textbookfilter.Edition+ "%' , '%" + textbookfilter.Quality + "%' , '%" + textbookfilter.Author + "%')"
		rows, err := s.dbAccess.Query(querystring)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		//define a list to hold all returned textbooks
		textbookList := TextbookList{}
		textbookList.Textbooks = []TextbookFilterResult{}

		var modulecode, id, name, edition, quality, author string

		//read the returned textbook into textbook struct this repeats for each textbook found.
		for rows.Next() {
			err = rows.Scan(&modulecode, &id, &name, &edition, &quality, &author)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from Textbook List...")
				fmt.Println(err.Error())
				return
			}
			//append the textbook struct to the textbook list.
			textbookList.Textbooks = append(textbookList.Textbooks, TextbookFilterResult{modulecode, id, name, edition, quality, author})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from Textbook List...")
			return
		}

		js, jserr := json.Marshal(textbookList)

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

func (s *Server) handleremovetextbook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//retrieve ID from advert service
		gettextbookid := r.URL.Query().Get("id")
		var textbookDeleted bool
		querystring := "SELECT * FROM public.deletetextbook('" + gettextbookid + "')"
		err := s.dbAccess.QueryRow(querystring).Scan(&textbookDeleted)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to delete textbook")
			fmt.Println("Error in communicating with database to delete textbook")
			return
		}

		//set response variables

		deleteTextbookResult := DeleteTextbookResult{}
		deleteTextbookResult.TextbookDeleted = textbookDeleted
		deleteTextbookResult.TextbookID = gettextbookid

		if textbookDeleted {
			deleteTextbookResult.Message = "textbook Successfully Deleted!"
		} else {
			deleteTextbookResult.Message = "Unable to Delete textbook!"
		}

		//convert struct back to JSON
		js, jserr := json.Marshal(deleteTextbookResult)

		//error occured when trying to convert struct to a JSON object
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to delete textbook")
			return
		}

		//return back to advert service
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)

	}
}

func (s *Server) handleaddnote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		//get JSON payload
		note := Note{}
		err := json.NewDecoder(r.Body).Decode(&note)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to add note ")
			return
		}

		// set response variables.
		var noteAdded bool
		var id string

		// write querystring
		querystring := "SELECT * FROM public.addnote('" + note.ModuleCode + "')"

		// query the database.
		err = s.dbAccess.QueryRow(querystring).Scan(&noteAdded, &id)
		
		//check for response error of 500
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to add Note")
			fmt.Println(err.Error())
			fmt.Println("Error in communicating with database to add Note")
			return
		}

		// define response struct
		postNoteResult := NoteResult{}

		// populate response struct variables.
		if(noteAdded){
			postNoteResult.NoteAdded = noteAdded
			postNoteResult.NoteID = id
			postNoteResult.Message = "Note Successfully Added!"
		}else{
			postNoteResult.NoteAdded = noteAdded
			postNoteResult.NoteID = id
			postNoteResult.Message = "Failed to add Note!"
		}

		//convert struct back to JSON
		js, jserr := json.Marshal(postNoteResult)
		
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

func (s *Server) handleupdatenote() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
	//get JSON payload
	note := UpdateNote{}
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Bad JSON provided to update note")
		return
	}

	//set response variables
	var NoteUpdated bool
	var msg string

	//communcate with the database
	querystring := "SELECT * FROM public.updatenote('" + note.NoteID + "', '" + note.ModuleCode + "')"
	err = s.dbAccess.QueryRow(querystring).Scan(&NoteUpdated, &msg)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Unable to process DB Function to update note")
		fmt.Println("Error in communicating with database to update note")
		return
	}

	updateNoteResult := UpdateNoteResult{}
	updateNoteResult.NoteUpdated = NoteUpdated
	updateNoteResult.Message = msg

	//convert struct back to JSON
	js, jserr := json.Marshal(updateNoteResult)

	//error occured when trying to convert struct to a JSON object
	if jserr != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Unable to create JSON object from DB result to update Note")
		return
	}

	//return back to advert service
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(js)
	}
}

func (s *Server) handlegetnotesbyfilter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get JSON payload
		notefilter := NoteFilter{}
		err := json.NewDecoder(r.Body).Decode(&notefilter)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to filter Note ")
			return
		}
		//Build Query for Filtering by prepending and appending % to the filtering queries.
		querystring := "SELECT * FROM getnotesbyfilter('%" + notefilter.ModuleCode + "%')"
		rows, err := s.dbAccess.Query(querystring)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		//define a list to hold all returned textbooks
		noteList := NoteList{}
		noteList.Notes = []NoteFilterResult{}

		var modulecode, id string

		//read the returned textbook into textbook struct this repeats for each textbook found.
		for rows.Next() {
			err = rows.Scan(&id, &modulecode)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from Note List...")
				fmt.Println(err.Error())
				return
			}
			//append the textbook struct to the textbook list.
			noteList.Notes = append(noteList.Notes, NoteFilterResult{modulecode, id})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from Note List...")
			return
		}

		js, jserr := json.Marshal(noteList)

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

func (s *Server) handleremovenote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//retrieve ID from advert service
		getnoteid := r.URL.Query().Get("id")
		var noteDeleted bool
		querystring := "SELECT * FROM public.deletenote('" + getnoteid + "')"
		err := s.dbAccess.QueryRow(querystring).Scan(&noteDeleted)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to delete note")
			fmt.Println("Error in communicating with database to delete note")
			return
		}

		//set response variables

		deleteNoteResult := DeleteNoteResult{}
		deleteNoteResult.NoteDeleted = noteDeleted
		deleteNoteResult.NoteID = getnoteid

		if noteDeleted {
			deleteNoteResult.Message = "note Successfully Deleted!"
		} else {
			deleteNoteResult.Message = "Unable to Delete note!"
		}

		//convert struct back to JSON
		js, jserr := json.Marshal(deleteNoteResult)

		//error occured when trying to convert struct to a JSON object
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to delete note")
			return
		}

		//return back to advert service
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)

	}
}


func (s *Server) handleaddtutor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		//get JSON payload
		tutor := Tutor{}
		err := json.NewDecoder(r.Body).Decode(&tutor)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to add Tutor ")
			return
		}

		// set response variables.
		var tutorAdded bool
		var id string

		// write querystring
		querystring := "SELECT * FROM public.addtutor('" + tutor.ModuleCode + "','" + tutor.Subject + "','" + tutor.YearCompleted + "','" + tutor.Venue + "','" + tutor.NotesIncluded + "', '" + tutor.Terms + "')"
	
		// query the database.
		err = s.dbAccess.QueryRow(querystring).Scan(&tutorAdded, &id)
		
		//check for response error of 500
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to add Tutor")
			fmt.Println(err.Error())
			fmt.Println("Error in communicating with database to add Tutor")
			return
		}

		// define response struct
		postTutorResult := TutorResult{}

		// populate response struct variables.
		if(tutorAdded){
			postTutorResult.TutorAdded = tutorAdded
			postTutorResult.TutorID = id
			postTutorResult.Message = "Tutor Successfully Added!"
		}else{
			postTutorResult.TutorAdded = tutorAdded
			postTutorResult.TutorID = id
			postTutorResult.Message = "Failed to add Tutor!"
		}

		//convert struct back to JSON
		js, jserr := json.Marshal(postTutorResult)
		
		//error occured when trying to convert struct to a JSON object
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to post Tutor")
			return
		}

		//return back to advert service
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleupdatetutor() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
	//get JSON payload
	tutor := UpdateTutor{}
	err := json.NewDecoder(r.Body).Decode(&tutor)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Bad JSON provided to update tutor")
		return
	}

	//set response variables
	var TutorUpdated bool
	var msg string

	//communcate with the database
	querystring := "SELECT * FROM public.updatetutor('" + tutor.TutorID + "', '" + tutor.ModuleCode + "', '" + tutor.Subject + "', '" + tutor.YearCompleted + "', '" + tutor.Venue + "', '" + tutor.NotesIncluded + "', '" + tutor.Terms + "')"
	err = s.dbAccess.QueryRow(querystring).Scan(&TutorUpdated, &msg)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Unable to process DB Function to update tutor")
		fmt.Println("Error in communicating with database to update tutor")
		return
	}

	updateTutorResult := UpdateTutorResult{}
	updateTutorResult.TutorUpdated = TutorUpdated
	updateTutorResult.Message = msg

	//convert struct back to JSON
	js, jserr := json.Marshal(updateTutorResult)

	//error occured when trying to convert struct to a JSON object
	if jserr != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Unable to create JSON object from DB result to update Tutor")
		return
	}

	//return back to advert service
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(js)
	}
}

func (s *Server) handlegettutorsbyfilter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get JSON payload
		tutorfilter := TutorFilter{}
		err := json.NewDecoder(r.Body).Decode(&tutorfilter)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to filter Tutor ")
			return
		}

		//Build Query for Filtering by prepending and appending % to the filtering queries.
		querystring := "SELECT * FROM gettutorbyfilter('%" + tutorfilter.ModuleCode + "%', '%" + tutorfilter.Subject + "%' , '%" + tutorfilter.YearCompleted + "%' , '%" + tutorfilter.Venue + "%' , '%" + tutorfilter.NotesIncluded + "%' , '%" + tutorfilter.Terms + "%')"
		rows, err := s.dbAccess.Query(querystring)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		//define a list to hold all returned tutors
		tutorList := TutorList{}
		tutorList.Tutors = []TutorFilterResult{}

		var modulecode, id, subject, yearcompleted, venue, notesincluded, terms  string

		//read the returned Tutor into Tutor struct this repeats for each Tutor found.
		for rows.Next() {
			err = rows.Scan(&id, &modulecode, &subject, &yearcompleted, &venue, &notesincluded, &terms)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from Tutor List...")
				fmt.Println(err.Error())
				return
			}
			//append the tutor struct to the tutor list.
			tutorList.Tutors = append(tutorList.Tutors, TutorFilterResult{modulecode, id, subject, yearcompleted, venue, notesincluded, terms})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from Tutor List...")
			return
		}

		js, jserr := json.Marshal(tutorList)

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

func (s *Server) handleremovetutor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//retrieve ID from advert service
		gettutorid := r.URL.Query().Get("id")
		var tutorDeleted bool
		querystring := "SELECT * FROM public.deletetutor('" + gettutorid + "')"
		err := s.dbAccess.QueryRow(querystring).Scan(&tutorDeleted)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to delete tutor")
			fmt.Println("Error in communicating with database to delete tutor")
			return
		}

		//set response variables

		deleteTutorResult := DeleteTutorResult{}
		deleteTutorResult.TutorDeleted = tutorDeleted
		deleteTutorResult.TutorID = gettutorid

		if tutorDeleted {
			deleteTutorResult.Message = "tutor Successfully Deleted!"
		} else {
			deleteTutorResult.Message = "Unable to Delete tutor!"
		}

		//convert struct back to JSON
		js, jserr := json.Marshal(deleteTutorResult)

		//error occured when trying to convert struct to a JSON object
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to delete tutor")
			return
		}

		//return back to advert service
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)

	}
}


func (s *Server) handleaddaccomodation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		//get JSON payload
		accomodation := Accomodation{}
		err := json.NewDecoder(r.Body).Decode(&accomodation)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to add Accomodation ")
			return
		}

		// set response variables.
		var accomodationAdded bool
		var id string

		// write querystring
		querystring := "SELECT * FROM public.addaccomodation('" + accomodation.AccomodationTypeCode + "','" + accomodation.InstitutionName + "','" + accomodation.Location + "','" + accomodation.DistanceToCampus + "')"

		// query the database.
		err = s.dbAccess.QueryRow(querystring).Scan(&accomodationAdded, &id)
		
		//check for response error of 500
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to add Accomodation")
			fmt.Println(err.Error())
			fmt.Println("Error in communicating with database to add Accomodation")
			return
		}

		// define response struct
		postAccomodationResult := AccomodationResult{}

		// populate response struct variables.
		if(accomodationAdded){
			postAccomodationResult.AccomodationAdded = accomodationAdded
			postAccomodationResult.AccomodationID = id
			postAccomodationResult.Message = "Accomodation Successfully Added!"
		}else{
			postAccomodationResult.AccomodationAdded = accomodationAdded
			postAccomodationResult.AccomodationID = id
			postAccomodationResult.Message = "Failed to add Accomodation!"
		}

		//convert struct back to JSON
		js, jserr := json.Marshal(postAccomodationResult)
		
		//error occured when trying to convert struct to a JSON object
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to post Accomodation")
			return
		}

		//return back to advert service
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleupdateaccomodation() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
	//get JSON payload
	accomodation := UpdateAccomodation{}
	err := json.NewDecoder(r.Body).Decode(&accomodation)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Bad JSON provided to update accomodation")
		return
	}

	//set response variables
	var AccomodationUpdated bool
	var msg string

	//communcate with the database
	querystring := "SELECT * FROM public.updateaccomodation('" + accomodation.AccomodationID + "', '" + accomodation.AccomodationTypeCode + "', '" + accomodation.InstitutionName + "', '" + accomodation.Location + "', '" + accomodation.DistanceToCampus + "')"
	err = s.dbAccess.QueryRow(querystring).Scan(&AccomodationUpdated, &msg)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Unable to process DB Function to update accomodation")
		fmt.Println("Error in communicating with database to update accomodation")
		return
	}

	updateAccomodationResult := UpdateAccomodationResult{}
	updateAccomodationResult.AccomodationUpdated = AccomodationUpdated
	updateAccomodationResult.Message = msg

	//convert struct back to JSON
	js, jserr := json.Marshal(updateAccomodationResult)

	//error occured when trying to convert struct to a JSON object
	if jserr != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Unable to create JSON object from DB result to update Accomodation")
		return
	}

	//return back to advert service
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(js)
	}
}


func (s *Server) handlegetaccomodationsbyfilter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get JSON payload
		accomodationfilter := AccomodationFilter{}
		err := json.NewDecoder(r.Body).Decode(&accomodationfilter)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to filter accomodation ")
			return
		}

		//Build Query for Filtering by prepending and appending % to the filtering queries.
		querystring := "SELECT * FROM getaccomodationbyfilter('%" + accomodationfilter.AccomodationTypeCode + "%', '%" + accomodationfilter.InstitutionName + "%' , '%" + accomodationfilter.Location + "%' , '%" + accomodationfilter.DistanceToCampus + "%')"
		rows, err := s.dbAccess.Query(querystring)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		//define a list to hold all returned accomodations
		accomodationList := AccomodationList{}
		accomodationList.Accomodations = []AccomodationFilterResult{}

		var id, accomodationtypecode, institutionname, location, distancetocampus string

		//read the returned accomodation into accomodation struct this repeats for each accomodation found.
		for rows.Next() {
			err = rows.Scan(&id, &institutionname, &accomodationtypecode, &location, &distancetocampus)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from accomodation List...")
				fmt.Println(err.Error())
				return
			}
			//append the accomodation struct to the accomodation list.
			accomodationList.Accomodations = append(accomodationList.Accomodations, AccomodationFilterResult{id, accomodationtypecode, institutionname, location, distancetocampus})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from Accomodation List...")
			return
		}

		js, jserr := json.Marshal(accomodationList)

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


func (s *Server) handleremoveaccomodation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//retrieve ID from advert service
		getaccomodationid := r.URL.Query().Get("id")
		var accomodationDeleted bool
		querystring := "SELECT * FROM public.deleteaccomodation('" + getaccomodationid + "')"
		err := s.dbAccess.QueryRow(querystring).Scan(&accomodationDeleted)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to delete accomodation")
			fmt.Println("Error in communicating with database to delete accomodation")
			return
		}

		//set response variables

		deleteAccomodationResult := DeleteAccomodationResult{}
		deleteAccomodationResult.AccomodationDeleted = accomodationDeleted
		deleteAccomodationResult.AccomodationID = getaccomodationid

		if accomodationDeleted {
			deleteAccomodationResult.Message = "accomodation Successfully Deleted!"
		} else {
			deleteAccomodationResult.Message = "Unable to Delete accomodation!"
		}

		//convert struct back to JSON
		js, jserr := json.Marshal(deleteAccomodationResult)

		//error occured when trying to convert struct to a JSON object
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to delete accomodation")
			return
		}

		//return back to advert service
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)

	}
}