package params

// UsePulseChainNetworkConfig uses the PulseChain beacon chain mainnet network config.
func UsePulseChainNetworkConfig() {
	cfg := BeaconNetworkConfig().Copy()
	cfg.ContractDeploymentBlock = 17233000
	cfg.BootstrapNodes = []string{
		"enr:-MK4QLMoST7zES5B03faU_ANy-dZp0I1fyLOcGRBKexc4-bgZeuKZYWLOx_RahC5Wa2pE8B-5VBvpXf3RfplWHDferGGAYgLF_jCh2F0dG5ldHOIAAAAAAAAAACEZXRoMpCEXuUqAAADbP__________gmlkgnY0gmlwhBLYujGJc2VjcDI1NmsxoQM6gonfZ2At1gVixDiXYTT49JGCA9mk-qXH_HbgxAshNYhzeW5jbmV0cwCDdGNwgjLIg3VkcIIu4A", // bootnode-aws-us-east-2-001
		"enr:-MK4QFJATDWtgTvu8EFa6Ukzn_04Aj0GOpCvRxDaaXH8IScfImfRAFFa8Oz0_6FXX5jxlZwUVXvsez9pSswceG3zCVyGAYgLF_e_h2F0dG5ldHOIAAAAAAAAAACEZXRoMpCEXuUqAAADbP__________gmlkgnY0gmlwhBK8WXeJc2VjcDI1NmsxoQPgYyuwBxuNAxq_rqXaoc4IfOPZRUW1A6DeD-1OOlz6G4hzeW5jbmV0cwCDdGNwgjLIg3VkcIIu4A", // bootnode-aws-us-east-2-002
		"enr:-MK4QBrmFJu1tQ7M-wRryZ_lAlM8xBFQ6Et2j5OoJK6rgJnlSCzU4ASwShkhnewN5dG7mwxO3MGnLauepIMaNctDEEOGAYgLF_xnh2F0dG5ldHOIAAAAAAAAAACEZXRoMpCEXuUqAAADbP__________gmlkgnY0gmlwhBLY3ZiJc2VjcDI1NmsxoQNrGUin_QhnH664dMUYCEr25_MBZnqsOm4qIlqB0KMw0YhzeW5jbmV0cwCDdGNwgjLIg3VkcIIu4A", // bootnode-aws-us-east-2-003
	}
	OverrideBeaconNetworkConfig(cfg)
}

// PulseChainConfig defines the config for the PulseChain beacon chain mainnet.
func PulseChainConfig() *BeaconChainConfig {
	cfg := MainnetConfig().Copy()
	cfg.ConfigName = PulseChainName
	cfg.PresetBase = "pulsechain"

	// preset overrides
	cfg.BaseRewardFactor = 64000
	cfg.EffectiveBalanceIncrement = 1 * 1e15
	cfg.MaxEffectiveBalance = 32 * 1e15

	// config overrides
	cfg.TerminalTotalDifficulty = "58750003716598352947541"
	cfg.MinGenesisActiveValidatorCount = 4096
	cfg.MinGenesisTime = 1683776400
	cfg.GenesisForkVersion = []byte{0x00, 0x00, 0x03, 0x69}
	cfg.GenesisDelay = 300
	cfg.AltairForkVersion = []byte{0x00, 0x00, 0x03, 0x6a}
	cfg.AltairForkEpoch = 1
	cfg.BellatrixForkVersion = []byte{0x00, 0x00, 0x03, 0x6b}
	cfg.BellatrixForkEpoch = 2
	cfg.CapellaForkVersion = []byte{0x00, 0x00, 0x03, 0x6c}
	cfg.CapellaForkEpoch = 3
	cfg.SecondsPerSlot = 10
	cfg.EjectionBalance = 16 * 1e15
	cfg.DepositChainID = 369
	cfg.DepositNetworkID = 369
	cfg.DepositContractAddress = "0x3693693693693693693693693693693693693693"

	cfg.InitializeForkSchedule()
	return cfg
}
