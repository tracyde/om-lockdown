package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tracyde/om-lockdown/certificates"

	"github.com/tracyde/om-lockdown/banner"
	"github.com/tracyde/om-lockdown/session"
)

var (
	bannerFile = flag.String("banner", "", "file to use as ssh banner")
	hostname   = flag.String("hostname", "", "resolvable fqdn or ip address of opsmanager vm")
	username   = flag.String("username", "ubuntu", "username used to connect to opsmanager vm - overwritten by `OM_VMUSERNAME`")
	password   = flag.String("password", "", "password used with username to connect to opsmanager vm - overwritten by `OM_VMPASSWORD`")
	certFile   = flag.String("cert", "", "PEM file to be used as opsmanager TLS certificate")
	keyFile    = flag.String("key", "", "PEM file to be used as opsmanager TLS private key")
)

func usage() {
	fmt.Fprintf(os.Stderr,
		"usage: om-lockdown \n"+
			"       om-lockdown -banner=/etc/issue\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {

	// Parse flags
	flag.Usage = usage
	flag.Parse()

	// Read environment variables
	if *username == "" {
		*username = os.Getenv("OM_VMUSERNAME")
	}

	if *password == "" {
		*password = os.Getenv("OM_VMPASSWORD")
	}

	// Check usage: command line args
	if *hostname == "" || *username == "" || *password == "" {
		fmt.Fprintln(os.Stderr, "missing args.")
		usage()
	}

	session := session.NewGeneric(*hostname, *username, *password)

	if *bannerFile != "" {
		err := banner.UpdateBanner(*bannerFile, session)
		if err != nil {
			log.Fatalf("Error running command: %s", err)
		}
	}

	if *certFile != "" || *keyFile != "" {
		if *certFile == "" || *keyFile == "" {
			log.Fatalf("Both cert and key options must be set!")
		}
		err := certificates.UpdateCertificates(*certFile, *keyFile, session)
		if err != nil {
			log.Fatalf("Error updating certificate: %s", err)
		}
	}
}
