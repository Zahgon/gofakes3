package main

// This test proved that version-id-marker requires key-marker to be passed
// when paginating object versions. This is a huge relief as it would've been
// really painful to support arbitrary combinations.
//
// This test flushed out that we are handling version-id-marker slightly wrong.
// The following debug dump shows that if version-id-marker is present in the
// query but contains an empty string, S3 rejects the request with an invalid
// argument error:
//
//	GET /<bucket>?key-marker=810dd187e5ad28a83e8b208558823cf45d8d6e46b1f0fa56f2712de97213c343bbec1b871a4bfa92859cd142f25b9eb135008e47ad65359c4229b7919384d51dfd8f5b15e88eaec302a38447f86ef5642e03b10eac8ba81db5ddc949d769edf49e123b809800fa28e3ddef44bf96fe65ab422065e3378c9ff0c4c37a68e05f19af9d4755fbbd846cacecdaca95eecf9b9802685e7853&max-keys=1&prefix=810dd187e5ad28a83e8b208558823cf45d8d6e46b1f0fa56f2712de97213c343bbec1b871a4bfa92859cd142f25b9eb13500&version-id-marker=&versions=
//	<Error>
//	  <Code>InvalidArgument</Code>
//	  <Message>A version-id marker cannot be empty.</Message>
//	  <ArgumentName>version-id-marker</ArgumentName><ArgumentValue></ArgumentValue>
//	  <RequestId>AC8DD6B6E4826D5C</RequestId>
//	  <HostId>prA+xN7N52Ovkry/Q8jMhvbCm2MtWiraIk2SOJqleareEgUY41CmpxD1Uvcaq6YiShSGoUlDbvY=</HostId>
//	</Error>
//
// It looks as if key-marker is not validated the same way - if it's not passed,
// it just starts from the start, and ignores version-id-marker if one is passed.
type S300004ListVersionsWithVersionMarkerButNoKeyMarker struct{}

func (s S300004ListVersionsWithVersionMarkerButNoKeyMarker) Run(ctx *Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Sanity check version length

// Passing no key marker, which should be an error:

// Passing key marker but empty version, which should fail:

// Passing version-id-marker but empty key-marker returns the first page for some reason:

// Passing key marker but no version, which should be fine:
