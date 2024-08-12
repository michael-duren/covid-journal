package handlers

import (
	"context"
	"covid-journal/internal/auth"
	"covid-journal/internal/database"
	"covid-journal/internal/server/middleware"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

func AuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	logger := middleware.GetLoggingContext(r.Context())
	queryCtx := middleware.GetQueryContext(r.Context())
	provider := chi.URLParam(r, "provider")
	logger.Infof("Provider is %s", provider)

	logger.Infof("In GetAuthCallbackFunction")
	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))

	logger.Infof("Calling gothic.CompleteUserAuth")
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Println("Error")
		fmt.Printf("%+v\n", err)
		http.Redirect(w, r, "/error", http.StatusInternalServerError)
	}

	logger.Infof("Calling gothic.CompleteUserAuth succeeded")

	dbUser, err := getDbUser(&user, queryCtx, r.Context())
	if errors.Is(err, sql.ErrNoRows) {
		dbUser, err = createDbUser(&user, queryCtx, r.Context())
		if err != nil {
			logger.Warnf("an error occured trying to create the db user based of the oauth information. error: %v", err)
			http.Redirect(w, r, "/error", http.StatusInternalServerError)
		}
	} else if err != nil {
		logger.Warnf("an unknown error occured with the database query. err:\n%v", err)
	}

	if err = createSession(dbUser, queryCtx, r, w); err != nil {
		logger.Warnf("unable to create user session. error: %v", err)
		http.Redirect(w, r, "/error", http.StatusInternalServerError)
	}

	// redirect after login
	http.Redirect(w, r, "/", http.StatusFound)
}

func BeginAuthHandler(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	logger := middleware.GetLoggingContext(r.Context())

	logger.Infof("Provider is %s", provider)
	if provider == "" {
		http.Error(w, "Provider is missing", http.StatusBadRequest)
		return
	}
	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))

	// try to get the user without re-authenticating
	if _, err := gothic.CompleteUserAuth(w, r); err == nil {
		fmt.Println("Hit login page without re-authenticating")
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	logger := middleware.GetLoggingContext(r.Context())
	userSession, isNewSession, err := middleware.GetUserSession(r)
	if err != nil {
		logger.Warnf("there was an error logging the user out: %v", err)
	}
	if isNewSession {
		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	userSession.Values = make(map[interface{}]interface{})
	userSession.Options.MaxAge = -1
	if err = userSession.Save(r, w); err != nil {
		logger.Warnf("unable to log user out. error: %v", err)
	}

	_ = gothic.Logout(w, r)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func getDbUser(user *goth.User, queryCtx *database.Queries, context context.Context) (database.User, error) {
	dbUser, err := queryCtx.GetUserByEmail(context, user.Email)
	if err != nil {
		return database.User{}, err
	}
	return dbUser, nil
}

func createDbUser(user *goth.User, queryCtx *database.Queries, context context.Context) (database.User, error) {
	newDbUser := database.CreateUserParams{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		OauthID: sql.NullString{
			String: user.UserID,
			Valid:  true,
		},
		AvatarUrl: sql.NullString{
			String: user.AvatarURL,
			Valid:  true,
		},
		LOCATION: sql.NullString{
			String: user.Location,
			Valid:  true,
		},
	}
	dbUser, err := queryCtx.CreateUser(context, newDbUser)
	if err != nil {
		return database.User{}, err
	}
	return dbUser, nil
}

// create a session & cookie for debuser and store in db
func createSession(dbUser database.User, queryCtx *database.Queries, r *http.Request, w http.ResponseWriter) error {
	store := middleware.GetSessionContext(r.Context())
	userSession, err := store.GetSession(r, string(auth.UserSessionK))
	if err != nil {
		return err
	}

	userSessionMap := map[string]string{
		"user-id":    string(dbUser.UserID),
		"oauth-id":   dbUser.OauthID.String,
		"first-name": dbUser.FirstName,
		"last-name":  dbUser.LastName,
		"email":      dbUser.Email,
		"avatar":     dbUser.AvatarUrl.String,
		"location":   dbUser.LOCATION.String,
	}

    sessionId := uuid.New().String()

	userSession.Values["user-id"] = userSessionMap["user-id"]
	userSession.Values["first-name"] = userSessionMap["first-name"]
	userSession.Values["last-name"] = userSessionMap["last-name"]
	userSession.Values["email"] = userSessionMap["email"]
	userSession.Values["avatar"] = userSessionMap["avatar"]
	userSession.Values["location"] = userSessionMap["location"]
    userSession.Values["session-id"] = sessionId

	err = userSession.Save(r, w)
	if err != nil {
		return err
	}

	sessionData, err := json.Marshal(userSessionMap)
	if err != nil {
		return err
	}

	// save to db
	dbSessionParams := database.CreateUserSessionParams{
		SessionID:   sessionId,
		UserID:      dbUser.UserID,
		SessionData: sessionData,
	}
	dbSession, err := queryCtx.CreateUserSession(r.Context(), dbSessionParams)
	if err != nil {
		return err
	}

	userSession.Values["session-id"] = dbSession.SessionID

	return nil
}
