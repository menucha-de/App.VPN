/*
 * App vpn
 *
 * API version: 1.0.0
 * Contact: info@menucha.de
 */

package vpn

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
)

const config string = "conf/default.ovpn"
const logfile string = "openvpn.log"

var mutex sync.Mutex
var process *os.Process

func init() {
	stat, err := os.Stat(config)
	if err != nil && os.IsNotExist(err) {
		return
	}
	if stat.Mode().Perm() == 0 {
		return
	}
	start()
}

func start() {
	mutex.Lock()
	defer mutex.Unlock()

	stdout, err := os.Create(logfile)
	if err != nil {
		log.Warn("Failed to create log file")
	}
	defer stdout.Close()
	var procAttr os.ProcAttr
	procAttr.Files = []*os.File{os.Stdin, stdout, stdout}
	log.Info("Starting OpenVPN")
	process, err = os.StartProcess("/usr/sbin/openvpn", []string{"--config", config}, &procAttr)
	if err != nil {
		log.Warn("Failed to start OpenVPN: %s\n", err)
	}
}

func stop() {
	mutex.Lock()
	defer mutex.Unlock()

	if process != nil {
		process.Kill()
		_, _ = process.Wait()
		process = nil
	}
}

// UploadOpenVPNConfig Uploads OpenVPN config
func UploadOpenVPNConfig(w http.ResponseWriter, r *http.Request) {
	f, err := os.Create(config)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	_, err = io.Copy(f, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// IsEnabled Returns the enable state if config exists
func IsEnabled(w http.ResponseWriter, r *http.Request) {
	stat, err := os.Stat(config)
	if err != nil && os.IsNotExist(err) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, strconv.FormatBool(stat.Mode().Perm() != 0))
}

// SetEnable Sets the enable state if config exists
func SetEnable(w http.ResponseWriter, r *http.Request) {
	_, err := os.Stat(config)
	if err != nil && os.IsNotExist(err) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	enable, err := strconv.ParseBool(string(b))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	stop()
	if enable {
		os.Chmod(config, 0644)
		start()
	} else {
		os.Chmod(config, 0)
	}
	w.WriteHeader(http.StatusAccepted)
}

// DownloadOpenVPNlog Provides log content
func DownloadOpenVPNlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.Header().Set("Content-Disposition", "attachment; filename=openvpn.log")
	http.ServeFile(w, r, logfile)
	w.WriteHeader(http.StatusOK)
}
