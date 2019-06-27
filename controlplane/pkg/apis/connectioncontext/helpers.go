package connectioncontext

import (
	"fmt"
	"net"
)

//Validate - Checks IPContext and returns error if IPContext is not valid
func (c *IPContext) Validate() error {
	if c == nil {
		return fmt.Errorf("connectionContext should not be nil")
	}
	for _, route := range c.GetRoutes() {
		if route.GetPrefix() == "" {
			return fmt.Errorf("connectionContext.Route.Prefix is required and cannot be empty/nil: %v", c)
		}
		_, _, err := net.ParseCIDR(route.GetPrefix())
		if err != nil {
			return fmt.Errorf("connectionContext.Route.Prefix should be a valid CIDR address: %v", c)
		}
	}

	for _, neighbor := range c.GetIpNeighbors() {
		if neighbor.GetIp() == "" {
			return fmt.Errorf("connectionContext.IpNeighbors.Ip is required and cannot be empty/nil: %v", c)
		}
		if neighbor.GetHardwareAddress() == "" {
			return fmt.Errorf("connectionContext.IpNeighbors.HardwareAddress is required and cannot be empty/nil: %v", c)
		}
	}
	return nil
}

//Validate - Checks DNSConfig and returns error if DNSConfig is not valid
func (c *DNSConfig) Validate() error {
	if c == nil {
		return fmt.Errorf("dnsConfig should not be nil")
	}
	if c.Prioritize && len(c.ResolvesDomains) == 0 {
		return fmt.Errorf("resolves domains can not be empty")
	}
	if len(c.DnsServerIps) == 0 {
		return fmt.Errorf("dnsConfig should have recoreds")
	}
	return nil
}

//MeetsRequirements - Checks IPContext with required fields and returns error if IPContext is not valid
func (c *IPContext) MeetsRequirements(original *IPContext) error {
	if c == nil {
		return fmt.Errorf("connectionContext should not be nil")
	}

	err := c.Validate()
	if err != nil {
		return err
	}
	if original.GetDstIpRequired() && c.GetDstIpAddr() == "" {
		return fmt.Errorf("connectionContext.DestIp is required and cannot be empty/nil: %v", c)
	}
	if original.GetSrcIpRequired() && c.GetSrcIpAddr() == "" {
		return fmt.Errorf("connectionContext.SrcIp is required cannot be empty/nil: %v", c)
	}

	return nil
}

func (c *ExtraPrefixRequest) IsValid() error {
	if c == nil {
		return fmt.Errorf("ExtraPrefixRequest should not be nil...")
	}

	if c.RequiredNumber < 1 {
		return fmt.Errorf("ExtraPrefixRequest.RequiredNumber should be positive number >=1: %v", c)
	}
	if c.RequestedNumber < 1 {
		return fmt.Errorf("ExtraPrefixRequest.RequestedNumber should be positive number >=1: %v", c)
	}

	if c.RequiredNumber > c.RequestedNumber {
		return fmt.Errorf("ExtraPrefixRequest.RequiredNumber should be less or equal to ExtraPrefixRequest.RequestedNumber >=1: %v", c)
	}

	if c.PrefixLen < 1 {
		return fmt.Errorf("ExtraPrefixRequest.PrefixLen should be positive number >=1: %v", c)
	}

	// Check protocols
	if c.AddrFamily == nil {
		return fmt.Errorf("ExtraPrefixRequest.AfFamily should not be nil...")
	}

	switch c.AddrFamily.Family {
	case IpFamily_IPV4:
		if c.PrefixLen > 32 {
			return fmt.Errorf("ExtraPrefixRequest.PrefixLen should be positive number >=1 and <=32 for IPv4 %v", c)
		}
	case IpFamily_IPV6:
		if c.PrefixLen > 128 {
			return fmt.Errorf("ExtraPrefixRequest.PrefixLen should be positive number >=1 and <=32 for IPv6 %v", c)
		}
	}

	return nil
}
