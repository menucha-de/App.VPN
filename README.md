# OpenVPN connection setup

## Upload config and restart VPN app
    curl -v --data-binary @default.ovpn -H "Cookie: TOKEN=#####" -X PUT raspberrypi/vpn/rest/default/config
(Please use a valid auth token, you can use the token which is transferred as a cookie while communicating with the UI)
    
## Getting logs
    http://raspberrypi/vpn/rest/default/log

## Getting IP address configuration
    http://raspberrypi/#/main/settings/network/tun0
