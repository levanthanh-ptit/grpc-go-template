package mongo_utils

import "go.mongodb.org/mongo-driver/mongo/options"

func ToFindOptions(opts ...interface{}) []*options.FindOptions {
	arr := []*options.FindOptions{}
	for _, v := range opts {
		op := v.(options.FindOptions)
		arr = append(arr, &op)
	}
	return arr
}

func ToFindOneOptions(opts ...interface{}) []*options.FindOneOptions {
	arr := []*options.FindOneOptions{}
	for _, v := range opts {
		op := v.(options.FindOneOptions)
		arr = append(arr, &op)
	}
	return arr
}

func ToInsertOneOptions(opts ...interface{}) []*options.InsertOneOptions {
	arr := []*options.InsertOneOptions{}
	for _, v := range opts {
		op := v.(options.InsertOneOptions)
		arr = append(arr, &op)
	}
	return arr
}

func ToInsertManyOptions(opts ...interface{}) []*options.InsertManyOptions {
	arr := []*options.InsertManyOptions{}
	for _, v := range opts {
		op := v.(options.InsertManyOptions)
		arr = append(arr, &op)
	}
	return arr
}

func ToUpdateOptions(opts ...interface{}) []*options.UpdateOptions {
	arr := []*options.UpdateOptions{}
	for _, v := range opts {
		op := v.(options.UpdateOptions)
		arr = append(arr, &op)
	}
	return arr
}

func ToDeleteOptions(opts ...interface{}) []*options.DeleteOptions {
	arr := []*options.DeleteOptions{}
	for _, v := range opts {
		op := v.(options.DeleteOptions)
		arr = append(arr, &op)
	}
	return arr
}
