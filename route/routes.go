package route

import (
    "net/http"
    "GO_REST_API/controllers"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes {
	Route{
                "Ping",
                "GET",
                "/", 
                controllers.Ping,
        },
	Route{ 
                "Blog",
                "GET",
                "/blogs", 
                controllers.GetBlogs,
        },
       	Route{ 
                "Blog",
                "GET",
                "/blogs/{title}",
                controllers.GetBlog,
        },
        Route{ 
                "Blog",
                "POST",
                "/blog", 
                controllers.PostBlog,
        },
}