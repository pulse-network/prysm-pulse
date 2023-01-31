package params

// UsePulseChainTestnetNetworkConfig uses the PulseChain beacon chain testnet network config.
func UsePulseChainTestnetNetworkConfig() {
	cfg := BeaconNetworkConfig().Copy()
	cfg.ContractDeploymentBlock = 16492700
	cfg.BootstrapNodes = []string{
		// TODO: Add bootnode ENRs
	}
	OverrideBeaconNetworkConfig(cfg)
}

// PulseChainTestnetConfig defines the config for the PulseChain beacon chain testnet.
func PulseChainTestnetConfig() *BeaconChainConfig {
	cfg := MainnetConfig().Copy()
	cfg.ConfigName = PulseChainTestnetName
	cfg.TerminalTotalDifficulty = "58750003716598352947541"
	cfg.MinGenesisActiveValidatorCount = 5000
	cfg.MinGenesisTime = 1674864000
	cfg.GenesisForkVersion = []byte{0x00, 0x00, 0x09, 0x42}
	cfg.GenesisDelay = 300
	cfg.AltairForkVersion = []byte{0x00, 0x00, 0x09, 0x43}
	cfg.AltairForkEpoch = 1
	cfg.BellatrixForkVersion = []byte{0x00, 0x00, 0x09, 0x44}
	cfg.BellatrixForkEpoch = 2
	cfg.SecondsPerSlot = 10
	cfg.MaxEffectiveBalance = 32 * 1e15
	cfg.EjectionBalance = 16 * 1e15
	cfg.DepositChainID = 942
	cfg.DepositNetworkID = 942
	cfg.DepositContractAddress = "0x3693693693693693693693693693693693693693"
	cfg.InitializeForkSchedule()
	return cfg
}
