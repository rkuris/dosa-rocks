package main

import (
    "crypto/tls"
    "net/http"
    "log"
)

func ApiServer(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte("This is an example server.\n"))
}

func redirectServer(w http.ResponseWriter, req *http.Request) {
    newUrl := "https://" + req.Host + req.URL.Path
    http.Redirect(w, req, newUrl, http.StatusTemporaryRedirect)
}

func main() {
    go http.ListenAndServe(":80", http.HandlerFunc(redirectServer))
    cfg := &tls.Config{
	MinVersion: tls.VersionTLS12,
	CurvePreferences: []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
	PreferServerCipherSuites: true,
	CipherSuites: []uint16{
	    tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
	    tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
            tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
            tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
            tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
            tls.TLS_RSA_WITH_AES_256_CBC_SHA,
        },
    }
    fs := http.FileServer(http.Dir(".well-known"))
    mux := http.NewServeMux()
    mux.Handle("/.well-known/", http.StripPrefix("/.well-known", fs))
    mux.HandleFunc("/api", ApiServer)
    srv := &http.Server{
	Addr: ":443",
	Handler: mux,
	TLSConfig: cfg,
	TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
    }
    log.Fatal(srv.ListenAndServeTLS("/etc/letsencrypt/live/dosa.authbank.com/fullchain.pem", "/etc/letsencrypt/live/dosa.authbank.com/privkey.pem"))
}
