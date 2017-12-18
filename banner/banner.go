package banner

import (
	"io/ioutil"
	"log"

	"github.com/tracyde/om-lockdown/session"
)

const (
	sshBanner    = "/etc/issue.net"
	loginBanner  = "/etc/issue"
	sshTmpBanner = "/tmp/banner.txt"
)

// UpdateBanner updates the opsman VM banner
func UpdateBanner(banner string, session *session.Session) error {
	b, err := ioutil.ReadFile(banner)
	if err != nil {
		return err
	}

	// Create tmp file with banner
	r, err := session.ExecuteCmd("cat << EOF > " + sshTmpBanner + "\n" + string(b) + "EOF")
	if err != nil {
		log.Printf("Error(s) occurred:\n%s", r)
		return err
	}

	// Copy tmp file to banner locations
	r, err = session.ExecuteCmd("echo \"" + session.Password + "\" | sudo -S -k " +
		"cp " + sshTmpBanner + " " + sshBanner + " && " +
		"echo \"" + session.Password + "\" | sudo -S -k " +
		"cp " + sshTmpBanner + " " + loginBanner)
	if err != nil {
		log.Printf("Error(s) occurred:\n%s", r)
		return err
	}

	// Remove tmp file with banner
	r, err = session.ExecuteCmd("rm " + sshTmpBanner)
	if err != nil {
		log.Printf("Error(s) occurred:\n%s", r)
		return err
	}

	return nil
}
