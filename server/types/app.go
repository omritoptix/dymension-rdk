package types

import (
	"encoding/json"
	"io"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

type (
	// AppOptions defines an interface that is passed into an application
	// constructor, typically used to set BaseApp options that are either supplied
	// via config file or through CLI arguments/flags. The underlying implementation
	// is defined by the server package and is typically implemented via a Viper
	// literal defined on the server Context. Note, casting Get calls may not yield
	// the expected types and could result in type assertion errors. It is recommend
	// to either use the cast package or perform manual conversion for safety.
	AppOptions interface {
		Get(string) interface{}
	}
	// ExportedApp represents an exported app state, along with
	// validators, consensus params and latest app height.
	ExportedApp struct {
		// AppState is the application state as JSON.
		AppState json.RawMessage
		// Validators is the exported validator set.
		Validators []tmtypes.GenesisValidator
		// Height is the app's latest block height.
		Height int64
		// ConsensusParams are the exported consensus params for ABCI.
		ConsensusParams *abci.ConsensusParams
		// Bech32 prefixs
		Bech32Prefix string
	}

	// AppExporter is a function that dumps all app state to
	// JSON-serializable structure and returns the current validator set.
	AppExporter func(log.Logger, dbm.DB, io.Writer, int64, bool, []string, AppOptions) (ExportedApp, error)
)
