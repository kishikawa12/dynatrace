package notification

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dtcookie/dynatrace/apis/cluster"
	"github.com/dtcookie/dynatrace/apis/problems"
	"github.com/dtcookie/dynatrace/log"
)

func newListener(config *Config, handler Handler) *listener {
	return &listener{config: config, handler: handler}
}

// listener TODO: documentation
type listener struct {
	handler Handler
	config  *Config
}

func (listener *listener) listen() {
	var clusterVersion string
	var err error

	clusterAPI := cluster.NewAPI(&listener.config.Credentials)
	if clusterVersion, err = clusterAPI.Get(); err != nil {
		log.Error(err)
		return
	}
	log.Info("Dynatrace Cluster Version: " + clusterVersion)
	http.HandleFunc("/", listener.handleHTTP)
	log.Info(fmt.Sprintf("Listening on port %d for incoming problem notifications.", listener.config.ListenPort))
	http.ListenAndServe(fmt.Sprintf(":%d", listener.config.ListenPort), nil)
}

func (listener *listener) handleHTTP(w http.ResponseWriter, request *http.Request) {
	var err error
	var body []byte

	fmt.Println("received")

	if request.Method != http.MethodPost {
		if listener.config.verbose {
			log.Warn(request.Method + " responding with " + http.StatusText(http.StatusMethodNotAllowed))
		}
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	if request.ContentLength == 0 {
		if listener.config.verbose {
			log.Warn("responding with " + http.StatusText(http.StatusBadRequest))
		}
		http.Error(w, http.StatusText(http.StatusBadRequest)+": missing request body", http.StatusBadRequest)
		return
	}
	var contentType string
	contentType = request.Header.Get("content-type")
	if contentType != "application/json" {
		if listener.config.verbose {
			log.Warn("responding with " + http.StatusText(http.StatusBadRequest) + ": expected content-type 'application/json'")
		}
		http.Error(w, http.StatusText(http.StatusBadRequest)+": expected content-type 'application/json'", http.StatusBadRequest)
		return
	}
	if body, err = ioutil.ReadAll(request.Body); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError)+": "+err.Error(), http.StatusInternalServerError)
		log.Error(err)
		return
	}

	var defNotification Default
	if err = json.Unmarshal(body, &defNotification); err != nil {
		log.Error(err)
		return
	}
	if (defNotification.Title == "Dynatrace problem notification test run") || (defNotification.PID == "999999") {
		log.Info("Dynatrace problem notification test run successful")
		http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
		return
	}

	if listener.config.verbose {
		log.Info("received problem notification " + toJSON(defNotification))
	} else {
		if defNotification.PID != "" {
			log.Info("received problem notification for PID " + defNotification.PID)
		}
	}

	if defNotification.PID == "" {
		if listener.config.verbose {
			log.Warn("received problem notification without PID")
		}
		http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
		return
	}
	if listener.config.verbose {
		log.Info("querying for problem details")
	}
	problemAPI := problems.NewAPI(&listener.config.Credentials)
	var problem *problems.Problem
	if problem, err = problemAPI.Get(defNotification.PID); err != nil {
		log.Warn("querying for problem details failed: " + err.Error())
		return
	}

	problemEvent := ProblemEvent{Notification: &defNotification, Problem: problem}
	listener.handler.Handle(&problemEvent)

	http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
}

func toJSON(v interface{}) string {
	var err error
	var bytes []byte
	if bytes, err = json.Marshal(v); err != nil {
		return err.Error()
	}
	return string(bytes)
}
