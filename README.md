# OM-Lockdown
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

![locked out](https://media.giphy.com/media/o0pzviHOcqABa/giphy.gif)

## Purpose

Allows options for hardening a Pivotal OpsManager VM.

## Installation

To install `om-lockdown` go to [Releases](https://github.com/tracyde/om-lockdown/releases)

## Example

```
$ om-lockdown \
    -hostname 192.168.56.101 \
    -banner ~/banner.txt \
    -cert ~/om-cert.crt  \
    -key ~/om-cert.key
```

### Options

#### `-hostname`

The `-hostname` flag accepts either the IP address or FQDN of your Pivotal
OpsManager VM.

#### `-username`

Specify the username that `om-lockdown` will use to login to your Pivotal
OpsManager VM by setting the `-username` flag. If this flag is not set then
the default user `ubuntu` is used.

__This flag overrides the environment variable `OM_VMUSERNAME`.__

#### `-password`

The `-password` flag is the password associated with the username flag.

__This flag overrides the environment variable `OM_VMPASSWORD`.__

#### `-banner`

The `-banner` flag takes a path to a text file. That text file will be
set as both the login and ssh banner (`/etc/issue` and `/etc/issue.net`) for 
the Pivotal OpsManager VM.

#### `-cert`

The `-cert` flag takes a path to a PEM encoded certificate file. That certificate
file will be used as the Pivotal OpsManager nginx TLS certificate 
(`/var/tempest/cert/tempest.crt`). The nginx service will also be restarted to 
make the changes take effect.

__The `-cert` flag must be used in conjunction with the `-key` flag.__

#### `-key`

The `-key` flag takes a path to a PEM encoded RSA private key file. That private
key file will be used as the Pivotal OpsManager nginx TLS private key 
(`/var/tempest/cert/tempest.key`). The nginx service will also be restarted to 
make the changes take effect.

__The `-key` flag must be used in conjunction with the `-cert` flag.__
