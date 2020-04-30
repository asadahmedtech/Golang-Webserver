package auth

// import(

// )

// // Serve a reverse proxy for a given url
// func serveReverseProxy(target string, w http.ResponseWriter, r *http.Request) {
// 	// parse the url
// 	url, _ := url.Parse(target)

// 	// create the reverse proxy
// 	proxy := httputil.NewSingleHostReverseProxy(url)

// 	// Update the headers to allow for SSL redirection
// 	r.URL.Host = url.Host
// 	r.URL.Scheme = url.Scheme
// 	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
// 	r.Host = url.Host

// 	// Note that ServeHttp is non blocking and uses a go routine under the hood
// 	proxy.ServeHTTP(w, r)
// }
