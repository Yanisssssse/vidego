package transcoding

import "context"

type Transcoder struct {
}

func (t *Transcoder) Transcode(ctx context.Context, vidkey string) error { return nil }

func (t *Transcoder) Cancel(ctx context.Context, vidkey string) {}

func (t *Transcoder) HLSP(ctx context.Context, vidkey string) ([]byte, error) {
	return nil, nil
}
