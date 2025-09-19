package migrate

import (
	"encoding/json"
	"fmt"
	"io"
)

type Config struct {
	ImportPaths map[string]string
	Modules     []string
}

var DefaultConfig = Config{
	ImportPaths: map[string]string{
		"github.com/ipfs/go-bitswap":                     "github.com/nnlgsakib/wwfs-sdk/bitswap",
		"github.com/ipfs/go-ipfs-files":                  "github.com/nnlgsakib/wwfs-sdk/files",
		"github.com/ipfs/tar-utils":                      "github.com/nnlgsakib/wwfs-sdk/tar",
		"github.com/ipfs/interface-go-ipfs-core":         "github.com/nnlgsakib/wwfs-sdk/coreiface",
		"github.com/ipfs/go-unixfs":                      "github.com/nnlgsakib/wwfs-sdk/ipld/unixfs",
		"github.com/ipfs/go-pinning-service-http-client": "github.com/nnlgsakib/wwfs-sdk/pinning/remote/client",
		"github.com/ipfs/go-path":                        "github.com/nnlgsakib/wwfs-sdk/path",
		"github.com/ipfs/go-namesys":                     "github.com/nnlgsakib/wwfs-sdk/namesys",
		"github.com/ipfs/go-mfs":                         "github.com/nnlgsakib/wwfs-sdk/mfs",
		"github.com/ipfs/go-ipfs-provider":               "github.com/nnlgsakib/wwfs-sdk/provider",
		"github.com/ipfs/go-ipfs-pinner":                 "github.com/nnlgsakib/wwfs-sdk/pinning/pinner",
		"github.com/ipfs/go-ipfs-keystore":               "github.com/nnlgsakib/wwfs-sdk/keystore",
		"github.com/ipfs/go-filestore":                   "github.com/nnlgsakib/wwfs-sdk/filestore",
		"github.com/ipfs/go-ipns":                        "github.com/nnlgsakib/wwfs-sdk/ipns",
		"github.com/ipfs/go-blockservice":                "github.com/nnlgsakib/wwfs-sdk/blockservice",
		"github.com/ipfs/go-ipfs-chunker":                "github.com/nnlgsakib/wwfs-sdk/chunker",
		"github.com/ipfs/go-fetcher":                     "github.com/nnlgsakib/wwfs-sdk/fetcher",
		"github.com/ipfs/go-ipfs-blockstore":             "github.com/nnlgsakib/wwfs-sdk/blockstore",
		"github.com/ipfs/go-ipfs-posinfo":                "github.com/nnlgsakib/wwfs-sdk/filestore/posinfo",
		"github.com/ipfs/go-ipfs-util":                   "github.com/nnlgsakib/wwfs-sdk/util",
		"github.com/ipfs/go-ipfs-ds-help":                "github.com/nnlgsakib/wwfs-sdk/datastore/dshelp",
		"github.com/ipfs/go-verifcid":                    "github.com/nnlgsakib/wwfs-sdk/verifcid",
		"github.com/ipfs/go-ipfs-exchange-offline":       "github.com/nnlgsakib/wwfs-sdk/exchange/offline",
		"github.com/ipfs/go-ipfs-routing":                "github.com/nnlgsakib/wwfs-sdk/routing",
		"github.com/ipfs/go-ipfs-exchange-interface":     "github.com/nnlgsakib/wwfs-sdk/exchange",
		"github.com/ipfs/go-merkledag":                   "github.com/nnlgsakib/wwfs-sdk/ipld/merkledag",
		"github.com/boxo/ipld/car":                       "github.com/ipld/go-car",

		// Pre Boxo rename
		"github.com/ipfs/go-libipfs/gateway":               "github.com/nnlgsakib/wwfs-sdk/gateway",
		"github.com/ipfs/go-libipfs/bitswap":               "github.com/nnlgsakib/wwfs-sdk/bitswap",
		"github.com/ipfs/go-libipfs/files":                 "github.com/nnlgsakib/wwfs-sdk/files",
		"github.com/ipfs/go-libipfs/tar":                   "github.com/nnlgsakib/wwfs-sdk/tar",
		"github.com/ipfs/go-libipfs/coreiface":             "github.com/nnlgsakib/wwfs-sdk/coreiface",
		"github.com/ipfs/go-libipfs/unixfs":                "github.com/nnlgsakib/wwfs-sdk/ipld/unixfs",
		"github.com/ipfs/go-libipfs/pinning/remote/client": "github.com/nnlgsakib/wwfs-sdk/pinning/remote/client",
		"github.com/ipfs/go-libipfs/path":                  "github.com/nnlgsakib/wwfs-sdk/path",
		"github.com/ipfs/go-libipfs/namesys":               "github.com/nnlgsakib/wwfs-sdk/namesys",
		"github.com/ipfs/go-libipfs/mfs":                   "github.com/nnlgsakib/wwfs-sdk/mfs",
		"github.com/ipfs/go-libipfs/provider":              "github.com/nnlgsakib/wwfs-sdk/provider",
		"github.com/ipfs/go-libipfs/pinning/pinner":        "github.com/nnlgsakib/wwfs-sdk/pinning/pinner",
		"github.com/ipfs/go-libipfs/keystore":              "github.com/nnlgsakib/wwfs-sdk/keystore",
		"github.com/ipfs/go-libipfs/filestore":             "github.com/nnlgsakib/wwfs-sdk/filestore",
		"github.com/ipfs/go-libipfs/ipns":                  "github.com/nnlgsakib/wwfs-sdk/ipns",
		"github.com/ipfs/go-libipfs/blockservice":          "github.com/nnlgsakib/wwfs-sdk/blockservice",
		"github.com/ipfs/go-libipfs/chunker":               "github.com/nnlgsakib/wwfs-sdk/chunker",
		"github.com/ipfs/go-libipfs/fetcher":               "github.com/nnlgsakib/wwfs-sdk/fetcher",
		"github.com/ipfs/go-libipfs/blockstore":            "github.com/nnlgsakib/wwfs-sdk/blockstore",
		"github.com/ipfs/go-libipfs/filestore/posinfo":     "github.com/nnlgsakib/wwfs-sdk/filestore/posinfo",
		"github.com/ipfs/go-libipfs/util":                  "github.com/nnlgsakib/wwfs-sdk/util",
		"github.com/ipfs/go-libipfs/datastore/dshelp":      "github.com/nnlgsakib/wwfs-sdk/datastore/dshelp",
		"github.com/ipfs/go-libipfs/verifcid":              "github.com/nnlgsakib/wwfs-sdk/verifcid",
		"github.com/ipfs/go-libipfs/exchange/offline":      "github.com/nnlgsakib/wwfs-sdk/exchange/offline",
		"github.com/ipfs/go-libipfs/routing":               "github.com/nnlgsakib/wwfs-sdk/routing",
		"github.com/ipfs/go-libipfs/exchange":              "github.com/nnlgsakib/wwfs-sdk/exchange",

		// Unmigrated things
		"github.com/ipfs/go-libipfs/blocks": "github.com/ipfs/go-block-format",
		"github.com/nnlgsakib/wwfs-sdk/blocks":       "github.com/ipfs/go-block-format",
	},
	Modules: []string{
		"github.com/ipfs/go-bitswap",
		"github.com/ipfs/go-ipfs-files",
		"github.com/ipfs/tar-utils",
		"gihtub.com/ipfs/go-block-format",
		"github.com/ipfs/interface-go-ipfs-core",
		"github.com/ipfs/go-unixfs",
		"github.com/ipfs/go-pinning-service-http-client",
		"github.com/ipfs/go-path",
		"github.com/ipfs/go-namesys",
		"github.com/ipfs/go-mfs",
		"github.com/ipfs/go-ipfs-provider",
		"github.com/ipfs/go-ipfs-pinner",
		"github.com/ipfs/go-ipfs-keystore",
		"github.com/ipfs/go-filestore",
		"github.com/ipfs/go-ipns",
		"github.com/ipfs/go-blockservice",
		"github.com/ipfs/go-ipfs-chunker",
		"github.com/ipfs/go-fetcher",
		"github.com/ipfs/go-ipfs-blockstore",
		"github.com/ipfs/go-ipfs-posinfo",
		"github.com/ipfs/go-ipfs-util",
		"github.com/ipfs/go-ipfs-ds-help",
		"github.com/ipfs/go-verifcid",
		"github.com/ipfs/go-ipfs-exchange-offline",
		"github.com/ipfs/go-ipfs-routing",
		"github.com/ipfs/go-ipfs-exchange-interface",
		"github.com/ipfs/go-libipfs",
	},
}

func ReadConfig(r io.Reader) (Config, error) {
	var config Config
	err := json.NewDecoder(r).Decode(&config)
	if err != nil {
		return Config{}, fmt.Errorf("reading and decoding config: %w", err)
	}
	return config, nil
}
