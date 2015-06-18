package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

const (
	defaultStaticCacheSize = 1024 * 1024 * 64 // 64 MiB
)

// Convenience function for serving only a single file
// (quick and easy way to view a README.md file)
func serveStaticFile(filename, colonPort string) {
	log.Info("Serving " + filename + " on " + serverHost + colonPort)
	mux := http.NewServeMux()
	cache := newFileCache(defaultStaticCacheSize, false, 0) // 64 MiB cache, no cache compression, no per-file size limit
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Server", versionString)
		filePage(w, req, filename, nil, nil, cache)
	})
	HTTPserver := newServerConfiguration(mux, false, serverHost+colonPort)
	if err := HTTPserver.ListenAndServe(); err != nil {
		// Can't serve HTTP, give up
		fatalExit(err)
	}
}
