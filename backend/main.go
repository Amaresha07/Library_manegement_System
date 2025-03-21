package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var secretKey = []byte("your-secret-key")

type Admin struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
type User struct {
	Usn      string `json:"usn"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Book struct {
	BookID   int    `json:"book_id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Year     int    `json:"year"`
	Count    int    `json:"count"`
	CoverURL string `json:"cover_url"`
}

var DB *sql.DB

func Register(w http.ResponseWriter, r *http.Request) {
	var loginUser Admin

	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, err := DB.Exec("INSERT INTO admin (name, password) VALUES($1, $2)", loginUser.Name, loginUser.Password)
	if err != nil {
		http.Error(w, "Failed to insert user", http.StatusInternalServerError)
		return
	}
	fmt.Println("Received:", loginUser.Name, loginUser.Password)

	response := map[string]string{
		"message": "Registered Successful",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func Login(w http.ResponseWriter, r *http.Request) {
	var loginUser Admin

	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var storedpassword string
	err := DB.QueryRow("SELECT password FROM admin  WHERE name =$1", loginUser.Name).Scan(&storedpassword)
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}
	if loginUser.Password != storedpassword {
		log.Println("Incorrect password")
	}

	response := map[string]string{
		"message": "Login Successful",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getbooks(w http.ResponseWriter, r *http.Request) {

	rows, err := DB.Query("SELECT book_id, title, author,year, count,cover_url FROM books")
	if err != nil {
		http.Error(w, "Unable to get books", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.BookID, &b.Title, &b.Author, &b.Year, &b.Count); err != nil {
			http.Error(w, "Error scanning books", http.StatusInternalServerError)
			return
		}
		books = append(books, b)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func addbook(w http.ResponseWriter, r *http.Request) {
	var book Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Unable to add book", http.StatusBadRequest)
		return
	}
	fmt.Println("Received book:", book)

	_, err := DB.Exec("INSERT INTO books( title, author, year ,count) VALUES ( $1, $2, $3,$4)",
		book.Title, book.Author, book.Year, book.Count)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to add book", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Book added successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Update a book
func updatebook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	BookID := vars["id"]

	if BookID == "" {
		http.Error(w, "Missing book ID", http.StatusBadRequest)
		return
	}

	// id, err := strconv.Atoi(idstr)
	// if err != nil {
	// 	http.Error(w, "Invalid book ID", http.StatusBadRequest)
	// 	return
	// }
	var book Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	result, err := DB.Exec("UPDATE books SET title = $1, author = $2, count = $3, year=$4 WHERE book_id = $5",
		book.Title, book.Author, book.Count, book.Year, BookID)

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "No book found with given ID", http.StatusNotFound)
		log.Println("No rows affected")
		return
	}

	fmt.Println(err)

	fmt.Println("Received Data from Frontend:")
	fmt.Printf("Title: %s\n", book.Title)
	fmt.Printf("Author: %s\n", book.Author)
	fmt.Printf("Count: %v\n", book.Count)
	fmt.Printf("year: %v\n", book.Year)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to update book", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Book updated successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func Deletebook(w http.ResponseWriter, r *http.Request) {
	// var book Book
	// if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
	// 	http.Error(w, "Failed to parse request body", http.StatusBadRequest)
	// 	return
	// }
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(w, "ID parameter is missing", http.StatusBadRequest)
		return
	}

	_, err := DB.Exec("DELETE FROM books WHERE book_id = $1", id)
	if err != nil {
		http.Error(w, "Error deleting book", http.StatusInternalServerError)
		return
	}
	response := map[string]string{"message": "Book deleted successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func Serachbook(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON request body
	var searchRequest struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}

	if err := json.NewDecoder(r.Body).Decode(&searchRequest); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Prepare the SQL query with the decoded values
	rows, err := DB.Query("SELECT book_id, title, author, count FROM books WHERE title = $1 OR author = $2", searchRequest.Title, searchRequest.Author)

	if err != nil {
		fmt.Println("Error querying books:", err)
		http.Error(w, "Error fetching books", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Collect the books in a slice
	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.BookID, &book.Title, &book.Author, &book.Count); err != nil {
			http.Error(w, "Error scanning books", http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}

	// Send the response outside the loop
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usn := claims["usn"].(string)
		return usn, nil
	}

	return "", jwt.ErrInvalidKey
}
func borrowBook(w http.ResponseWriter, r *http.Request) {

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}

	// ✅ Extract the token part from "Bearer <token>"
	tokenString := authHeader[len("Bearer "):]

	// ✅ Verify the token
	usn, err := VerifyToken(tokenString)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	BookID := vars["id"]

	fmt.Println("Received BookID:", BookID)

	if BookID == "" {
		http.Error(w, "ID parameter is missing", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(BookID)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	tx, err := DB.Begin()
	if err != nil {
		http.Error(w, "Transaction error", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback() // Ensure rollback if anything fails

	// Check if book is available
	var count int
	var bookName string
	err = tx.QueryRow("SELECT count, title FROM books WHERE book_id = $1", id).Scan(&count, &bookName)
	if err != nil || count <= 0 {
		http.Error(w, "Book not found", http.StatusBadRequest)
		return
	}

	if count <= 0 {
		http.Error(w, "Book not available", http.StatusBadRequest)
		return
	}

	// Reduce book count
	_, err = tx.Exec("UPDATE books SET count = count - 1 WHERE book_id = $1 AND count > 0", BookID)
	if err != nil {
		http.Error(w, "Failed to update book count", http.StatusInternalServerError)
		return
	}

	// Insert into borrowed_books

	_, err = tx.Exec("INSERT INTO borrowed_books (usn,bookname, book_id) VALUES ($1, $2,$3)", usn, bookName, BookID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to record borrowing", http.StatusInternalServerError)
		return
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		http.Error(w, "Transaction commit failed", http.StatusInternalServerError)
		return
	}

	// Success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Book borrowed successfully"})
}
func returnBook(w http.ResponseWriter, r *http.Request) {
	var returnRequest struct {
		Usn      string `json:"usn"`
		BookName string `json:"bookname"`
		BookID   int    `json:"book_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&returnRequest); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	tx, err := DB.Begin()
	if err != nil {
		http.Error(w, "Transaction error", http.StatusInternalServerError)
		return
	}

	_, err = tx.Exec("DELETE FROM borrowed_books WHERE usn = $1 AND book_id = $2", returnRequest.Usn, returnRequest.BookID)
	if err != nil {
		tx.Rollback()
		return
	}

	_, err = tx.Exec("UPDATE books SET count = count + 1 WHERE book_id = $1", returnRequest.BookID)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Book returned successfully"})
}

// ✅ Function to generate JWT token
func GenerateToken(usn string) (string, error) {
	claims := jwt.MapClaims{
		"usn": usn,                                   // User identifier
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	// Create the token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	return token.SignedString(secretKey)
}

func Loginuser(w http.ResponseWriter, r *http.Request) {
	var loginUser User

	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var storedpassword string
	err := DB.QueryRow("SELECT password FROM users5  WHERE usn =$1", loginUser.Usn).Scan(&storedpassword)
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	if loginUser.Password != storedpassword {
		log.Println("Incorrect password") // ✅ Log the error
		http.Error(w, "Incorrect password", http.StatusUnauthorized)
		return // ✅ Stop further execution
	}
	token, err := GenerateToken(loginUser.Usn)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": "Login Successful",
		"token":   token,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func Registeruser(w http.ResponseWriter, r *http.Request) {
	var loginUser User

	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// fmt.Println(loginUser)
	// fmt.Println("Received:", loginUser.Usn, loginUser.Name, loginUser.Password)

	_, err := DB.Exec("INSERT INTO users5 (usn,name, password) VALUES($1, $2,$3)", loginUser.Usn, loginUser.Name, loginUser.Password)
	if err != nil {
		http.Error(w, "Failed to insert user", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": "Registered Successful",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	db, err := sql.Open("postgres", "postgresql://demo_ub3u_user:TNVfdu2b716AGyj0wFfUzeK7BKB4GLMD@dpg-cus4n3qn91rc73dhmpvg-a.oregon-postgres.render.com/demo_ub3u")
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	DB = db
	if err := db.Ping(); err != nil {
		log.Fatal("Database is not reachable:", err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("OPTIONS")

	router.HandleFunc("/register", Register).Methods("POST")
	router.HandleFunc("/login", Login).Methods("POST")
	router.HandleFunc("/userregister", Registeruser).Methods("POST")
	router.HandleFunc("/userlogin", Loginuser).Methods("POST")
	router.HandleFunc("/books", getbooks).Methods("GET")
	router.HandleFunc("/books/add", addbook).Methods("POST")
	router.HandleFunc("/update/{id}", updatebook).Methods("PUT")
	router.HandleFunc("/books/delete/{id}", Deletebook).Methods("DELETE")
	router.HandleFunc("/books/search", Serachbook).Methods("POST")
	router.HandleFunc("/borrow/{id}", borrowBook).Methods("POST")
	router.HandleFunc("/books/return", returnBook).Methods("POST")

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}), // Allow all origins (*), change this to specific frontend URL in production
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)

	fmt.Println("Server is running at port: 9090")
	log.Fatal(http.ListenAndServe(":9090", corsHandler(router)))
}
