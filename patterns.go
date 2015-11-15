package noise

var HandshakeNN = HandshakePattern{
	Name: "NN",
	Messages: [][]MessagePattern{
		{MessagePatternE},
		{MessagePatternE, MessagePatternDHEE},
	},
}

var HandshakeKN = HandshakePattern{
	Name:                 "KN",
	InitiatorPreMessages: []MessagePattern{MessagePatternS},
	Messages: [][]MessagePattern{
		{MessagePatternE},
		{MessagePatternE, MessagePatternDHEE, MessagePatternDHES},
	},
}

var HandshakeNK = HandshakePattern{
	Name:                 "NK",
	ResponderPreMessages: []MessagePattern{MessagePatternS},
	Messages: [][]MessagePattern{
		{MessagePatternE, MessagePatternDHES},
		{MessagePatternE, MessagePatternDHEE},
	},
}

var HandshakeKK = HandshakePattern{
	Name:                 "KK",
	InitiatorPreMessages: []MessagePattern{MessagePatternS},
	ResponderPreMessages: []MessagePattern{MessagePatternS},
	Messages: [][]MessagePattern{
		{MessagePatternE, MessagePatternDHES, MessagePatternDHSS},
		{MessagePatternE, MessagePatternDHEE, MessagePatternDHES},
	},
}

var HandshakeNE = HandshakePattern{
	Name:                 "NE",
	ResponderPreMessages: []MessagePattern{MessagePatternS, MessagePatternE},
	Messages: [][]MessagePattern{
		{MessagePatternE, MessagePatternDHEE, MessagePatternDHSE},
		{MessagePatternE, MessagePatternDHEE},
	},
}

var HandshakeKE = HandshakePattern{
	Name:                 "KE",
	InitiatorPreMessages: []MessagePattern{MessagePatternS},
	ResponderPreMessages: []MessagePattern{MessagePatternS, MessagePatternE},
	Messages: [][]MessagePattern{
		{MessagePatternE, MessagePatternDHEE, MessagePatternDHES, MessagePatternDHSE},
		{MessagePatternE, MessagePatternDHEE, MessagePatternDHES},
	},
}

var HandshakeNX = HandshakePattern{
	Name: "NX",
	Messages: [][]MessagePattern{
		{MessagePatternE},
		{MessagePatternE, MessagePatternDHEE, MessagePatternS, MessagePatternDHSE},
	},
}

var HandshakeKX = HandshakePattern{
	Name:                 "KX",
	InitiatorPreMessages: []MessagePattern{MessagePatternS},
	Messages: [][]MessagePattern{
		{MessagePatternE},
		{MessagePatternE, MessagePatternDHEE, MessagePatternDHES, MessagePatternS, MessagePatternDHSE},
	},
}

var HandshakeXN = HandshakePattern{
	Name: "XN",
	Messages: [][]MessagePattern{
		{MessagePatternE},
		{MessagePatternE, MessagePatternDHEE},
		{MessagePatternS, MessagePatternDHSE},
	},
}

var HandshakeIN = HandshakePattern{
	Name: "IN",
	Messages: [][]MessagePattern{
		{MessagePatternS, MessagePatternE},
		{MessagePatternE, MessagePatternDHEE, MessagePatternDHES},
	},
}

var HandshakeXK = HandshakePattern{
	Name:                 "XK",
	ResponderPreMessages: []MessagePattern{MessagePatternS},
	Messages: [][]MessagePattern{
		{MessagePatternE, MessagePatternDHES},
		{MessagePatternE, MessagePatternDHEE},
		{MessagePatternS, MessagePatternDHSE},
	},
}

var HandshakeIK = HandshakePattern{
	Name:                 "IK",
	ResponderPreMessages: []MessagePattern{MessagePatternS},
	Messages: [][]MessagePattern{
		{MessagePatternE, MessagePatternDHES, MessagePatternS, MessagePatternDHSS},
		{MessagePatternE, MessagePatternDHEE, MessagePatternDHES},
	},
}

var HandshakeXE = HandshakePattern{
	Name:                 "XE",
	ResponderPreMessages: []MessagePattern{MessagePatternS, MessagePatternE},
	Messages: [][]MessagePattern{
		{MessagePatternE, MessagePatternDHEE, MessagePatternDHES},
		{MessagePatternE, MessagePatternDHEE},
		{MessagePatternS, MessagePatternDHSE},
	},
}

var HandshakeIE = HandshakePattern{
	Name:                 "IE",
	ResponderPreMessages: []MessagePattern{MessagePatternS, MessagePatternE},
	Messages: [][]MessagePattern{
		{MessagePatternE, MessagePatternDHEE, MessagePatternDHES, MessagePatternS, MessagePatternDHSE},
		{MessagePatternE, MessagePatternDHEE, MessagePatternDHES},
	},
}

var HandshakeXX = HandshakePattern{
	Name: "XX",
	Messages: [][]MessagePattern{
		{MessagePatternE},
		{MessagePatternE, MessagePatternDHEE, MessagePatternS, MessagePatternDHSE},
		{MessagePatternS, MessagePatternDHSE},
	},
}

var HandshakeIX = HandshakePattern{
	Name: "IX",
	Messages: [][]MessagePattern{
		{MessagePatternS, MessagePatternE},
		{MessagePatternE, MessagePatternDHEE, MessagePatternDHES, MessagePatternS, MessagePatternDHSE},
	},
}

var HandshakeN = HandshakePattern{
	Name:                 "N",
	ResponderPreMessages: []MessagePattern{MessagePatternS},
	Messages: [][]MessagePattern{
		{MessagePatternE, MessagePatternDHES},
	},
}

var HandshakeK = HandshakePattern{
	Name:                 "K",
	InitiatorPreMessages: []MessagePattern{MessagePatternS},
	ResponderPreMessages: []MessagePattern{MessagePatternS},
	Messages: [][]MessagePattern{
		{MessagePatternE, MessagePatternDHES, MessagePatternDHSS},
	},
}

var HandshakeX = HandshakePattern{
	Name:                 "X",
	ResponderPreMessages: []MessagePattern{MessagePatternS},
	Messages: [][]MessagePattern{
		{MessagePatternE, MessagePatternDHES, MessagePatternS, MessagePatternDHSS},
	},
}
