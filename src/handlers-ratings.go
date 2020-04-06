package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (s *Server) handleratebuyer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handleaddchat Has Been Called!")
		//get JSON payload
		rating := StartRating{}
		err := json.NewDecoder(r.Body).Decode(&rating)
		fmt.Println("Handle rate buyer Has Been Called..")
		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to rate buyer ")
			return
		}
		//set response variables
		var buyerrated bool
		var ratingid string

		//communcate with the database
		querystring := "SELECT * FROM public.ratebuyer('" + rating.AdvertisementID + "','" + rating.BuyerID + "','" + rating.SellerID + "','" + rating.BuyerRating + "','" + rating.BuyerComments + "')"

		//retrieve message from database tt set to JSON object
		err = s.dbAccess.QueryRow(querystring).Scan(&buyerrated, &ratingid)

		//check for response error of 500
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to rate buyer")
			fmt.Println(err.Error())
			fmt.Println("Error in communicating with database to rate buyer")
			return
		}

		//set JSON object variables for response
		startratingResult := StartRatingResult{}
		startratingResult.BuyerRated = buyerrated
		startratingResult.RatingID = ratingid

		if buyerrated {
			startratingResult.Message = "Buyer sucessfully rated!"
		} else {
			startratingResult.Message = "Buyer has already been rated!"
		}

		//convert struct back to JSON
		js, jserr := json.Marshal(startratingResult)

		//error occured when trying to convert struct to a JSON object
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to rate buyer")
			return
		}

		//return back to advert service
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handlerateseller() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Rate seller Has Been Called...")
		// declare a updateUser struct.
		sellerrrating := RateSeller{}
		// convert received JSON payload into the declared struct.
		err := json.NewDecoder(r.Body).Decode(&sellerrrating)
		//check for errors when converting JSON payload into struct.
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to rate seller")
			return
		}
		// declare variables to catch response from database.
		var sellerRated bool
		// building query string.
		querystring := "SELECT * FROM public.rateseller('" + sellerrrating.RatingID + "','" + sellerrrating.SellerRating + "','" + sellerrrating.SellerComments + "')"
		// query the database and read results into variables.
		err = s.dbAccess.QueryRow(querystring).Scan(&sellerRated)
		// check for errors with reading database result into variables.
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, err.Error())
			fmt.Println("Error in communicating with database to rate seller")
			return
		}
		// instansiate result struct.
		rateSellerResult := RateSellerResult{}
		rateSellerResult.SellerRated = sellerRated

		if sellerRated {
			rateSellerResult.Message = "Seller sucessfully rated!"
		} else {
			rateSellerResult.Message = "Seller has not been rated!"
		}
		// convert struct into JSON payload to send to service that called this fuction.
		js, jserr := json.Marshal(rateSellerResult)
		// check for errors in converting struct to JSON.
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to rate seller")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handlegetoutstandingratings() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle getoustanndingratings Has Been Called..")
		userid := r.URL.Query().Get("userid")

		rows, err := s.dbAccess.Query("SELECT * FROM public.getoutstandingratings('" + userid + "')")
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to get oustanding ratings")
			return
		}
		defer rows.Close()

		outstandingRatingList := OutstandingRatingList{}
		outstandingRatingList.Oustandingratings = []GetOutstandingResult{}

		var id string
		var username string
		var price string
		var title string
		var description string

		for rows.Next() {
			err = rows.Scan(&id, &username, &price, &title, &description)

			if title == "" {
				outstandingRatingList.Oustandingratings = append(outstandingRatingList.Oustandingratings, GetOutstandingResult{id, username, price, "Advertisement", description})
			} else {
				outstandingRatingList.Oustandingratings = append(outstandingRatingList.Oustandingratings, GetOutstandingResult{id, username, price, title, description})
			}

		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from Outstanding ratings list...")
			return
		}

		js, jserr := json.Marshal(outstandingRatingList)

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

func (s *Server) handlegetsellerratings() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle getsellerratings Has Been Called..")
		userid := r.URL.Query().Get("userid")

		rows, err := s.dbAccess.Query("SELECT * FROM public.sellerratings('" + userid + "')")
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to get oustanding ratings")
			return
		}
		defer rows.Close()

		previousRatingList := PreviousRatingList{}
		previousRatingList.Previousratings = []GetPreviousResult{}

		var id string
		var username string
		var rating string
		var comment string

		for rows.Next() {
			err = rows.Scan(&id, &username, &rating, &comment)

			previousRatingList.Previousratings = append(previousRatingList.Previousratings, GetPreviousResult{id, username, rating, comment})

		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from Previous ratings list...")
			return
		}

		js, jserr := json.Marshal(previousRatingList)

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

func (s *Server) handlegetbuyerratings() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle getsellerratings Has Been Called..")
		userid := r.URL.Query().Get("userid")

		rows, err := s.dbAccess.Query("SELECT * FROM public.buyerratings('" + userid + "')")
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to get oustanding ratings")
			return
		}
		defer rows.Close()

		previousRatingList := PreviousRatingList{}
		previousRatingList.Previousratings = []GetPreviousResult{}

		var id string
		var username string
		var rating string
		var comment string

		for rows.Next() {
			err = rows.Scan(&id, &username, &rating, &comment)

			previousRatingList.Previousratings = append(previousRatingList.Previousratings, GetPreviousResult{id, username, rating, comment})

		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from Previous ratings list...")
			return
		}

		js, jserr := json.Marshal(previousRatingList)

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

func (s *Server) handlegetinterestedbuyers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		interested := InterestedBuyers{}
		err := json.NewDecoder(r.Body).Decode(&interested)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to send message content ")
			return
		}

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.getinterestedbuyers('" + interested.UserID + "','" + interested.AdvertisementID + "')")
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		interestedRatingList := InterestedRatingList{}
		interestedRatingList.Interestedbuyers = []GetInterestedBuyersResult{}

		var username string
		var advertisementid string
		var advertisementtype string
		var sellerid string
		var buyerid string

		for rows.Next() {
			err = rows.Scan(&username, &advertisementid, &advertisementtype, &sellerid, &buyerid)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from Interestedratings List...")
				fmt.Println(err.Error())
				return
			}
			interestedRatingList.Interestedbuyers = append(interestedRatingList.Interestedbuyers, GetInterestedBuyersResult{username, advertisementid, advertisementtype, sellerid, buyerid})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from Interestedratings List...")
			return
		}

		js, jserr := json.Marshal(interestedRatingList)

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
func (s *Server) handlegetratingstodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle getratingstodoHas Been Called..")

		// retrieving the ID of the user that is requested to be deleted.
		userid := r.URL.Query().Get("userid")

		// declaring variable to catch response from database.
		var outstandingRatings bool

		// building query string.
		querystring := "SELECT * FROM public.ratingstodo('" + userid + "')"

		// querying the database and reading response from database into variable.
		err := s.dbAccess.QueryRow(querystring).Scan(&outstandingRatings)

		// check for errors with reading response from database into variables.
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, err.Error())
			fmt.Println("Error in communicating with database to the selected chat")
			return
		}

		// declaring result struct for delete user.
		unreadResult := OutstandingRatingResult{}
		unreadResult.OutstandingRatings = outstandingRatings

		// convert struct into JSON payload to send to service that called this function.
		js, jserr := json.Marshal(unreadResult)

		// check to see if any errors occured with coverting to JSON.
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to delete user")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)

	}
}

func (s *Server) handlegetbuyerdashboard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle getbuyerdashboard Been Called..")

		// retrieving the ID of the user that is requested to be deleted.
		userid := r.URL.Query().Get("userid")

		// declaring variable to catch response from database.
		var average string

		// building query string.
		querystring := "SELECT * FROM public.buyerdashboard('" + userid + "')"

		// querying the database and reading response from database into variable.
		err := s.dbAccess.QueryRow(querystring).Scan(&average)

		// check for errors with reading response from database into variables.
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, err.Error())
			fmt.Println("Error in communicating with database to the selected chat")
			return
		}

		// declaring result struct for delete user.
		averageResult := AverageResult{}

		if average == "" {
			averageResult.Average = "Work"
		} else {
			newaverage := strings.Replace(average, "0", "", -1)
			averageResult.Average = newaverage
		}

		// convert struct into JSON payload to send to service that called this function.
		js, jserr := json.Marshal(averageResult)

		// check to see if any errors occured with coverting to JSON.
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to delete user")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)

	}
}

func (s *Server) handlegetsellerdashboard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle getsellerdashboard Been Called..")

		// retrieving the ID of the user that is requested to be deleted.
		userid := r.URL.Query().Get("userid")

		// declaring variable to catch response from database.
		var average string

		// building query string.
		querystring := "SELECT * FROM public.sellerdashboard('" + userid + "')"

		// querying the database and reading response from database into variable.
		err := s.dbAccess.QueryRow(querystring).Scan(&average)

		// check for errors with reading response from database into variables.
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, err.Error())
			fmt.Println("Error in communicating with database to the selected chat")
			return
		}

		// declaring result struct for delete user.
		averageResult := AverageResult{}

		if average == "" {
			averageResult.Average = "Work"
		} else {
			newaverage := strings.Replace(average, "0", "", -1)
			averageResult.Average = newaverage
		}
		// convert struct into JSON payload to send to service that called this function.
		js, jserr := json.Marshal(averageResult)

		// check to see if any errors occured with coverting to JSON.
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to delete user")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)

	}
}
