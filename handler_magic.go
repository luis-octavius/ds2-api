package main

import (
	"context"
	"net/http"
)

var ctx = context.Background()

func (cfg *apiConfig) handlerGetMiracles(w http.ResponseWriter, _ *http.Request) {
	miracles, err := cfg.db.GetMiracles(ctx)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error returning the miracles from database", err)
		return
	}

	respondWithJSON(w, http.StatusOK, miracles)
}

func (cfg *apiConfig) handlerGetSorceries(w http.ResponseWriter, _ *http.Request) {
	sorceries, err := cfg.db.GetHexes(ctx)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error returning sorceries from database", err)
		return
	}

	respondWithJSON(w, http.StatusOK, sorceries)
}

func (cfg *apiConfig) handlerGetPyromancies(w http.ResponseWriter, _ *http.Request) {
	pyromancies, err := cfg.db.GetPyromancies(ctx)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error returning pyromancies from database", err)
		return
	}

	respondWithJSON(w, http.StatusOK, pyromancies)
}

func (cfg *apiConfig) handlerGetHexes(w http.ResponseWriter, _ *http.Request) {
	hexes, err := cfg.db.GetHexes(ctx)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error returning hexes from database", err)
	}

	respondWithJSON(w, http.StatusOK, hexes)
}
