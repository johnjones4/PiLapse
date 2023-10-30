#!/bin/bash

CONFIG_DIR=/etc/pilapse
BIN_DIR=/usr/local/bin
SYSTEMD_DIR=/etc/systemd/system/

echo "Removing configuration files"
rm -rf "$CONFIG_DIR" || true

echo "Removing binaries"
rm "$BIN_DIR/pilapse" || true

echo "Removing Systemd configuration"
systemctl stop pilapse.service || true
systemctl disable pilapse.service || true
rm "$SYSTEMD_DIR/pilapse.service" || true

echo "Uninstallation complete"
