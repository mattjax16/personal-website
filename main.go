package personal_website

import "net/http"

/// Here will be the main file to run the server for my website

/// Importing the necessary packages
"github.com/gorilla/mux",
		import (_ "net/http",
		"html/template",
		"fmt",
		"log",
		"os",
		"io/ioutil",
		"encoding/json",
		"database/sql",
		_ "github.com/go-sql-driver/mysql",
		"time",
		"strconv",
		"strings",
		"regexp",
		"sort",
		"math/rand",
		"bufio",
		"path/filepath",
)

/// MAIN FUNCTION
func main() {

	/// Setting up the router
	router := mux.NewRouter()

	/// Setting up the static files
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

}