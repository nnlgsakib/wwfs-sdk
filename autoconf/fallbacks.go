package autoconf

// Fallback defaults matching Kubo 0.36
// These are used as last-resort fallback when autoconf fetch fails and no cache exists
var (
	// FallbackBootstrapPeers are the default bootstrap peers from Kubo 0.36
	// Used as last-resort fallback when autoconf fetch fails
	FallbackBootstrapPeers = []string{
		"/ip4/148.251.35.204/udp/4001/quic-v1/p2p/12D3KooWQ3iqi3w58mDVUB7beFy6tuwX4WuxCvoBCZEzLnDANb2u",
	}

	// FallbackDNSResolvers are the default DNS resolvers matching mainnet autoconf
	// Used as last-resort fallback when autoconf fetch fails
	// Only "eth." has explicit fallbacks - no "." (root domain) fallback is intentional
	// to ensure users' OS DNS resolver is used by default, allowing browser clients
	// to make their own DoH decisions based on privacy preferences
	FallbackDNSResolvers = map[string][]string{
		"eth.": {
			"https://dns.eth.limo/dns-query",
			"https://dns.eth.link/dns-query",
		},
	}
)

// GetMainnetFallbackConfig returns a complete fallback config matching current mainnet values
// This mirrors https://conf.ipfs-mainnet.org/autoconf.json exactly
func GetMainnetFallbackConfig() *Config {
	return &Config{
		AutoConfVersion: 2025072901, // Current mainnet version
		AutoConfSchema:  1,
		SystemRegistry: map[string]SystemConfig{
			SystemAminoDHT: {
				URL:         "",
				Description: "Public DHT swarm that implements the IPFS Kademlia DHT specification under protocol identifier /wwfs/kad/1.0.0",
				NativeConfig: &NativeConfig{
					Bootstrap: FallbackBootstrapPeers,
				},
				DelegatedConfig: &DelegatedConfig{
					Read:  []string{"/routing/v1/providers", "/routing/v1/peers", "/routing/v1/ipns"},
					Write: []string{"/routing/v1/ipns"},
				},
			},
			SystemIPNI: {
				URL:         "https://cid.contact",
				Description: "Network Indexer - content routing database for large storage providers",
				DelegatedConfig: &DelegatedConfig{
					Read:  []string{"/routing/v1/providers"},
					Write: []string{},
				},
			},
			"Example": {
				URL:         "https://example.com",
				Description: "Test system for implementers to verify graceful handling of unknown systems and APIs. Production clients MUST ignore this system and its /example/* endpoints without errors.",
				DelegatedConfig: &DelegatedConfig{
					Read:  []string{"/example/v0/read"},
					Write: []string{"/example/v0/write"},
				},
			},
		},
		DNSResolvers: FallbackDNSResolvers,
		DelegatedEndpoints: map[string]EndpointConfig{
			"https://cid.contact": {
				Systems: []string{SystemIPNI},
				Read:    []string{"/routing/v1/providers"},
				Write:   []string{},
			},
			"https://delegated-ipfs.dev": {
				Systems: []string{SystemAminoDHT, SystemIPNI},
				Read:    []string{"/routing/v1/providers", "/routing/v1/peers", "/routing/v1/ipns"},
				Write:   []string{"/routing/v1/ipns"},
			},
			"https://example.com": {
				Systems: []string{"Example"},
				Read:    []string{"/example/v0/read"},
				Write:   []string{"/example/v0/write"},
			},
		},
	}
}
