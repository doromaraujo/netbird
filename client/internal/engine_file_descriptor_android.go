//go:build android

package internal

import (
	log "github.com/sirupsen/logrus"
	"strings"
)

func (e *Engine) getFileDescriptor() (int, error) {
	address := e.config.WgAddr
	mtu := int(e.config.MTU)
	dnsEnabled := !e.config.DisableDNS

	log.Debugf("dns server is nil? %t", e.dnsServer == nil)
	log.Debugf("route manager is nil? %t", e.routeManager == nil)

	var dns string
	var routesString string
	var searchDomainsString string

	if dnsEnabled && e.dnsServer != nil && e.routeManager != nil {
		dns = e.dnsServer.DnsIP().String()
		log.Debugf("dns: %v", dns)

		searchDomains := e.dnsServer.SearchDomains()
		routes := e.routeManager.InitialRouteRange()

		routesString = strings.Join(routes, ";")
		searchDomainsString = strings.Join(searchDomains, ";")
	}

	return e.mobileDep.TunAdapter.ConfigureInterface(address, mtu, dns, searchDomainsString, routesString)
}
