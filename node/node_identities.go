package node

const (
	legacyKeyFileName       = "key.bin"
	keyDir                  = "identities"
	supervisedIDKeyFileName = "local.key"
)

// NewIdentity creates a new identity, saves it to `keyDir/supervisedIDKeyFileName` in the config directory and
// initializes app.signers with that identity.
func (app *App) NewIdentity() error { _ = "STUB: not implemented"; return nil }

// LoadIdentities loads all existing identities from the config directory.
func (app *App) LoadIdentities() error { _ = "STUB: not implemented"; return nil }

// skip subdirectories and files in them

// skip files that are not identity files

// make sure all keys are unique
