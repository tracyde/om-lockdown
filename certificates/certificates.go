package certificates

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"log"

	"github.com/tracyde/om-lockdown/session"
)

const (
	certTmpPublic  = "/tmp/public.crt"
	certPublic     = "/var/tempest/cert/tempest.crt"
	certTmpPrivate = "/tmp/private.key"
	certPrivate    = "/var/tempest/cert/tempest.key"
)

// UpdateCertificates updates the certificates on the opsman VM
func UpdateCertificates(cert, key string, session *session.Session) error {
	c, err := ioutil.ReadFile(cert)
	if err != nil {
		return err
	}

	k, err := ioutil.ReadFile(key)
	if err != nil {
		return err
	}

	// Verify cert that was passed in is a valid cert
	block, _ := pem.Decode(c)
	if block == nil || block.Type != "CERTIFICATE" {
		return errors.New("failed to parse certificate PEM")
	}
	_, err = x509.ParseCertificate(block.Bytes)
	if err != nil {
		return err
	}

	// Verify key that was passed in is a valid private key
	block, _ = pem.Decode(k)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return errors.New("failed to parse private key PEM")
	}

	// Create tmp file with cert
	r, err := session.ExecuteCmd("cat << EOF > " + certTmpPublic + "\n" + string(c) + "EOF")
	if err != nil {
		log.Printf("Error(s) occurred:\n%s", r)
		return err
	}

	// Create tmp file with private key
	r, err = session.ExecuteCmd("cat << EOF > " + certTmpPrivate + "\n" + string(k) + "EOF")
	if err != nil {
		log.Printf("Error(s) occurred:\n%s", r)
		return err
	}

	// TODO see if there is a way to test combination of cert and private key

	// TODO add logic to test certificate with opsman fqdn, ensure cert covers the fqdn

	// Copy tmp file to final public cert location
	r, err = session.ExecuteCmd("echo \"" + session.Password + "\" | sudo -S -k " +
		"cp " + certTmpPublic + " " + certPublic)
	if err != nil {
		log.Printf("Error(s) occurred:\n%s", r)
		return err
	}

	// Copy tmp file to final private key location
	r, err = session.ExecuteCmd("echo \"" + session.Password + "\" | sudo -S -k " +
		"cp " + certTmpPrivate + " " + certPrivate)
	if err != nil {
		log.Printf("Error(s) occurred:\n%s", r)
		return err
	}

	// Remove tmp cert and key
	r, err = session.ExecuteCmd("rm " + certTmpPublic + " && rm " + certTmpPrivate)
	if err != nil {
		log.Printf("Error(s) occurred:\n%s", r)
		return err
	}

	// Restart nginx process to make changes take effect
	r, err = session.ExecuteCmd("echo \"" + session.Password + "\" | sudo -S -k /etc/init.d/nginx stop &&" +
		"echo \"" + session.Password + "\" | sudo -S -k /etc/init.d/nginx start")
	if err != nil {
		log.Printf("Error(s) occurred:\n%s", r)
		return err
	}

	return nil
}
