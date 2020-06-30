# QuickStart

```go
res, err := s.db.Collection(ns).InsertOne(context.Background(), Account{
		Namespace: ns,
		Username:  username,
		Metadata:  metadata,
		Secret: Secret{
			Salt: salt,
			Hash: hash,
		},
		Timestamp: timestamp.Create(),
	})
```