package helpers

// createShortURLDoc
// @Summary      Create a short URL
// @Description  Takes a long URL and returns a unique 6-character short code.
// @Tags         URL Operations
// @Accept       json
// @Produce      json
// @Param        request  body      api.Body		true  "Long URL Payload"
// @Success      200      {object}  models.Response
// @Failure      400      {object}  models.Response
// @Router       /create [post]
func createShortURLDoc() {}

// redirectURLDoc
// @Summary      Redirect to Original URL
// @Description  Redirects the user to the original long URL using the short code.
// @Tags         URL Operations
// @Param        code  path  string  true  "The 6-character short code"
// @Success      301   {string} string "Redirecting..."
// @Failure      404   {object}  models.Response
// @Router       /{code} [get]
func redirectURLDoc() {}