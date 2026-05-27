package s3afero

import (
	"time"

	"github.com/spf13/afero"
)

var modBaseTime = time.Date(2019, 1, 1, 12, 0, 0, 0, time.UTC)

type modTimeCalc func() (time.Duration, error)

func modTimeFsCalc(fs afero.Fs) modTimeCalc { _ = "STUB: not implemented"; return *new(modTimeCalc) }

// modTimeResolution returns a best-effort guess at the resolution of the file
// modification time for a given afero.Fs.
func modTimeResolution(fs afero.Fs) (dur time.Duration, rerr error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// FIXME(bw): I was writing a fancy algorithm to search for this, but it's trickier
// than it first appears. My first attempt was to simply set the mod time to
// something known to be representable on all the filesystems we support,
// subtracting 1ns and seeing what that rounds down to... works fine for NTFS
// but FAT32 rounds up in this situation!

// ext4, APFS

// NTFS

// HFS+
// FAT32
