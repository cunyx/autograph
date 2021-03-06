package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"go.mozilla.org/autograph/signer"
)

// MonitoringInputData is the data signed by the monitoring handler
var MonitoringInputData = []byte(`AUTOGRAPH MONITORING`)

func (a *autographer) addMonitoring(monitoring authorization) error {
	if monitoring.Key == "" {
		return nil
	}
	if _, ok := a.auths["monitor"]; ok {
		return fmt.Errorf("user 'monitor' is reserved for monitoring, duplication is not permitted")
	}
	monitoring.hawkMaxTimestampSkew = time.Minute
	a.auths["monitor"] = monitoring
	return nil
}

func (a *autographer) handleMonitor(w http.ResponseWriter, r *http.Request) {
	rid := getRequestID(r)
	starttime := time.Now()
	userid, err := a.authorize(r, []byte(""))
	if err != nil {
		httpError(w, r, http.StatusUnauthorized, "authorization verification failed: %v", err)
		return
	}
	if userid != "monitor" {
		httpError(w, r, http.StatusUnauthorized, "user is not permitted to call this endpoint")
		return
	}

	sigerrstrs := make([]string, len(a.signers))
	sigresps := make([]signatureresponse, len(a.signers))
	var wg sync.WaitGroup
	for i, s := range a.signers {
		wg.Add(1)

		go func(i int, s signer.Signer) {
			defer wg.Done()

			if _, ok := s.(signer.DataSigner); !ok {
				sigerrstrs[i] = fmt.Sprintf("signer %q does not implement the DataSigner interface", s.Config().ID)
				return
			}

			// base64 of the string 'AUTOGRAPH MONITORING'
			sig, err := s.(signer.DataSigner).SignData(MonitoringInputData, s.(signer.DataSigner).GetDefaultOptions())
			if err != nil {
				sigerrstrs[i] = fmt.Sprintf("signing failed with error: %v", err)
				return
			}

			encodedsig, err := sig.Marshal()
			if err != nil {
				sigerrstrs[i] = fmt.Sprintf("encoding failed with error: %v", err)
				return
			}

			sigerrstrs[i] = ""
			sigresps[i] = signatureresponse{
				Ref:       id(),
				Type:      s.Config().Type,
				Mode:      s.Config().Mode,
				SignerID:  s.Config().ID,
				PublicKey: s.Config().PublicKey,
				Signature: encodedsig,
				X5U:       s.Config().X5U,
			}
		}(i, s)
	}
	wg.Wait()
	for _, errstr := range sigerrstrs {
		if errstr != "" {
			httpError(w, r, http.StatusInternalServerError, errstr)
			return
		}
	}

	respdata, err := json.Marshal(sigresps)
	if err != nil {
		httpError(w, r, http.StatusInternalServerError, "signing failed with error: %v", err)
		return
	}
	if a.debug {
		log.Printf("signature response: %s", respdata)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(respdata)
	log.WithFields(log.Fields{
		"rid":     rid,
		"user_id": userid,
		"t":       int32(time.Since(starttime) / time.Millisecond), //  request processing time in ms
	}).Info("monitoring operation succeeded")
}
