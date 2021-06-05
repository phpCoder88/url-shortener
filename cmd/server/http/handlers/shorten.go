package handlers

import "net/http"

// swagger:operation POST /shorten shortenURL
//
// Creates a new short URL for given URL
//
// ---
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// parameters:
//   - in: body
//     name: url
//     description: URL to shorten
//     schema:
//       type: object
//       required:
//         - fullURL
//       properties:
//         fullURL:
//           type: string
//
// Responses:
//   200:
//     description: Short URL already exists
//     schema:
//       type: object
//       properties:
//         shortURL:
//           type: string
//           description: Short URL
//   201:
//     description: Created short URL
//     schema:
//       type: object
//       properties:
//         shortURL:
//           type: string
//           description: Short URL
//   405:
//     description: "Invalid input"
func ShortenEndpoint(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
}
