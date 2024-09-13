package web

import (
	"ergo.services/ergo/act"
	"ergo.services/ergo/gen"
	"ergo.services/ergo/meta"
	"net/http"
	"time"
)

type WebPool struct {
	act.Pool
}

func (w *WebPool) Init(args ...any) (act.PoolOptions, error) {
	var webOptions meta.WebServerOptions
	var poolOptions act.PoolOptions

	mux := http.NewServeMux()

	// Create and spawn root handler meta-process.
	root := meta.CreateWebHandler(meta.WebHandlerOptions{
		RequestTimeout: 5 * time.Hour,
	})
	rootId, err := w.SpawnMeta(root, gen.MetaOptions{})
	if err != nil {
		return poolOptions, err
	}

	mux.Handle("/", root)
	w.Log().Info("started WebHandler to serve '/' (meta-process: %s)", rootId)

	webOptions.Port = 9090
	webOptions.Host = "localhost"
	webOptions.Handler = mux

	webserver, err := meta.CreateWebServer(webOptions)
	if err != nil {
		return poolOptions, err
	}
	webserverId, err := w.SpawnMeta(webserver, gen.MetaOptions{})
	if err != nil {
		// Clean up
		webserver.Terminate(err)
		return poolOptions, err
	}

	w.Log().Info("started Web server %s: use http://%s:%d/", webserverId, webOptions.Host, webOptions.Port)

	poolOptions.WorkerFactory = NewWorker

	return poolOptions, nil
}

func NewWebPool() gen.ProcessBehavior {
	return &WebPool{}
}
