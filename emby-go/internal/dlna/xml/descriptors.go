package xml

// DeviceDescriptor returns the UPnP device descriptor XML.
func DeviceDescriptor(port int) string {
	return `<?xml version="1.0"?>
<root xmlns="urn:schemas-upnp-org:device-1-0">
  <specVersion>
    <major>1</major>
    <minor>0</minor>
  </specVersion>
  <device>
    <deviceType>urn:schemas-upnp-org:device:MediaServer:1</deviceType>
    <friendlyName>Emby Server</friendlyName>
    <manufacturer>Emby</manufacturer>
    <manufacturerURL>https://emby.media</manufacturerURL>
    <modelName>Emby Server</modelName>
    <modelNumber>1.0</modelNumber>
    <modelURL>https://emby.media</modelURL>
    <serialNumber>EMBY</serialNumber>
    <UDN>uuid:emby-go</UDN>
    <iconList>
      <icon>
        <mimetype>image/png</mimetype>
        <width>48</width>
        <height>48</height>
        <depth>32</depth>
        <url>/logo.png</url>
      </icon>
    </iconList>
    <serviceList>
      <service>
        <serviceType>urn:schemas-upnp-org:service:ConnectionManager:1</serviceType>
        <serviceId>ConnectionManager</serviceId>
        <controlURL>/upnp/control/ConnectionManager</controlURL>
        <eventSubURL>/upnp/event/ConnectionManager</eventSubURL>
      </service>
      <service>
        <serviceType>urn:schemas-upnp-org:service:ContentDirectory:1</serviceType>
        <serviceId>ContentDirectory</serviceId>
        <controlURL>/upnp/control/ContentDirectory</controlURL>
        <eventSubURL>/upnp/event/ContentDirectory</eventSubURL>
      </service>
    </serviceList>
  </device>
</root>`
}

// ContentDirectoryAction returns a ContentDirectory SOAP action response.
func ContentDirectoryAction(action string, body string) string {
	switch action {
	case "GetSystemInfo":
		return `<?xml version="1.0"?>
<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">
  <s:Body>
    <u:GetSystemInfoResponse xmlns:u="urn:schemas-upnp-org:service:ContentDirectory:1">
      <NumberOfDisplays>1</NumberOfDisplays>
      <VideoEncoderMaxStreamCount>4</VideoEncoderMaxStreamCount>
      <VideoTranscodingCapacity>4</VideoTranscodingCapacity>
    </u:GetSystemInfoResponse>
  </s:Body>
</s:Envelope>`
	case "Browse":
		return `<?xml version="1.0"?>
<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">
  <s:Body>
    <u:BrowseResponse xmlns:u="urn:schemas-upnp-org:service:ContentDirectory:1">
      <Result>&lt;DIDL-Lite xmlns="urn:schemas-upnp-org:metadata-1-0/DIDL-Lite/"&gt;&lt;/DIDL-Lite&gt;</Result>
      <NumberReturned>0</NumberReturned>
      <TotalMatches>0</TotalMatches>
      <UpdateID>0</UpdateID>
    </u:BrowseResponse>
  </s:Body>
</s:Envelope>`
	default:
		return `<?xml version="1.0"?>
<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">
  <s:Body>
    <s:Fault>
      <faultcode>s:Client</faultcode>
      <faultstring>UPnPError</faultstring>
      <detail>
        <UPnPError xmlns="urn:schemas-upnp-org:control-1-0">
          <errorCode>701</errorCode>
          <errorDescription>Action not supported</errorDescription>
        </UPnPError>
      </detail>
    </s:Fault>
  </s:Body>
</s:Envelope>`
	}
}
