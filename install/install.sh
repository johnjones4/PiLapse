#!/bin/bash

CONFIG_DIR=/etc/pilapse
BIN_DIR=/usr/local/bin
SYSTEMD_DIR=/etc/systemd/system/

echo "Creating configuration director $CONFIG_DIR"
mkdir "$CONFIG_DIR"

echo "Installing configuration files"
cp config.json "$CONFIG_DIR/"
cp env "$CONFIG_DIR/"

echo "Installing binares"
cp pilapse "$BIN_DIR/"

echo "Installing Systemd service"
cp pilapse.service "$SYSTEMD_DIR/"
systemctl daemon-reload
systemctl enable pilapse.service
systemctl start pilapse.service

echo "Installtion complete. Please refer back to the Readme for further instructions"
