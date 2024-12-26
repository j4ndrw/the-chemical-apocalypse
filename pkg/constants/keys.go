package constants

type keys struct {
	PlayerIdle            string
	PlayerMoveDown        string
	PlayerMoveUp          string
	PlayerMoveForward     string
	AntiPeanutIdle        string
	AntiPeanutMoveDown    string
	AntiPeanutMoveUp      string
	AntiPeanutMoveForward string
	Roam                  string
}

var Keys = keys{
	PlayerIdle:            "playeridle",
	PlayerMoveDown:        "playermovedown",
	PlayerMoveUp:          "playermoveup",
	PlayerMoveForward:     "playermoveforward",
	AntiPeanutIdle:        "antipeanutidle",
	AntiPeanutMoveDown:    "antipeanutmovedown",
	AntiPeanutMoveUp:      "antipeanutmoveup",
	AntiPeanutMoveForward: "antipeanutmoveforward",
	Roam:                  "roam",
}
