# nullable

This is a POC implementation of a `Nullable` type that allows knowing whether a field was nullable and/or optional in Go, meant for JSON serialization and deserialization.

It relies on Go generics, but rather than using a custom struct, it relies on a map as its underlying type. The problem with using a struct is that:
- on deserialization from JSON, any `null`` value in the JSON input combined with a struct pointer leads Go to skip the UnmarhalJSON for that value
- on serializaton to JSON, we do want the benefits of a `nil` pointer to be combined with `omitempty` in order for a non-specified field to not appear in the resulting JSON

Using a map as underlying type allows using `nil` while still making use of a non-explicit-pointer to a struct.

This is inspired from [this post by KumanekoSakura](https://github.com/golang/go/issues/64515#issuecomment-1841024193).

See some examples of usage in `nullable_test.go`.
