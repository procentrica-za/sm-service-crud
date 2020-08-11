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
	Username       string `json:"username"`
	Password       string `json:"password"`
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Email          string `json:"email"`
	InsitutionName string `json:"institutionname"`
}

type getUser struct {
	UserID         string `json:"id"`
	Username       string `json:"username"`
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Email          string `json:"email"`
	InsitutionName string `json:"institutionname"`
	AdsRemaining   string `json:"adsremaining"`
	Message        string `json:"message"`
	GotUser        bool   `json:"gotuser"`
}

type updateUser struct {
	UserID         string `json:"id"`
	Username       string `json:"username"`
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Email          string `json:"email"`
	InsitutionName string `json:"institutionname"`
}

type UpdatePassword struct {
	UserID          string `json:"id"`
	CurrentPassword string `json:"currentpassword"`
	Password        string `json:"password"`
}

type UpdatePasswordResult struct {
	PasswordUpdated bool   `json:"passwordupdated"`
	Message         string `json:"message"`
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
	Institution  string `json:"institution"`
	Message      string `json:"message"`
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

//Forgot password
type UserEmail struct {
	Email string `json:"email"`
}

type ForgotPasswordResult struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Message  string `json:"message"`
}

//advert crud
type PostAdvertisement struct {
	UserID            string `json:"userid"`
	IsSelling         string `json:"isselling"`
	AdvertisementType string `json:"advertisementtype"`
	EntityID          string `json:"entityid"`
	Price             string `json:"price"`
	Description       string `json:"description"`
	Institution       string `json:"institution"`
}

type PostAdvertisementResult struct {
	AdvertisementPosted bool   `json:"advertisementposted"`
	ID                  string `json:"id"`
	Message             string `json:"message"`
}

type UpdateAdvertisement struct {
	AdvertisementID   string `json:"id"`
	UserID            string `json:"userid"`
	IsSelling         string `json:"isselling"`
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
	IsSelling         bool   `json:"isselling"`
	AdvertisementType string `json:"advertisementtype"`
	EntityID          string `json:"entityid"`
	Price             string `json:"price"`
	Description       string `json:"description"`
	Message           string `json:"message"`
}

type getAdvertisements struct {
	AdvertisementID   string `json:"id"`
	UserID            string `json:"userid"`
	IsSelling         bool   `json:"isselling"`
	AdvertisementType string `json:"advertisementtype"`
	EntityID          string `json:"entityid"`
	Price             string `json:"price"`
	Description       string `json:"description"`
}

type TypeAdvertisementList struct {
	TypeAdvertisements []getAdvertisements `json:"typeadvertisements"`
}

type AdvertisementList struct {
	Advertisements []getAdvertisements `json:"advertisements"`
}

type GetUserAdvertisementResult struct {
	AdvertisementID   string `json:"advertisementid"`
	IsSelling         bool   `json:"isselling"`
	AdvertisementType string `json:"advertisementtype"`
	EntityID          string `json:"entityid"`
	Price             string `json:"price"`
	Description       string `json:"description"`
}

type UserAdvertisementList struct {
	UserAdvertisements []GetUserAdvertisementResult `json:"useradvertisements"`
}

type GetTextbookAdvertisementsResult struct {
	AdvertisementID   string `json:"advertisementid"`
	UserID            string `json:"userid"`
	Isselling         bool   `json:"isselling"`
	AdvertisementType string `json:"advertisementtype"`
	Price             string `json:"price"`
	Description       string `json:"description"`
	TextbookID        string `json:"textbookid"`
	TextbookName      string `json:"textbookname"`
	Edition           string `json:"edition"`
	Quality           string `json:"quality"`
	Author            string `json:"author"`
	ModuleCode        string `json:"modulecode"`
	InsitutionName    string `json:"institutionname"`
	RowNumber         string `json:"rownumber"`
}

type TextbookAdvertisementList struct {
	listEmpty bool                              `json:"listempty"`
	Textbooks []GetTextbookAdvertisementsResult `json:"textbooks"`
}

type GetTutorAdvertisementsResult struct {
	Advertisementid   string `json:"advertisementid"`
	Userid            string `json:"userid"`
	Isselling         bool   `json:"isselling"`
	Advertisementtype string `json:"advertisementtype"`
	Price             string `json:"price"`
	Description       string `json:"description"`
	Tutorid           string `json:"tutorid"`
	Subject           string `json:"subject"`
	Yearcompleted     string `json:"yearcompleted"`
	Venue             string `json:"venue"`
	Notesincluded     string `json:"notesincluded"`
	Terms             string `json:"terms"`
	Modulecode        string `json:"modulecode"`
	InsitutionName    string `json:"institutionname"`
	RowNumber         string `json:"rownumber"`
}

type TutorAdvertisementList struct {
	listEmpty bool                           `json:"listempty"`
	Tutors    []GetTutorAdvertisementsResult `json:"tutors"`
}

type GetAccomodationAdvertisementsResult struct {
	Advertisementid      string `json:"advertisementid"`
	Userid               string `json:"userid"`
	Isselling            bool   `json:"isselling"`
	Advertisementtype    string `json:"advertisementtype"`
	Price                string `json:"price"`
	Description          string `json:"description"`
	AccomodationID       string `json:"accomodationid"`
	Accomodationtypecode string `json:"accomodationtypecode"`
	Location             string `json:"location"`
	Distancetocampus     string `json:"distancetocampus"`
	InsitutionName       string `json:"institutionname"`
	RowNumber            string `json:"rownumber"`
}

type AccomodationAdvertisementList struct {
	listEmpty     bool                                  `json:"listempty"`
	Accomodations []GetAccomodationAdvertisementsResult `json:"accomodations"`
}

type GetNoteAdvertisementsResult struct {
	Advertisementid   string `json:"advertisementid"`
	Userid            string `json:"userid"`
	Isselling         bool   `json:"isselling"`
	Advertisementtype string `json:"advertisementtype"`
	Price             string `json:"price"`
	Description       string `json:"description"`
	NoteID            string `json:"noteid"`
	ModuleCode        string `json:"modulecode"`
	InsitutionName    string `json:"institutionname"`
	RowNumber         string `json:"rownumber"`
}

type NoteAdvertisementList struct {
	listEmpty bool                          `json:"listempty"`
	Notes     []GetNoteAdvertisementsResult `json:"notes"`
}

type Textbook struct {
	ModuleCode string `json:"modulecode"`
	Name       string `json:"name"`
	Edition    string `json:"edition"`
	Quality    string `json:"quality"`
	Author     string `json:"author"`
}

type TextbookResult struct {
	TextbookAdded bool   `json:"textbookadded"`
	TextbookID    string `json:"id"`
	ModuleCode    string `json:"modulecode"`
	Name          string `json:"name"`
	Edition       string `json:"edition"`
	Quality       string `json:"quality"`
	Author        string `json:"author"`
}

type UpdateTextbook struct {
	TextbookID string `json:"id"`
	ModuleCode string `json:"modulecode"`
	Name       string `json:"name"`
	Edition    string `json:"edition"`
	Quality    string `json:"quality"`
	Author     string `json:"author"`
}

type UpdateTextbookResult struct {
	TextbookUpdated bool   `json:"textbookupdated"`
	Message         string `json:"message"`
}

type TextbookFilter struct {
	ModuleCode string `json:"modulecode"`
	Name       string `json:"name"`
	Edition    string `json:"edition"`
	Quality    string `json:"quality"`
	Author     string `json:"author"`
}

type TextbookFilterResult struct {
	ModuleCode string `json:"modulecode"`
	ID         string `'json:"id"`
	Name       string `json:"name"`
	Edition    string `json:"edition"`
	Quality    string `json:"quality"`
	Author     string `json:"author"`
}

type TextbookList struct {
	Textbooks []TextbookFilterResult `json:"textbooks"`
}

type DeleteTextbookResult struct {
	TextbookDeleted bool   `json:"Textbookdeleted"`
	TextbookID      string `json:"id"`
	Message         string `json:"message"`
}

type Note struct {
	ID         string `json:"id"`
	ModuleCode string `json:"modulecode"`
}

type NoteResult struct {
	NoteAdded bool   `json:"noteadded"`
	NoteID    string `json:"id"`
	Message   string `json:"message"`
}

type UpdateNote struct {
	NoteID     string `json:"id"`
	ModuleCode string `json:"modulecode"`
}

type UpdateNoteResult struct {
	NoteUpdated bool   `json:"noteupdated"`
	Message     string `json:"message"`
}

type NoteFilter struct {
	ModuleCode string `json:"modulecode"`
}

type NoteFilterResult struct {
	ID         string `json:"id"`
	ModuleCode string `json:"modulecode"`
}

type NoteList struct {
	Notes []NoteFilterResult `json:"notes"`
}

type DeleteNoteResult struct {
	NoteDeleted bool   `json:"Notedeleted"`
	NoteID      string `json:"id"`
	Message     string `json:"message"`
}

type Tutor struct {
	ID            string `json:"id"`
	ModuleCode    string `json:"modulecode"`
	Subject       string `json:"subject"`
	YearCompleted string `json:"yearcompleted"`
	Venue         string `json:"venue"`
	NotesIncluded string `json:"notesincluded"`
	Terms         string `json:"terms"`
}

type TutorResult struct {
	TutorAdded bool   `json:"tutoradded"`
	TutorID    string `json:"id"`
	Message    string `json:"message"`
}

type UpdateTutor struct {
	TutorID       string `json:"id"`
	ModuleCode    string `json:"modulecode"`
	Subject       string `json:"subject"`
	YearCompleted string `json:"yearcompleted"`
	Venue         string `json:"venue"`
	NotesIncluded string `json:"notesincluded"`
	Terms         string `json:"terms"`
}

type UpdateTutorResult struct {
	TutorUpdated bool   `json:"tutorupdated"`
	Message      string `json:"message"`
}

type TutorFilter struct {
	ModuleCode    string `json:"modulecode"`
	Subject       string `json:"subject"`
	YearCompleted string `json:"yearcompleted"`
	Venue         string `json:"venue"`
	NotesIncluded string `json:"notesincluded"`
	Terms         string `json:"terms"`
}

type TutorFilterResult struct {
	ID            string `json:"id"`
	ModuleCode    string `json:"modulecode"`
	Subject       string `json:"subject"`
	YearCompleted string `json:"yearcompleted"`
	Venue         string `json:"venue"`
	NotesIncluded string `json:"notesincluded"`
	Terms         string `json:"terms"`
}

type TutorList struct {
	Tutors []TutorFilterResult `json:"tutors"`
}

type DeleteTutorResult struct {
	TutorDeleted bool   `json:"Tutordeleted"`
	TutorID      string `json:"id"`
	Message      string `json:"message"`
}

type Accomodation struct {
	ID                   string `json:"id"`
	AccomodationTypeCode string `json:"accomodationtypecode"`
	InstitutionName      string `json:"institutionname"`
	Location             string `json:"location"`
	DistanceToCampus     string `json:"distancetocampus"`
}

type AccomodationResult struct {
	AccomodationAdded bool   `json:"accomodationadded"`
	AccomodationID    string `json:"id"`
	Message           string `json:"message"`
}

type UpdateAccomodation struct {
	AccomodationID       string `json:"id"`
	AccomodationTypeCode string `json:"accomodationtypecode"`
	InstitutionName      string `json:"institutionname"`
	Location             string `json:"location"`
	DistanceToCampus     string `json:"distancetocampus"`
}

type UpdateAccomodationResult struct {
	AccomodationUpdated bool   `json:"accomodationupdated"`
	Message             string `json:"message"`
}

type AccomodationFilter struct {
	AccomodationTypeCode string `json:"accomodationtypecode"`
	InstitutionName      string `json:"institutionname"`
	Location             string `json:"location"`
	DistanceToCampus     string `json:"distancetocampus"`
}

type AccomodationFilterResult struct {
	ID                   string `json:"id"`
	AccomodationTypeCode string `json:"accomodationtypecode"`
	InstitutionName      string `json:"institutionname"`
	Location             string `json:"location"`
	DistanceToCampus     string `json:"distancetocampus"`
}

type AccomodationList struct {
	Accomodations []AccomodationFilterResult `json:"accomodations"`
}

type DeleteAccomodationResult struct {
	AccomodationDeleted bool   `json:"Accomodationdeleted"`
	AccomodationID      string `json:"id"`
	Message             string `json:"message"`
}

type DeleteAdvertisementsResult struct {
	AdvertisementsDeleted bool   `json:"advertisementsdeleted"`
	Message               string `json:"message"`
}

type CardImage struct {
	EntityID string `json:"entityid"`
	FilePath string `json:"filepath"`
	FileName string `json:"filename"`
}

type CardImageBatch struct {
	Images []CardImage `json:"images"`
}

type CardImageRequest struct {
	EntityID string `json:"entityid"`
}

type CardImageBatchRequest struct {
	Cards []CardImageRequest `json:"cards"`
}

//Messaging types
type StartChat struct {
	SellerID          string `json:"sellerid"`
	BuyerID           string `json:"buyerid"`
	AdvertisementType string `json:"advertisementtype"`
	AdvertisementID   string `json:"advertisementid"`
}

type StartChatResult struct {
	ChatPosted bool   `json:"chatposted"`
	ChatID     string `json:"chatid"`
	Message    string `json:"message"`
}

type ChatID struct {
	ChatID string `json:"id"`
}

type DeleteChatResult struct {
	ChatDeleted bool   `json:"chatposted"`
	Message     string `json:"message"`
}

type GetActiveChatResult struct {
	ChatID            string `json:"chatid"`
	AdvertisementType string `json:"advertisementtype"`
	AdvertisementID   string `json:"advertisementid"`
	UserName          string `json:"username"`
	Price             string `json:"price"`
	Title             string `json:"title"`
	Description       string `json:"description"`
	Message           string `json:"message"`
	MessageDate       string `json:"messagedate"`
	IsRead            string `json:"isread"`
	MessageAuthor     string `json:"messageauthor"`
}

type ActiveChatList struct {
	ActiveChats []GetActiveChatResult `json:"activechats"`
}

type GetMessageResult struct {
	MessageID   string `json:"messageid"`
	UserName    string `json:"username"`
	Message     string `json:"message"`
	MessageDate string `json:"messagedate"`
}

type MessageList struct {
	Messages []GetMessageResult `json:"messages"`
}

type SendMessage struct {
	ChatID   string `json:"chatid"`
	AuthorID string `json:"authorid"`
	Message  string `json:"message"`
}

type UnreadResult struct {
	UnreadMessages bool `json:"unreadmessages"`
}

type UploadImage struct {
	EntityID    string `json:"entityid"`
	FilePath    string `json:"filepath"`
	FileName    string `json:"filename"`
	IsMainImage bool   `json:"ismainimage"`
}

type UploadImageResult struct {
	ImageInserted bool   `json:"imageinserted"`
	Message       string `json:"message"`
}

//rate types
type StartRating struct {
	AdvertisementID string `json:"advertisementid"`
	BuyerID         string `json:"buyerid"`
	SellerID        string `json:"sellerid"`
	BuyerRating     string `json:"buyerrating"`
	BuyerComments   string `json:"buyercomments"`
}

type StartRatingResult struct {
	BuyerRated bool   `json:"buyerrated"`
	RatingID   string `json:"ratingid"`
	Message    string `json:"message"`
}

type RateSeller struct {
	RatingID       string `json:"ratingid"`
	SellerRating   string `json:"sellerrating"`
	SellerComments string `json:"sellercomments"`
}

type RateSellerResult struct {
	SellerRated bool   `json:"sellerrated"`
	Message     string `json:"message"`
}

type GetOutstandingResult struct {
	RatingID    string `json:"ratingid"`
	UserName    string `json:"username"`
	Price       string `json:"price"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type OutstandingRatingList struct {
	Oustandingratings []GetOutstandingResult `json:"outstandingratings"`
}

type OutstandingRatingResult struct {
	OutstandingRatings bool `json:"outstandingratings"`
}

type GetPreviousResult struct {
	RatingID string `json:"ratingid"`
	UserName string `json:"username"`
	Rating   string `json:"rating"`
	Comment  string `json:"comment"`
}

type PreviousRatingList struct {
	Previousratings []GetPreviousResult `json:"previousratings"`
}

type InterestedBuyers struct {
	UserID          string `json:"userid"`
	AdvertisementID string `json:"advertisementid"`
}

type GetInterestedBuyersResult struct {
	UserName          string `json:"username"`
	AdvertisementID   string `json:"advertisementid"`
	AdvertisementType string `json:"advertisementtype"`
	SellerID          string `json:"sellerid"`
	BuyerID           string `json:"buyerid"`
}

type InterestedRatingList struct {
	Interestedbuyers []GetInterestedBuyersResult `json:"interestedbuyers"`
}

type ModuleCode struct {
	Modulecode string `json:"code"`
}

type ModuleCodeList struct {
	Modulecodes []ModuleCode `json:"modulecodes"`
}

type InstitutionName struct {
	Institutionname string `json:"institutionname"`
}

type InstitutionNameList struct {
	Institutionnames []InstitutionName `json:"institutionnames"`
}

type AverageResult struct {
	Average string `json:"average"`
}

type RequestOtpResult struct {
	Sent        bool   `json:"sent"`
	Message     string `'json:"message"`
	Phonenumber string `json:"phonenumber"`
	Otp         string `json:"otp"`
}

type ValidateOtp struct {
	UserID string `json:"userid"`
	Otp    string `json:"otp"`
}

type ValidateOtpResult struct {
	Validated bool   `json:"validated"`
	Message   string `'json:"message"`
}

type PurchaseAdvertisement struct {
	ID      string `json:"id"`
	Ammount string `json:"ammount"`
}
type PurchaseAdvertisementResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type getPassword struct {
	UserID   string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	GotUser  bool   `json:"gotuser"`
}

type Status struct {
	Isverified bool `json:"isverified"`
}

type Config struct {
	ListenServePort string
}
