package main

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type Server struct {
	dbAccess *sql.DB
	router   *mux.Router
}

type UserID struct {
	UserID string `json:"id"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
}

type getUser struct {
	UserID   string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
}

type updateUser struct {
	UserID   string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
}

type UpdateUserResult struct {
	UserUpdated bool   `json:"userupdated"`
	Message     string `json:"message"`
}

type DeleteUserResult struct {
	UserDeleted bool   `json:"userdeleted"`
	UserID      string `json:"id"`
	Message     string `json:"message"`
}

type LoginUserResult struct {
	UserID       string `json:"id"`
	Username     string `json:"username"`
	UserLoggedIn bool   `json:"userloggedin"`
	Message string `json:"message"`
}

type RegisterUserResult struct {
	UserCreated string `json:"usercreated"`
	Username    string `json:"username"`
	UserID      string `json:"id"`
	Message     string `json:"message"`
}

type dbConfig struct {
	UserName        string
	Password        string
	DatabaseName    string
	Port            string
	PostgresHost    string
	PostgresPort    string
	ListenServePort string
}

//advert crud
type PostAdvertisement struct {
	UserID            string `json:"userid"`
	AdvertisementType string `json:"advertisementtype"`
	EntityID          string `json:"entityid"`
	Price             string `json:"price"`
	Description       string `json:"description"`
}

type PostAdvertisementResult struct {
	AdvertisementPosted bool   `json:"advertisementposted"`
	ID                  string `json:"id"`
	Message             string `json:"message"`
}

type UpdateAdvertisement struct {
	AdvertisementID   string `json:"id"`
	UserID            string `json:"userid"`
	AdvertisementType string `json:"advertisementtype"`
	EntityID          string `json:"entityid"`
	Price             string `json:"price"`
	Description       string `json:"description"`
}

type UpdateAdvertisementResult struct {
	AdvertisementUpdated bool   `json:"advertisementupdated"`
	Message              string `json:"message"`
}

type DeleteAdvertisementResult struct {
	AdvertisementDeleted bool   `json:"advertisementdeleted"`
	AdvertisementID      string `json:"id"`
	Message              string `json:"message"`
}

type AdvertisementID struct {
	AdvertisementID string `json:"id"`
}

type getAdvertisement struct {
	AdvertisementID   string `json:"id"`
	UserID            string `json:"userid"`
	AdvertisementType string `json:"advertisementtype"`
	EntityID          string `json:"entityid"`
	Price             string `json:"price"`
	Description       string `json:"description"`
}


type TypeAdvertisementList struct {
	TypeAdvertisements []getAdvertisement `json:"typeadvertisements"`
}

type AdvertisementList struct {
	Advertisements []getAdvertisement `json:"advertisements"`
}

type GetUserAdvertisementResult struct {
	AdvertisementID   string `json:"advertisementid"`
	AdvertisementType string `json:"advertisementtype"`
	EntityID          string `json:"entityid"`
	Price             string `json:"price"`
	Description       string `json:"description"`
}

type UserAdvertisementList struct {
	UserAdvertisements []GetUserAdvertisementResult `json:"useradvertisements"`
}

type Textbook struct {
	ModuleCode string `json:"modulecode"`
	Name string `json:"name"`
	Edition string `json:"edition"`
	Quality string `json:"quality"`
	Author string `json:"author"`
}

type TextbookResult struct {
	TextbookAdded bool `json:"textbookadded"`
	TextbookID string `json:"id"`
	Message string `json:"message"`
}

type UpdateTextbook struct {
	TextbookID string `json:"id"`
	ModuleCode string `json:"modulecode"`
	Name string `json:"name"`
	Edition string `json:"edition"`
	Quality string `json:"quality"`
	Author string `json:"author"`
}

type UpdateTextbookResult struct {
	TextbookUpdated bool `json:"textbookupdated"`
	Message string `json:"message"`
}

type TextbookFilter struct {
	ModuleCode string `json:"modulecode"`
	Name string `json:"name"`
	Edition string `json:"edition"`
	Quality string `json:"quality"`
	Author string `json:"author"`
}

type TextbookFilterResult struct {
	ModuleCode string `json:"modulecode"`
	ID string `'json:"id"`
	Name string `json:"name"`
	Edition string `json:"edition"`
	Quality string `json:"quality"`
	Author string `json:"author"`
}


type TextbookList struct {
	Textbooks []TextbookFilterResult `json:"textbooks"`
}

type DeleteTextbookResult struct {
	TextbookDeleted bool   `json:"Textbookdeleted"`
	TextbookID      string `json:"id"`
	Message              string `json:"message"`
}

type Note struct {
	ModuleCode string `json:"modulecode"`
}

type NoteResult struct {
	NoteAdded bool `json:"noteadded"`
	NoteID string `json:"id"`
	Message string `json:"message"`
}

type UpdateNote struct {
	NoteID string `json:"id"`
	ModuleCode string `json:"modulecode"`
}

type UpdateNoteResult struct {
	NoteUpdated bool `json:"noteupdated"`
	Message string `json:"message"`
}

type NoteFilter struct {
	ModuleCode string `json:"modulecode"`
}

type NoteFilterResult struct {
	ID string `json:"id"`
	ModuleCode string `json:"modulecode"`
}

type NoteList struct {
	Notes []NoteFilterResult `json:"notes"`
}

type DeleteNoteResult struct {
	NoteDeleted bool   `json:"Notedeleted"`
	NoteID      string `json:"id"`
	Message              string `json:"message"`
}

type Tutor struct {
	ModuleCode string `json:"modulecode"`
	Subject string `json:"subject"`
	YearCompleted string `json:"yearcompleted"`
	Venue string `json:"venue"`
	NotesIncluded string `json:"notesincluded"`
	Terms string `json:"terms"`
}

type TutorResult struct {
	TutorAdded bool `json:"tutoradded"`
	TutorID string `json:"id"`
	Message string `json:"message"`
}

type UpdateTutor struct {
	TutorID string `json:"id"`
	ModuleCode string `json:"modulecode"`
	Subject string `json:"subject"`
	YearCompleted string `json:"yearcompleted"`
	Venue string `json:"venue"`
	NotesIncluded string `json:"notesincluded"`
	Terms string `json:"terms"`
}

type UpdateTutorResult struct {
	TutorUpdated bool `json:"tutorupdated"`
	Message string `json:"message"`
}


type TutorFilter struct {
	ModuleCode string `json:"modulecode"`
	Subject string `json:"subject"`
	YearCompleted string `json:"yearcompleted"`
	Venue string `json:"venue"`
	NotesIncluded string `json:"notesincluded"`
	Terms string `json:"terms"`
}

type TutorFilterResult struct {
	ID string `json:"id"`
	ModuleCode string `json:"modulecode"`
	Subject string `json:"subject"`
	YearCompleted string `json:"yearcompleted"`
	Venue string `json:"venue"`
	NotesIncluded string `json:"notesincluded"`
	Terms string `json:"terms"`
}

type TutorList struct {
	Tutors []TutorFilterResult `json:"tutors"`
}

type DeleteTutorResult struct {
	TutorDeleted bool   `json:"Tutordeleted"`
	TutorID      string `json:"id"`
	Message              string `json:"message"`
}

type Accomodation struct {
	AccomodationTypeCode string `json:"accomodationtypecode"`
	InstitutionName string `json:"institutionname"`
	Location string `json:"location"`
	DistanceToCampus string `json:"distancetocampus"`
}

type AccomodationResult struct {
	AccomodationAdded bool `json:"accomodationadded"`
	AccomodationID string `json:"id"`
	Message string `json:"message"`
}

type UpdateAccomodation struct {
	AccomodationID string `json:"id"`
	AccomodationTypeCode string `json:"accomodationtypecode"`
	InstitutionName string `json:"institutionname"`
	Location string `json:"location"`
	DistanceToCampus string `json:"distancetocampus"`
}

type UpdateAccomodationResult struct {
	AccomodationUpdated bool `json:"accomodationupdated"`
	Message string `json:"message"`
}


type AccomodationFilter struct {
	AccomodationTypeCode string `json:"accomodationtypecode"`
	InstitutionName string `json:"institutionname"`
	Location string `json:"location"`
	DistanceToCampus string `json:"distancetocampus"`
}

type AccomodationFilterResult struct {
	ID string `json:"id"`
	AccomodationTypeCode string `json:"accomodationtypecode"`
	InstitutionName string `json:"institutionname"`
	Location string `json:"location"`
	DistanceToCampus string `json:"distancetocampus"`
}

type AccomodationList struct {
	Accomodations []AccomodationFilterResult `json:"accomodations"`
}

type DeleteAccomodationResult struct {
	AccomodationDeleted bool   `json:"Accomodationdeleted"`
	AccomodationID      string `json:"id"`
	Message              string `json:"message"`
}

type DeleteAdvertisementsResult struct {
	AdvertisementsDeleted bool   `json:"advertisementsdeleted"`
	Message               string `json:"message"`
}
type Config struct {
	ListenServePort string
}
