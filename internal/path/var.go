package path

const (
	// VarBlockID is the id of the block variable.
	VarBlockID = "block"
	// VarBlock is the var path of the block variable.
	VarBlock = "{" + VarBlockID + ":" + reToken + "}"
	// VarInstanceID is the id of the instance variable.
	VarInstanceID = "instance"
	// VarInstance is the var path of the instance variable.
	VarInstance = "{" + VarInstanceID + ":" + reToken + "}"
)
