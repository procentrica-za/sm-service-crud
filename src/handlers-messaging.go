package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (s *Server) handleaddchat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handleaddchat Has Been Called!")
		//get JSON payload
		chat := StartChat{}
		err := json.NewDecoder(r.Body).Decode(&chat)
		fmt.Println("Handle add chat Has Been Called..")
		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to add a new chat ")
			return
		}
		//set response variables
		var chatposted bool
		var id string

		//communcate with the database
		querystring := "SELECT * FROM public.addchat('" + chat.SellerID + "','" + chat.BuyerID + "','" + chat.AdvertisementType + "','" + chat.AdvertisementID + "')"

		//retrieve message from database tt set to JSON object
		err = s.dbAccess.QueryRow(querystring).Scan(&chatposted, &id)

		//check for response error of 500
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to add chat")
			fmt.Println(err.Error())
			fmt.Println("Error in communicating with database to add new chat")
			return
		}

		//set JSON object variables for response
		startchatResult := StartChatResult{}
		startchatResult.ChatPosted = chatposted
		startchatResult.ChatID = id

		if chatposted {
			startchatResult.Message = "Chat Successfully Started!"
		} else {
			startchatResult.Message = "Unable to start chat!"
		}

		//convert struct back to JSON
		js, jserr := json.Marshal(startchatResult)

		//error occured when trying to convert struct to a JSON object
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to post chat")
			return
		}

		//return back to advert service
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handledeletechat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Delete Chat Has Been Called..")

		// retrieving the ID of the user that is requested to be deleted.
		getchatid := r.URL.Query().Get("chatid")
		chatid := ChatID{}
		chatid.ChatID = getchatid

		// declaring variable to catch response from database.
		var chatDeleted bool

		// building query string.
		querystring := "SELECT * FROM public.deletechat('" + chatid.ChatID + "')"

		// querying the database and reading response from database into variable.
		err := s.dbAccess.QueryRow(querystring).Scan(&chatDeleted)

		// check for errors with reading response from database into variables.
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, err.Error())
			fmt.Println("Error in communicating with database to the selected chat")
			return
		}

		// declaring result struct for delete user.
		deleteChatResult := DeleteChatResult{}
		deleteChatResult.ChatDeleted = chatDeleted

		// determine if user was infact deleted.
		if chatDeleted {
			deleteChatResult.Message = "Chat Successfully Deleted!"
		} else {
			deleteChatResult.Message = "Unable to Selected Chat!"
		}

		// convert struct into JSON payload to send to service that called this function.
		js, jserr := json.Marshal(deleteChatResult)

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

func (s *Server) handlegetactivechats() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle get active chats Has Been Called..")
		userid := r.URL.Query().Get("userid")

		rows, err := s.dbAccess.Query("SELECT * FROM public.getactivechats('" + userid + "')")
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to get active chats")
			return
		}
		defer rows.Close()

		activeChatList := ActiveChatList{}
		activeChatList.ActiveChats = []GetActiveChatResult{}

		var id string
		var username string
		var message string
		var messagedate string
		var advertisementtype string
		var advertisementid string
		var price string
		var title string
		var description string
		var isread string
		var messageauthor string

		for rows.Next() {
			id = ""
			username = ""
			message = ""
			messagedate = ""
			advertisementtype = ""
			advertisementid = ""
			price = ""
			title = ""
			description = ""
			isread = ""
			messageauthor = ""

			err = rows.Scan(&id, &advertisementtype, &advertisementid, &username, &price, &title, &description, &message, &messagedate, &isread, &messageauthor)
			if messagedate == "" && message == "" {
				activeChatList.ActiveChats = append(activeChatList.ActiveChats, GetActiveChatResult{id, advertisementtype, advertisementid, username, price, title, description, "Please select chat to send a message.", "", "", ""})
			} else {
				r := strings.NewReplacer("T", " ", "Z", "")
				newmessagedate := r.Replace(messagedate)
				newmessagedate = newmessagedate[:len(newmessagedate)-10]
				activeChatList.ActiveChats = append(activeChatList.ActiveChats, GetActiveChatResult{id, advertisementtype, advertisementid, username, price, title, description, message, newmessagedate, isread, messageauthor})
			}
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from Active chat List...")
			return
		}

		js, jserr := json.Marshal(activeChatList)

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

func (s *Server) handlegetmessages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle  get messages Has Been Called..")
		//retrieve URL from ad service
		userid := r.URL.Query().Get("userid")
		chatid := r.URL.Query().Get("chatid")

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.getchat('" + userid + "','" + chatid + "')")
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		messagesList := MessageList{}
		messagesList.Messages = []GetMessageResult{}

		var messageid string
		var username string
		var message string
		var messagedate string

		for rows.Next() {
			err = rows.Scan(&messageid, &username, &message, &messagedate)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from Messages List...")
				fmt.Println(err.Error())
				return
			}
			r := strings.NewReplacer("T", " ", "Z", "")
			newmessagedate := r.Replace(messagedate)
			newmessagedate = newmessagedate[:len(newmessagedate)-10]
			messagesList.Messages = append(messagesList.Messages, GetMessageResult{messageid, username, message, newmessagedate})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from Advertisement List...")
			return
		}

		js, jserr := json.Marshal(messagesList)

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

func (s *Server) handleaddmessage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		messagecontent := SendMessage{}
		err := json.NewDecoder(r.Body).Decode(&messagecontent)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to send message content ")
			return
		}

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.sendmessage('" + messagecontent.ChatID + "','" + messagecontent.AuthorID + "','" + messagecontent.Message + "')")
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		messagesList := MessageList{}
		messagesList.Messages = []GetMessageResult{}

		var messageid string
		var username string
		var message string
		var messagedate string

		for rows.Next() {
			err = rows.Scan(&messageid, &username, &message, &messagedate)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from Messages List...")
				fmt.Println(err.Error())
				return
			}
			r := strings.NewReplacer("T", " ", "Z", "")
			newmessagedate := r.Replace(messagedate)
			newmessagedate = newmessagedate[:len(newmessagedate)-10]
			messagesList.Messages = append(messagesList.Messages, GetMessageResult{messageid, username, message, newmessagedate})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from Advertisement List...")
			return
		}

		js, jserr := json.Marshal(messagesList)

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

func (s *Server) handlegetunreadmessages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Delete Chat Has Been Called..")

		// retrieving the ID of the user that is requested to be deleted.
		userid := r.URL.Query().Get("userid")

		// declaring variable to catch response from database.
		var unreadMessages bool

		// building query string.
		querystring := "SELECT * FROM public.unreadmessages('" + userid + "')"

		// querying the database and reading response from database into variable.
		err := s.dbAccess.QueryRow(querystring).Scan(&unreadMessages)

		// check for errors with reading response from database into variables.
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, err.Error())
			fmt.Println("Error in communicating with database to the selected chat")
			return
		}

		// declaring result struct for delete user.
		unreadResult := UnreadResult{}
		unreadResult.UnreadMessages = unreadMessages

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
