#!/bin/bash
# Upgrade Emby Server from C# to Go version

set -e

echo "=== Emby Server Upgrade Script ==="
echo "This script will upgrade your Emby Server from C# to Go version."
echo ""

# Check if running as root
if [ "$(id -u)" -ne 0 ]; then
    echo "This script must be run as root"
    exit 1
fi

# Backup current data
echo "Backing up current data..."
BACKUP_DIR="/tmp/emby-backup-$(date +%Y%m%d-%H%M%S)"
mkdir -p "$BACKUP_DIR"

if [ -d "/var/lib/emby-server" ]; then
    cp -r /var/lib/emby-server "$BACKUP_DIR/data"
    echo "Data backed up to $BACKUP_DIR/data"
fi

if [ -d "/etc/emby-server" ]; then
    cp -r /etc/emby-server "$BACKUP_DIR/config"
    echo "Config backed up to $BACKUP_DIR/config"
fi

if [ -d "/var/log/emby-server" ]; then
    cp -r /var/log/emby-server "$BACKUP_DIR/logs"
    echo "Logs backed up to $BACKUP_DIR/logs"
fi

echo ""
echo "Backup complete."
echo ""

# Stop current service
echo "Stopping current Emby Server service..."
systemctl stop emby-server 2>/dev/null || true
systemctl stop emby-server-csharp 2>/dev/null || true

echo ""
echo "Installing Go version..."

# Install Go version
if [ -f "/tmp/emby-server" ]; then
    cp /tmp/emby-server /opt/emby-server/emby-server
    echo "Binary installed."
else
    echo "Binary not found at /tmp/emby-server"
    echo "Please download the Go binary first and place it at /tmp/emby-server"
    exit 1
fi

# Set permissions
chown -R emby:emby /opt/emby-server
chown -R emby:emby /var/lib/emby-server
chown -R emby:emby /var/log/emby-server
chown -R emby:emby /etc/emby-server

# Start new service
echo "Starting new Emby Server service..."
systemctl daemon-reload
systemctl start emby-server
systemctl enable emby-server

echo ""
echo "=== Upgrade Complete ==="
echo "Emby Server has been upgraded to the Go version."
echo "Access the web interface at http://localhost:8096"
echo ""
echo "Backup saved to: $BACKUP_DIR"
echo "If you need to rollback, run:"
echo "  systemctl stop emby-server"
echo "  cp -r $BACKUP_DIR/data/* /var/lib/emby-server/"
echo "  cp -r $BACKUP_DIR/config/* /etc/emby-server/"
echo "  systemctl start emby-server-csharp"
