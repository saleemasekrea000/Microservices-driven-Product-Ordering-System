package http

import (
	"net/http"
)

// errorResponse is a generic helper function for sending JSON-formatted error messages to the
// client with a given status code. It takes an interface{} type as the message parameter, which
// gives us more flexibility over the values we can include in the response.
//
// The function first constructs an envelope (a simple map[string]interface{}) containing the
// error message. The envelope is then passed to the writeJSON() helper function, which is
// responsible for formatting the JSON response and writing it to the client.
//
// If writeJSON() returns an error, we log it and fall back to sending an empty response with
// a 500 Internal Server Error status code to the client. This is a last-ditch effort to handle
// unexpected errors gracefully.
func errorResponse(w http.ResponseWriter, _ *http.Request, status int, message interface{}) {
	// Construct the envelope (map) that will contain our error message
	env := Envelope{"error": message}

	// Write the response using the writeJSON() helper function. If this function returns an error,
	// we log it and fall back to sending an empty response with a 500 Internal Server Error
	// status code to the client.
	err := WriteJSON(w, status, env, nil)
	if err != nil {

		// Set the status code to 500 Internal Server Error, as a last-ditch effort to handle the
		// error gracefully.
		w.WriteHeader(500)
	}
}

// serverErrorResponse method is used when our application encounters an unexpected problem
// at runtime. it logs the detailed error message, then uses the errorResponse() helper to send a
// 500 Internal Server Error status code and JSON response (containing the generic error message)
// to the client
func ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {

	message := "the server encountered a problem and could not process your request"
	errorResponse(w, r, 500, message)
}

// notFoundResponse method is used to send a 404 Not Found status code and JSON response to the
// client.
func NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	errorResponse(w, r, http.StatusNotFound, message)
}
