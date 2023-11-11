package handlers

import (
	"net/http"

	"dev.azure.com/wole0010243/scheduler/_git/scheduler/util"
)

func HandleReadiness(w http.ResponseWriter, r *http.Request) {
    util.RespondWithJSON(w, 200, struct{}{})
}
