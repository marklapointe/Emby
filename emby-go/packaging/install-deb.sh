#!/bin/bash
# Install Emby Server on Debian/Ubuntu systems

set -e

# Check if running as root
if [ "$(id -u)" -ne 0 ]; then
    echo "This script must be run as root"
    exit 1
fi

# Install dependencies
apt-get update
apt-get install -y wget ca-certificates ffmpeg

# Create emby user
if ! id -u emby > /dev/null 2>&1; then
    addgroup --system emby
    adduser --system --ingroup emby emby
fi

# Create directories
mkdir -p /opt/emby-server
mkdir -p /var/lib/emby-server
mkdir -p /var/log/emby-server
mkdir -p /etc/emby-server

# Copy binary
if [ -f "/tmp/emby-server" ]; then
    cp /tmp/emby-server /opt/emby-server/emby-server
else
    echo "Binary not found at /tmp/emby-server"
    exit 1
fi

# Set permissions
chown -R emby:emby /opt/emby-server
chown -R emby:emby /var/lib/emby-server
chown -R emby:emby /var/log/emby-server
chown -R emby:emby /etc/emby-server

# Create systemd service
cat > /etc/systemd/system/emby-server.service << EOF
[Unit]
Description=Emby Server
After=network.target

[Service]
Type=simple
User=emby
Group=emby
ExecStart=/opt/emby-server/emby-server
Restart=on-failure
RestartSec=5
Environment=EMBY_SERVER_CONFIG_DIR=/etc/emby-server
Environment=EMBY_SERVER_DATA_DIR=/var/lib/emby-server
Environment=EMBY_SERVER_LOG_DIR=/var/log/emby-server

[Install]
WantedBy=multi-user.target
EOF

# Enable and start service
systemctl daemon-reload
systemctl enable emby-server
systemctl start emby-server

echo "Emby Server installed successfully!"
echo "Access the web interface at http://localhost:8096"
