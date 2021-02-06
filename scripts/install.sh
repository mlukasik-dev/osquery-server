#!/bin/bash

# Exit if any of the commands fails.
set -e

if [ $(id -u) -ne 0 ]; then
  echo "Need to run as root"
  exit 1
fi

OSQUERY_TARBAL=osquery-4.6.0_1.linux_x86_64.tar.gz
OSQUERY_DOWNLOAD_LINK=https://pkg.osquery.io/linux/$OSQUERY_TARBAL
INSTALL_DIR=/usr/local/aosquery

echo "Downloading osquery"
curl $OSQUERY_DOWNLOAD_LINK > $OSQUERY_TARBAL

echo "Creating osquery user"
useradd -c 'Osquery user' -p 'password' -m osquery

echo "Unpacking osquery"
sudo -u osquery tar -C /home/osquery -xzf $OSQUERY_TARBAL
rm $OSQUERY_TARBAL

cp /home/osquery/usr/local/share/osquery/osquery.example.conf /home/osquery/etc/osquery/osquery.conf

sudo systemctl start osqueryd

exit 0