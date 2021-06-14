package handlers

import (
	"encoding/json"
	"net/http"
)

// swagger:operation GET /report URLReport
//
// Returns report information
//
// ---
// produces:
// - application/json
//
// responses:
//   200:
//     description: returns report information
//     schema:
//       type: array
//       items:
//         type: object
//         properties:
//           fullURL:
//             type: string
//             description: Full URL
//           visits:
//             type: integer
//             description: Visit counts of URL
func (h *Handler) ReportEndpoint(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
	all, err := h.container.ShortenerService.FindAll()
	if err != nil {
		return
	}

	err = json.NewEncoder(res).Encode(all)
	if err != nil {
		return
	}
}
