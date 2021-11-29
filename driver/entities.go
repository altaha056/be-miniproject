package driver

import(
	UserModel "antonio/features/user/data"
)

type Entity struct{
	Model interface{}
}

func collectionEntities() []Entity{
	return []Entity{
		{UserModel.User{}},
	}
}