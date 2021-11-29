package driver

import(
	UserModel "antonio/features/user/data"
)

type Entity struct{
	Model interface{}
}

func registerEntities() []Entity{
	return []Entity{
		{UserModel.User{}},
	}
}