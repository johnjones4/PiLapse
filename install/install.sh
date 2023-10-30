#!/bin/bash

CONFIG_DIR=/etc/pilapse
BIN_DIR=/usr/local/bin
SYSTEMD_DIR=/etc/systemd/system/

echo "Creating configuration director $CONFIG_DIR"
mkdir "$CONFIG_DIR"

echo "Installing configuration files"
cp -f config.json "$CONFIG_DIR/"
cp -f env "$CONFIG_DIR/"

echo "Installing binares"
cp -f pilapse "$BIN_DIR/"

echo "Installing Systemd service"
systemctl stop pilapse.service || true
systemctl disable pilapse.service || true
cp -f pilapse.service "$SYSTEMD_DIR/"
systemctl daemon-reload
systemctl enable pilapse.service
systemctl start pilapse.service

echo "Installtion complete. Please refer back to the Readme for further instructions"
