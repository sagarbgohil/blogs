package blogs

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)


func V1Router() http.Handler {
    r := chi.NewRouter()

    r.Post("/", CreateBlog)
    r.Get("/", GetAllBlogs)
    r.Get("//{id}", GetBlogById)
    r.Delete("//{id}", DeleteBlog)
	r.Put("//{id}", UpdateBlog)
	r.Delete("//bulk", DeleteBlogs)

    return r
}