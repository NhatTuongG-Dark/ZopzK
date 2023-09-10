package JsonParse

type Options struct {
	FakeSlaves struct {
		Status bool `json:"Status"`
		MinGen int `json:"MinGenSet"`
		MaxGen int `json:"MaxGenSet"`
	} `json:"FakeSlaves"`

	SlaveTransition struct {
		Status bool `json:"Status"`
		LoopbackPort string `json:"LoopbackPort"`
	} `json:"SlaveTransition"`
}

type Mirai struct {
	AttackPrefix string `json:"AttackPrefix"`
	CurrentCNC string `json:"CurrentCNC"`
	Username string `json:"Username"`
	Password string `json:"Password"`

	Methods []struct {
		Name string `json:"Name"`
		Description string `json:"Description"`
		Args []string `json:"Args"`
	} `json:"Methods"`
}


type Slaves struct {
	AttackPrefix string `json:"AttackPrefix"`
	Slaves []struct {
		ID        uint16   `json:"ID"`
		Name  string   `json:"Name"`
		Description string `json:"Description"`
		DefaultPort string `json:"DefaultPort"`
		Admin bool `json:"Admin"`
		Vip bool `json:"Vip"`
	} `json:"Slaves"`
}

type AttackSync struct {
	Attacks []struct {
		Name        string   `json:"Name"`
		MethodName  string   `json:"AttackName"`
		Description string   `json:"Description"`
		API        string `json:"API"`

		Moderation struct {
			LimitMaxTime bool `json:"LimitMaxTime"`
			MaxTimeAllow int `json:"MaxTime"`
		} `json:"moderation"`

		DefaultPort string `json:"DefaultPort"`
		AdminMethod bool     `json:"AdminMethod"`
		VipMethod   bool   `json:"VipMethod"`
	} `json:"Attacks"`
}


type ConfigSync struct {
	Masters struct {
		MasterPort   string `json:"MasterPort"`
		MaxAuthTries int    `json:"MaxAuthTries"`
	} `json:"Masters"`

	Slaves struct {
		Status bool `json:"Status"`
		Slaves string `json:"Slaves"`
	}
	
	SQL struct {
		SQLHost     string `json:"SQL-Host"`
		SQLUsername string `json:"SQL-Username"`
		SQLPassword string `json:"SQL-Password"`
		SQLName     string `json:"SQL-Name"`
		SQLAudit    struct {
			Status     bool   `json:"Status"`
			LogDetails string `json:"LogDetails"`
			Username   string `json:"Username"`
		} `json:"SQL-Audit"`
	} `json:"SQL"`

	Branding struct {
		AppName	string `json:"AppName"`
		MaxConcurrentsReached string `json:"MaxConcurrentsReached"`
		OverAllowedTime string `json:"OverAllowedTime"`
	} `json:"Branding"`

	DisabledCommands []string `json:"DisabledCommands"`

	Controls struct {
		Catpcha struct {
			Status	bool `json:"Status"`
			Header    string `json:"Header"`
			AllowedAttempts	int `json:"AllowedAttempts"`
			AdminBypass	bool `json:"AdminBypass"`
			Question struct {
				MinGen	int `json:"MinGen"`
				MaxGen	int `json:"MaxGen"`
			} `json:"Question"`
		} `json:"Catpcha"`

		MFA struct {
			ForceMFA bool `json:"ForceMFA"`
			AdminBypassForce bool `json:"AdminBypassForce"`
			AppName string `json:"AppName"`
		}

		Ongoing struct {
			ShowGlobalAttacks	bool `json:"ShowGlobalAttacks"`
		}`json:"Ongoing"`
		
	} `json:"Controls"`

	Plans struct {
		MinPassLen int `json:"minPassLen"`
		MaxSessions int `json:"MaxSessions"`
		Presets    []struct {
			Name        string `json:"Name"`
			Description string `json:"Description"`
			Concurrents int    `json:"Concurrents"`
			Cooldown    int    `json:"cooldown"`
			MaxTime     int    `json:"maxTime"`
			PowerSavingExempt bool `json:"PowerSavingExempt"`
			BypassBlacklist bool `json:"BypassBlacklist"`
			PlanLenDays int    `json:"PlanTimeDays"`
			MaxSessions int    `json:"MaxSessions"`
			Reseller    bool   `json:"Reseller"`
			Vip         bool   `json:"VIP"`
		} `json:"Presets"`
	} `json:"Plans"`

	Attacks struct {
		AttackDebug bool `json:"AttackDebug"`
		Blacklists	[]string `json:"Blacklists"`
	} `json:"Attacks"`
}

	
	
type LiveWireDLC struct {
	TableGradient struct {
		Status      bool   `json:"Status"`
		GradientOne string `json:"GradientOne"`
		GradientTwo string `json:"GradientTwo"`
	} `json:"TableGradient"`
	TitleSpinner struct {
		Active bool   `json:"Active"`
		Position string `json:"Position"`
		Frames []string `json:"Frames"`
	} `json:"TitleSpinner"`
}

type MethodSuggestion struct {
	Active bool `json:"Active"`
	ISPs   []struct {
		ASNs             []string `json:"ASNs"`
		RecommendMethods string   `json:"RecommendMethods"`
	} `json:"ISPs"`
}