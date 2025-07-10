package blogs

import (
	"encoding/json"
	"net/http"
)

func GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	blogs, err := getAll()
	if err != nil {
		http.Error(w, "Error fetching blogs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogs)
}

func GetBlogById(w http.ResponseWriter, r *http.Request) {
	blogId := r.URL.Query().Get("id")
	if blogId == "" {
		http.Error(w, "Blog ID is required", http.StatusBadRequest)
		return
	}

	blog, err := getById(blogId)
	if err != nil {
		http.Error(w, "Error fetching blog", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blog)
}

func CreateBlog(w http.ResponseWriter, r *http.Request) {
	var blog Blog
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	id, err := insertOne(blog)
	if err != nil {
		http.Error(w, "Error creating blog", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func UpdateBlog(w http.ResponseWriter, r *http.Request) {
	blogId := r.URL.Query().Get("id")
	if blogId == "" {
		http.Error(w, "Blog ID is required", http.StatusBadRequest)
		return
	}

	var blog Blog
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	id, err := updateOne(blogId, blog)
	if err != nil {
		http.Error(w, "Error updating blog", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	blogId := r.URL.Query().Get("id")
	if blogId == "" {
		http.Error(w, "Blog ID is required", http.StatusBadRequest)
		return
	}

	deletedCount, err := deleteOne(blogId)
	if err != nil {
		http.Error(w, "Error deleting blog", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"deleted_count": deletedCount})
}

func DeleteBlogs(w http.ResponseWriter, r *http.Request) {
	var blogIds []string
	if err := json.NewDecoder(r.Body).Decode(&blogIds); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	deletedCount, err := deleteMany(blogIds)
	if err != nil {
		http.Error(w, "Error deleting blogs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"deleted_count": deletedCount})
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"status": "ok"}
	json.NewEncoder(w).Encode(response)
}
