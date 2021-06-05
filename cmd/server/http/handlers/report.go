package handlers

import "net/http"

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
func ReportEndpoint(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
}
