package params

// UsePulseChainTestnetV4NetworkConfig uses the PulseChain beacon chain testnet network config.
func UsePulseChainTestnetV4NetworkConfig() {
	cfg := BeaconNetworkConfig().Copy()
	cfg.ContractDeploymentBlock = 16492700
	cfg.BootstrapNodes = []string{
		"enr:-MK4QPzTScyPIPS7UWhaz0v4cVYKTfN2-0leb-iFKGl_xFBjDOSJenQ91BwaZ7sUXApiCPihr_Mw4L5oZJ4W5vNqffuGAYd4QZA7h2F0dG5ldHOIAAAAAAAAAACEZXRoMpB3mwAIAAAJRAEAAAAAAAAAgmlkgnY0gmlwhAOOpsuJc2VjcDI1NmsxoQNJljCY9olY73J9aYbW-9Ix72ZNzKv1AeIt6BHSqtniV4hzeW5jbmV0cwCDdGNwgjLIg3VkcIIu4A", // bootnode-aws-us-east-2-001
		"enr:-MK4QF0dJA9H3BlxJE7HEjucJD47JYdlQyJ0wbQVNrbfgM_0A-OZEb8J1jjtmiaw3ytpl-FdmofS7QPW31j1n2WImTiGAYd4QUGRh2F0dG5ldHOIAAAAAAAAAACEZXRoMpB3mwAIAAAJRAEAAAAAAAAAgmlkgnY0gmlwhBJ2vKWJc2VjcDI1NmsxoQJsvKcVWN9inS4QTW8QsPSd053XZl6MZ5zmvBXLycbONYhzeW5jbmV0cwCDdGNwgjLIg3VkcIIu4A", // bootnode-aws-us-east-2-002
		"enr:-MK4QMqHY3x4vzN1yHXY7ZCAH8ylE8dSYT3ZWSNsUBd0MJ4dIYjz94vrEUvMZGsDfJXW1aHZlpszxScBLJEQovnha6WGAYd4QUMUh2F0dG5ldHOIAAAAAAAAAACEZXRoMpB3mwAIAAAJRAEAAAAAAAAAgmlkgnY0gmlwhAOKa5CJc2VjcDI1NmsxoQIt-VoTVWzXdUTO78hhF_25-AhXgOk0k-YG3gW_dzX6YIhzeW5jbmV0cwCDdGNwgjLIg3VkcIIu4A", // bootnode-aws-us-east-2-003
	}
	OverrideBeaconNetworkConfig(cfg)
}

// PulseChainTestnetV4Config defines the config for the PulseChain beacon chain testnet.
func PulseChainTestnetV4Config() *BeaconChainConfig {
	cfg := MainnetConfig().Copy()
	cfg.ConfigName = PulseChainTestnetV4Name
	cfg.PresetBase = "pulsechain"

	// preset overrides
	cfg.BaseRewardFactor = 64000
	cfg.EffectiveBalanceIncrement = 1 * 1e15
	cfg.MaxEffectiveBalance = 32 * 1e15

	// config overrides
	cfg.TerminalTotalDifficulty = "58750003716598352947541"
	cfg.MinGenesisActiveValidatorCount = 4096
	cfg.MinGenesisTime = 1674864000
	cfg.GenesisForkVersion = []byte{0x00, 0x00, 0x09, 0x43}
	cfg.GenesisDelay = 300
	cfg.AltairForkVersion = []byte{0x00, 0x00, 0x09, 0x44}
	cfg.AltairForkEpoch = 1
	cfg.BellatrixForkVersion = []byte{0x00, 0x00, 0x09, 0x45}
	cfg.BellatrixForkEpoch = 2
	cfg.CapellaForkVersion = []byte{0x00, 0x00, 0x09, 0x46}
	cfg.CapellaForkEpoch = 4200
	cfg.SecondsPerSlot = 10
	cfg.EjectionBalance = 16 * 1e15
	cfg.DepositChainID = 943
	cfg.DepositNetworkID = 943
	cfg.DepositContractAddress = "0x3693693693693693693693693693693693693693"

	cfg.InitializeForkSchedule()
	return cfg
}
