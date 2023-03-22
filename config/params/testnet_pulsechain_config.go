package params

// UsePulseChainTestnetNetworkConfig uses the PulseChain beacon chain testnet network config.
func UsePulseChainTestnetNetworkConfig() {
	cfg := BeaconNetworkConfig().Copy()
	cfg.ContractDeploymentBlock = 16492700
	cfg.BootstrapNodes = []string{
		"enr:-MK4QC37TFAfc973oOezRlVoOCygtjT-rlOoKbbuZNmrJ5dhXS-IfrsH3yhjNP0dfy3-UpyFZy2hy6lOE__ykFfj3lKGAYa95Co5h2F0dG5ldHOIAAAAAAAAAACEZXRoMpBbnJIGAAAJRP__________gmlkgnY0gmlwhAPsylWJc2VjcDI1NmsxoQOpEhsSXVShFW4yvaww_SI0A-H0pix0aJlOdYJgyIgbjYhzeW5jbmV0cwCDdGNwgjLIg3VkcIIu4A", // bootnode-aws-us-east-1-001
		"enr:-MK4QM1JpOXnj-zpjfPvG1GkCEvjYMg8dEk6t7VLtpFuionhBz59n2ZIwixpO2exzoNLMV4_v7jCHGQqi0zYtc-Gp3OGAYa94lhqh2F0dG5ldHOIAAAAAAAAAACEZXRoMpBbnJIGAAAJRP__________gmlkgnY0gmlwhCzKVYOJc2VjcDI1NmsxoQK24rqFwR7W3HJgLVVGDSMy8PiMculxF6VOJgAlG4wmXohzeW5jbmV0cwCDdGNwgjLIg3VkcIIu4A", // bootnode-aws-us-east-1-002
		"enr:-MK4QBAlroGpM1xtk7WzWP8lbKnI2UjVpoQNsKJeNRS-kVvFCswNyVRBZHwMOvfW2G3j0qaaDsUpMxXY-t9LdGAZgQGGAYa94IQmh2F0dG5ldHOIAAAAAAAAAACEZXRoMpBbnJIGAAAJRP__________gmlkgnY0gmlwhK6BYN-Jc2VjcDI1NmsxoQJNoBFGkhcAMKIbrDPHoI7dYVAY99Z832TimlqhpoYo7YhzeW5jbmV0cwCDdGNwgjLIg3VkcIIu4A", // bootnode-aws-us-east-1-003
	}
	OverrideBeaconNetworkConfig(cfg)
}

// PulseChainTestnetConfig defines the config for the PulseChain beacon chain testnet.
func PulseChainTestnetConfig() *BeaconChainConfig {
	cfg := MainnetConfig().Copy()
	cfg.ConfigName = PulseChainTestnetName
	cfg.TerminalTotalDifficulty = "58750003716598352947541"
	cfg.MinGenesisActiveValidatorCount = 4096
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
