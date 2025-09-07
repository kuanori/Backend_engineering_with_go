package main

import (
	"app/internal/repository"
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type userKey string

const userCtx userKey = "user"

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromContext(r)

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

type FollowerUser struct {
	UserID int64 `json:"user_id"`
}

func (app *application) followUserHandler(w http.ResponseWriter, r *http.Request) {
	followerUser := getUserFromContext(r)

	// TODO: Revert back to auth userID from ctx
	var paylod FollowerUser
	if err := readJSON(w, r, &paylod); err != nil {
		app.badRequestResponse(w, r, err)
	}

	ctx := r.Context()

	if err := app.repository.Followers.Follow(ctx, followerUser.ID, paylod.UserID); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) unfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	unfollowedUser := getUserFromContext(r)

	// TODO: Revert back to auth userID from ctx
	var paylod FollowerUser
	if err := readJSON(w, r, &paylod); err != nil {
		app.badRequestResponse(w, r, err)
	}

	ctx := r.Context()

	if err := app.repository.Followers.Unfollow(ctx, unfollowedUser.ID, paylod.UserID); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) userContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)
		if err != nil {
			app.badRequestResponse(w, r, err)
			return
		}

		ctx := r.Context()

		user, err := app.repository.Users.GetById(ctx, userID)
		if err != nil {
			if err == repository.ErrNotFound {
				app.notFoundResponse(w, r, err)
				return
			}
			app.internalServerError(w, r, err)
			return
		}

		ctx = context.WithValue(ctx, userCtx, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserFromContext(r *http.Request) *repository.User {
	user, _ := r.Context().Value(userCtx).(*repository.User)
	return user
}
